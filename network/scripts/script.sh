#!/bin/bash

echo
echo " ____		_____			_			____		_____ "
echo "/ ___|	|_	 _|		/ \		|	_ \	|_	 _|"
echo "\___ \		| |		 / _ \	 | |_) |	 | |	"
echo " ___) |	 | |		/ ___ \	|	_ <		| |	"
echo "|____/		|_|	 /_/	 \_\ |_| \_\	 |_|	"
echo
echo "Build multi host network (BMHN) end-to-end test"
echo
CHANNEL_NAME="$1"
CHAINCODE_NAME="ptwist"
DELAY="3"
: ${CHANNEL_NAME:="ptwist"}
: ${TIMEOUT:="60"}
COUNTER=1
MAX_RETRY=5
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

echo "Channel name : "$CHANNEL_NAME

# verify the result of the end-to-end test
verifyResult () {
	if [ $1 -ne 0 ] ; then
		echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
		echo "========= ERROR !!! FAILED to execute End-2-End Scenario ==========="
		echo
	 		exit 1
	fi
}

setGlobals () {
	org=$1
	peer=$2

	CORE_PEER_LOCALMSPID="${org}MSP"
	CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/${org}.example.com/peers/peer0.${org}.example.com/tls/ca.crt
	CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/${org}.example.com/users/Admin@${org}.example.com/msp
	CORE_PEER_ADDRESS=peer${peer}.${org}.example.com:7051

	env | grep CORE
}

createChannel() {
	setGlobals MEDSOS 0


	peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx >&log.txt
	res=$?
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== Channel \"$CHANNEL_NAME\" is created successfully ===================== "
	echo
}

updateAnchorPeers() {
	for org in MEDSOS BFF BLUECITY; do
		setGlobals $org 0
			
		peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx >&log.txt
		res=$?
		cat log.txt
		verifyResult $res "Anchor peer update failed"
		echo "===================== Anchor peers for org \"$CORE_PEER_LOCALMSPID\" on \"$CHANNEL_NAME\" is updated successfully ===================== "
		sleep $DELAY
		echo
	done
}

## Sometimes Join takes time hence RETRY atleast for 5 times
joinWithRetry () {
	peer channel join -b $CHANNEL_NAME.block	>&log.txt
	res=$?
	cat log.txt
	if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
		COUNTER=` expr $COUNTER + 1`
		echo "PEER$1 failed to join the channel, Retry after 2 seconds"
		sleep $DELAY
		joinWithRetry $1
	else
		COUNTER=1
	fi
	verifyResult $res "After $MAX_RETRY attempts, PEER$ch has failed to Join the Channel"
}

joinChannel () {
	for org in MEDSOS BFF BLUECITY; do
		for peer in 0 1; do
			setGlobals $org $peer
			joinWithRetry $peer
			echo "===================== peer${peer}.${org} joined on the channel \"$CHANNEL_NAME\" ===================== "
			sleep $DELAY
			echo
		done
	done
}

installChaincode () {
	for org in MEDSOS BFF BLUECITY; do
		for peer in 0 1; do
			setGlobals ${org} ${peer}
			#peer chaincode install -n ptwist -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02 >&log.txt
			peer chaincode install -n $CHAINCODE_NAME -v 1.0 -p github.com/chaincode/ptwist >&log.txt
			res=$?
			cat log.txt
				verifyResult $res "Chaincode installation on remote peer peer${peer}.${org} has Failed"
			echo "===================== Chaincode is installed on remote peer peer${peer}.${org} ===================== "
			echo
		done
	done
}

instantiateChaincode () {
	setGlobals MEDSOS 0
	#peer chaincode instantiate -o orderer.example.com:7050 -C $CHANNEL_NAME -n $CHAINCODE_NAME -v 1.0 -c '{"Args":["a","100"]}' -P "OR ('MEDSOSMSP.member','BFFMSP.member', 'BLUECITYMSP.member')" >&log.txt
	peer chaincode instantiate -o orderer.example.com:7050 -C $CHANNEL_NAME -n $CHAINCODE_NAME -v 1.0 -c '{"Args":["a","100"]}' >&log.txt
	res=$?
	cat log.txt
	verifyResult $res "Chaincode instantiation on PEER$PEER on channel '$CHANNEL_NAME' failed"
	echo "===================== Chaincode Instantiation on PEER$PEER on channel '$CHANNEL_NAME' is successful ===================== "
	echo
}

chaincodeQuery () {
	org=${1}
	peer=${2}
	echo "===================== Querying on peer${peer}.${org} on channel '$CHANNEL_NAME'... ===================== "
	setGlobals ${org} ${peer}
	local rc=1
	local starttime=$(date +%s)

	# continue to poll
	# we either get a successful response, or reach TIMEOUT
	while test "$(($(date +%s)-starttime))" -lt "$TIMEOUT" -a $rc -ne 0
	do
		 sleep $DELAY
		 echo "Attempting to Query peer${peer}.${org} ...$(($(date +%s)-starttime)) secs"
		 peer chaincode query -C $CHANNEL_NAME -n ${CHAINCODE_NAME} -c '{"Args":["a"]}'
		 rc=$?
	done
	echo
	if test $rc -eq 0 ; then
		echo "===================== Query on peer${peer}.${org} on channel '$CHANNEL_NAME' is successful ===================== "
	else
		echo "!!!!!!!!!!!!!!! Query result on peer${peer}.${org} is INVALID !!!!!!!!!!!!!!!!"
					echo "================== ERROR !!! FAILED to execute End-2-End Scenario =================="
		echo
		exit 1
	fi
}

chaincodeInvoke () {
	PEER=$1
	setGlobals $PEER
	# while 'peer chaincode' command can get the orderer endpoint from the peer (if join was successful),
	# lets supply it directly as we know it using the "-o" option
	peer chaincode invoke -o orderer.example.com:7050 -C $CHANNEL_NAME -n mycc -c '{"Args":["invoke","a","b","10"]}' >&log.txt
	res=$?
	cat log.txt
	verifyResult $res "Invoke execution on PEER$PEER failed "
	echo "===================== Invoke transaction on PEER$PEER on channel '$CHANNEL_NAME' is successful ===================== "
	echo
}

# Create channel
echo "Creating channel..."
createChannel

# Join all the peers to the channel
echo "Having all peers join the channel..."
joinChannel

# Set the anchor peers for each org in the channel
echo "Updating anchor peers..."
updateAnchorPeers

## Install chaincode on all peers
echo "Installing chaincode..."
installChaincode

# Instantiate chaincode on
echo "Instantiating chaincode..."
instantiateChaincode

# Query on chaincode on Peer0/Org1
echo "Querying chaincode on BFF/peer0..."
#chaincodeQuery BFF 0
chaincodeQuery MEDSOS 0

# #Invoke on chaincode on Peer0/Org1
# echo "Sending invoke transaction on org1/peer0..."
# chaincodeInvoke 0

# #Query on chaincode on Peer1/Org1, check if the result is 90
# echo "Querying chaincode on org2/peer3..."
# chaincodeQuery 1 90

echo
echo "========= All GOOD, BMHN execution completed =========== "
echo

echo
echo " _____	 _	 _	 ____	 "
echo "| ____| | \ | | |	_ \	"
echo "|	_|	 |	\| | | | | | "
echo "| |___	| |\	| | |_| | "
echo "|_____| |_| \_| |____/	"
echo

exit 0

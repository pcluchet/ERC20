#!/bin/bash

readonly BASE="peer chaincode invoke -o orderer.example.com:7050"
readonly TLS="--tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

readonly fct=("totalSupply" "balanceOf" "allowance" "transfer" "approve" "transferFrom")
readonly usage=("io tsp" "io blo [address tokenOwner]" "io alo [address tokenOwner] [address spender]"			\
				"io trs [address to] [uint tokens]" "io apr [address spender] [uint tokens]"					\
				"io trf [address from] [address to] [uint tokens]")

# **************************************************************************** #
#			USAGE															   #
# **************************************************************************** #

function	fctUsage {
	printf ${fct["$1"]} ""
	printf "$2"
	echo ${usage["$1"]}
}

function	basicUsage {
	echo "----------> Usage 🔖  <----------" 
	echo ""

	for index in 0 1 2 3 4 5 ; do
		fctUsage $index ":\t"
	done
}

# **************************************************************************** #
#			PRIVATE															   #
# **************************************************************************** #

function	totalSupply {
	echo "---------------> Total Supply 💰  <---------------"
	echo ""

	cmd="-C mychannel -n mycc -c '{\"Args\":[\"totalSupply\"]}'"
	$("$BASE $TLS $cmd")
}

function	balanceOf {
	if [ $1 ]; then
		echo "---------------> Balance of $1 💵  <---------------"
		echo ""
	
		cmd="-C mychannel -n mycc -c '{\"Args\":[\"balanceOf\", \"$1\"]}'"
		$("$BASE $TLS $cmd")
	else
		fctUsage 1 ": "
	fi
}

function	allowance {
	if [ $1 ] && [ $2 ]; then
		echo "---------------> Allowance from $1 to $2 🤝  <---------------"
		echo ""

		cmd="-C mychannel -n mycc -c '{\"Args\":[\"allowance\", \"$1\", \"$2\"]}'"
		$("$BASE $TLS $cmd")
	else
		fctUsage 2 ": "
	fi
}

function	transfer {
	if [ $1 ] && [ $2 ]; then
		echo "---------------> Transfer to $1 of $2 📲  <---------------"
		echo ""

		cmd="-C mychannel -n mycc -c '{\"Args\":[\"transfer\", \"$1\", \"$2\"]}'"
		$("$BASE $TLS $cmd")
	else
		fctUsage 3 ": "
	fi
}

function	approve {
	if [ $1 ] && [ $2 ]; then
		echo "---------------> Approve from $1 of $2 👮  <---------------"
		echo ""

		cmd="-C mychannel -n mycc -c '{\"Args\":[\"approve\", \"$1\", \"$2\"]}'"
		$("$BASE $TLS $cmd")
	else
		fctUsage 4 ": "
	fi
}

function	transferFrom {
	if [ $1 ] && [ $2 ] && [ $3 ]; then
		echo "---------------> TransferFrom from $1 to $2 of $3 🚀  <---------------"
		echo ""

		cmd="-C mychannel -n mycc -c '{\"Args\":[\"transferFrom\", \"$1\", \"$2\", \"$3\"]}'"
		$("$BASE $TLS $cmd")
	else
		fctUsage 5 ": "
	fi
}


# **************************************************************************** #
#			PUBLIC															   #
# **************************************************************************** #

case $1 in
	tsp)
		totalSupply ;;
	blo)
		balanceOf $2 ;;
	alo)
		allowance $2 $3 ;;
	trs)
		transfer $2 $3 ;;
	apr)
		approve $2 $3 ;;
	trf)
		transferFrom $2 $3 $4 ;;
	*)
		basicUsage ;;
esac

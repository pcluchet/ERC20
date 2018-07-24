#!/bin/bash

#
# Copyright LOL
#
# SPDX-License-Identifier: Apache-2.0
#

os=$(uname)
if [[ ${os} == Linux ]]; then
	export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:$PWD/bin_linux:$PATH
else
	export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:$PWD/bin_macos:$PATH
fi
export FABRIC_CFG_PATH=${PWD}

# Ask user for confirmation to proceed
function askProceed () {
	read -p "Continue (y/n)? " ans
	case "$ans" in
		y|Y )
			echo "proceeding..."
			mkdir -p configtx channel-artifacts
			rm -rf configtx/* channel-artifacts/*
		;;
		n|N )
			echo "exiting..."
			exit 1
		;;
		* )
			echo "invalid response"
			askProceed
		;;
	esac
}



# Generates Org certs using cryptogen tool
function generateCerts (){
	which cryptogen
	if [ "$?" -ne 0 ]; then
		echo "cryptogen tool not found. exiting"
		exit 1
	fi
	echo
	echo "##########################################################"
	echo "##### Generate certificates using cryptogen tool #########"
	echo "##########################################################"
	if [ -d "crypto-config" ]; then
		rm -Rf crypto-config
	fi
	cryptogen generate --config=./crypto-config.yaml
	if [ "$?" -ne 0 ]; then
		echo "Failed to generate certificates..."
		exit 1
	fi
	echo
}

# Generate orderer genesis block, channel configuration transaction and
# anchor peer update transactions
function generateChannelArtifacts() {
	which configtxgen
	if [ "$?" -ne 0 ]; then
		echo "configtxgen tool not found. exiting"
		exit 1
	fi

	echo "##########################################################"
	echo "#########	Generating Orderer Genesis block ##############"
	echo "##########################################################"
	# Note: For some unknown reason (at least for now) the block file can't be
	# named orderer.genesis.block or the orderer will fail to launch!
	configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
	if [ "$?" -ne 0 ]; then
		echo "Failed to generate orderer genesis block..."
		exit 1
	fi
	echo
	echo "#################################################################"
	echo "### Generating channel configuration transaction 'channel.tx' ###"
	echo "#################################################################"
	configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
	if [ "$?" -ne 0 ]; then
		echo "Failed to generate channel configuration transaction..."
		exit 1
	fi

	echo
	echo "#################################################################"
	echo "#######		Generating anchor peer update for Org1MSP	 ##########"
	echo "#################################################################"
	orgs=(MEDSOS BFF BLUECITY)
	for org in ${orgs[@]}; do
		echo "generating \"${org}\""
		configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/${org}MSPanchors.tx -channelID $CHANNEL_NAME -asOrg ${org}MSP
		if [ "$?" -ne 0 ]; then
			echo "Failed to generate anchor peer update for ${org}MSP..."
			exit 1
		fi
		echo
	done
}

# Obtain the OS and Architecture string that will be used to select the correct
# native binaries for your platform
OS_ARCH=$(echo "$(uname -s|tr '[:upper:]' '[:lower:]'|sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
# timeout duration - the duration the CLI should wait for a response from
# another container before giving up
CLI_TIMEOUT=10
#default for delay
CLI_DELAY=3
# channel name defaults to "mychannel"
CHANNEL_NAME="ptwist"

EXPMODE="Generating certs and genesis block for"

# Announce what was requested

echo "${EXPMODE} with channel '${CHANNEL_NAME}' and CLI timeout of '${CLI_TIMEOUT}'"

# ask for confirmation to proceed
askProceed

# generate crypto-material
generateCerts
generateChannelArtifacts

#!/usr/bin/env bash

C_RED="\033[31;01m"
C_GREEN="\033[32;01m"
C_YELLOW="\033[33;01m"
C_BLUE="\033[34;01m"
C_PINK="\033[35;01m"
C_CYAN="\033[36;01m"
C_NO="\033[0m"

os=$(uname)
if [[ ${os} == Linux ]]; then
	export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:$PWD/bin_linux:$PATH
else
	export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:$PWD/bin_macos:$PATH
fi

export FABRIC_CFG_PATH="$PWD"
CHANNEL_NAME=ptwist

################################################################################
###                                FUNCTIONS                                 ###
################################################################################

function clean {
	mkdir -p ./crypto-config
	mkdir -p ./channel-artifacts
	rm -rf ./crypto-config/*
	rm -rf ./channel-artifacts/*
}

function _err {
	if [ $1 != 0 ]; then
		echo "Failed to generate $2..."
		exit 1
	fi
}

function generate {	
	### GENERATE CRYPTO MATERIALS
	cryptogen generate --config=./crypto-config.yaml
	_err $? "crypto material"
	### GENEREATE ORDERER GENESIS BLOCK
	configtxgen \
		-profile TwoOrgsOrdererGenesis \
		-outputBlock ./channel-artifacts/genesis.block
	_err $? "orderer genesis block"
	### CHANNLE CONFIGURATION
	configtxgen \
		-profile TwoOrgsChannel \
		-outputCreateChannelTx ./channel-artifacts/channel.tx \
		-channelID $CHANNEL_NAME
	_err $? "channel configuration"
}

function anchorPeer {
	for org in MEDSOS BFF BLUECITY; do
		configtxgen \
			-profile TwoOrgsChannel \
			-outputAnchorPeersUpdate ./channel-artifacts/${org}MSPanchors.tx \
			-channelID $CHANNEL_NAME \
			-asOrg ${org}MSP
		_err $? "anchor peer update for ${org}MSP"
	done
}

################################################################################
###                                   MAIN                                   ###
################################################################################

clean
generate
anchorPeer

#!/usr/bin/env bash

if [[ -t 1 ]]; then
	C_RED="\033[31;01m"
	C_GREEN="\033[32;01m"
	C_YELLOW="\033[33;01m"
	C_BLUE="\033[34;01m"
	C_PINK="\033[35;01m"
	C_CYAN="\033[36;01m"
	C_NO="\033[0m"
fi

################################################################################
###                                FUNCTIONS                                 ###
################################################################################

function		fail()
{
	printf "${C_RED}ERROR:${C_NO} %s\n" "${1}" >&2
	exit 1
}

################################################################################
###                                   MAIN                                   ###
################################################################################

##################################################
### INTRODUCTION
##################################################

os=$(uname)
if [[ ${os} == Linux ]]; then
	if [[ $(whoami) != root ]]; then
		fail "Linux users must be root to launch this script."
	fi
	case $(uname --machine) in
		x86)		computer=linux-386								;;
		x86_64)		computer=linux-amd64							;;
		armv6l)		computer=linux-armv6l							;;
		*)			fail "Linux arch must be x86 | x86_64 | armv6l"	;;
	esac
elif [[ ${os} == Darwin ]]; then
	if [[ $(uname --machine) != x86_64 ]]; then
		fail "Macos arch must be x86_64"
	fi
	computer=darwin-amd64
else
	fail "This script only handles Linux and Macos users."
fi

mkdir -p installation
cd installation

##################################################
### INSTALL GO
##################################################

go_file=go1.10.3.${computer}.tar.gz
wget https://dl.google.com/go/${go_file}
tar -C /usr/local -xzf ${go_file}
export PATH="${PATH}":/usr/local/go/bin

### INSTALL GO DEPENDENCIES
go get \
	github.com/hyperledger/fabric/core/chaincode/lib/cid \
	github.com/hyperledger/fabric/core/chaincode/shim \
	github.com/hyperledger/fabric/protos/peer

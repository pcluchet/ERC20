#!/usr/bin/env bash

C_RED="\033[31;01m"
C_GREEN="\033[32;01m"
C_YELLOW="\033[33;01m"
C_BLUE="\033[34;01m"
C_PINK="\033[35;01m"
C_CYAN="\033[36;01m"
C_NO="\033[0m"

################################################################################
###                                FUNCTIONS                                 ###
################################################################################

function		title()
{
	echo "################################################################################"
	echo "### ${1}"
	echo "################################################################################"
}

function		launch()
{
	docker exec ${1} bash -c "${2}"
	sleep 2
}

################################################################################
###                                   MAIN                                   ###
################################################################################

# WHY ?
#if [[ $(whoami) != root ]]; then
#	printf "you need to be root.\n" >&2
#	exit 0
#fi

pub_bank=$(launch	"central_bank"	"io pub --silent")
pub_bob=$(launch	"bob"			"io pub --silent")
pub_alice=$(launch	"alice"			"io pub --silent")

title "CENTRALBANK to BOB"
launch central_bank	"io apr ${pub_bob} 10"
launch central_bank	"io get ${pub_bank}"
launch bob			"io trf ${pub_bank} ${pub_bob} 10"
launch central_bank	"io blo ${pub_bank}"
launch central_bank	"io blo ${pub_bob}"

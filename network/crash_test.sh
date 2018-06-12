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

function		color()
{
	if [[ ${1} == OK ]]; then
		printf "${C_GREEN}"
	elif [[ ${1} == KO ]]; then
		printf "${C_RED}"
	else
		printf "${C_NO}"
	fi
}

function		title()
{
	color ${1}
	echo "################################################################################"
	if [[ ${1} == OK ]]; then
		printf "### ${C_CYAN}[${C_YELLOW}OK${C_CYAN}]\n"
	else
		printf "### ${C_RED}[${C_YELLOW}KO${C_RED}]\n"
	fi
	color ${1}
	printf "### ${C_NO}%s\n" "${2}"
	color ${1}
	echo "################################################################################"
	color
}

function		launch()
{
	docker exec ${1} bash -c "${2}"
	sleep 2
}

################################################################################
###                                   MAIN                                   ###
################################################################################

if [[ $(uname) == Linux ]] && [[ $(whoami) != root ]]; then
	printf "you need to be root.\n" >&2
	exit 0
fi

pub_bank=$(launch	"central_bank"	"io pub --silent")
pub_bob=$(launch	"bob"			"io pub --silent")
pub_alice=$(launch	"alice"			"io pub --silent")

##################################################
###                  SUCCESS                   ###
##################################################
title OK "CENTRALBANK to BOB"
launch central_bank	"io apr ${pub_bob} 8"
launch central_bank	"io get ${pub_bank}"
launch bob			"io trf ${pub_bank} ${pub_bob} 8"
launch central_bank	"io blo ${pub_bank}"
launch central_bank	"io blo ${pub_bob}"

title OK "CENTRALBANK to BOB by ALICE"
launch central_bank	"io apr ${pub_alice} 2"
launch central_bank	"io get ${pub_bank}"
launch alice		"io trf ${pub_bank} ${pub_bob} 2"
launch central_bank	"io blo ${pub_bank}"
launch central_bank	"io blo ${pub_bob}"

title OK "BOB to ALICE"
launch bob			"io trs ${pub_alice} 4"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo ${pub_alice}"

title OK "BOB to ALICE (0pc)"
launch bob			"io trs ${pub_alice} 0"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo ${pub_alice}"

title OK "BOB APPROVE AND DE-APPROVE to ALICE"
launch bob			"io apr ${pub_alice} 2"
launch central_bank	"io get ${pub_bob}"
launch bob			"io apr ${pub_alice} 0"
launch central_bank	"io get ${pub_bob}"

##################################################
###                  FAILURE                   ###
##################################################

title KO "BOB to ALICE (to much for bob)"
launch bob			"io trs ${pub_alice} 7"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo ${pub_alice}"

title KO "BOB to ALICE by ALICE (to much for bob to alice)"
launch bob			"io apr ${pub_alice} 10"
launch bob			"io get ${pub_bob}"
launch alice		"io trf ${pub_bob} ${pub_alice} 10"
launch central_bank	"io blo ${pub_alice}"
launch central_bank	"io blo ${pub_bob}"

title KO "BOB to BOB"
launch bob			"io trs ${pub_bob} 3"
launch central_bank	"io blo ${pub_bob}"

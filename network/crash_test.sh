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
	if [[ ${mode} == success ]]; then
		printf "${C_GREEN}"
	else
		printf "${C_RED}"
	fi
}

function		title()
{
	color
	echo "################################################################################"
	if [[ ${mode} == success ]]; then
		printf "### ${C_CYAN}[${C_YELLOW}OK${C_CYAN}]\n"
	else
		printf "### ${C_RED}[${C_YELLOW}KO${C_RED}]\n"
	fi
	color
	printf "### ${C_NO}%s\n" "${1}"
	color
	echo "################################################################################"
	printf "${C_NO}"
}

function		check_return()
{
	if [[ ${1} -eq 0 ]]; then
		printf "${C_CYAN}[${C_YELLOW}OK${C_CYAN}]\n${C_NO}"
	else
		printf "${C_RED}[${C_YELLOW}KO${C_RED}]\n${C_NO}"
	fi
}

function		launch()
{
	docker exec ${1} bash -c "${2}"
	check_return ${?}
	sleep 2
}

################################################################################
###                                   MAIN                                   ###
################################################################################

if [[ $(uname) == Linux ]] && [[ $(whoami) != root ]]; then
	printf "you need to be root.\n" >&2
	exit 0
fi

pub_bank=$(docker exec "central_bank"	bash -c "io pub --silent")
pub_bob=$(docker exec "bob"		bash -c "io pub --silent")
pub_alice=$(docker exec "alice"		bash -c "io pub --silent")

##################################################
###                  SUCCESS                   ###
##################################################
mode=success

title "CENTRALBANK to BOB by BOB"
launch central_bank	"io apr ${pub_bob} 8"
launch central_bank	"io get ${pub_bank}"
launch bob			"io trf ${pub_bank} ${pub_bob} 8"
launch central_bank	"io blo ${pub_bank}"
launch central_bank	"io blo ${pub_bob}"

title "CENTRALBANK to BOB by ALICE"
launch central_bank	"io apr ${pub_alice} 2"
launch central_bank	"io get ${pub_bank}"
launch alice		"io trf ${pub_bank} ${pub_bob} 2"
launch central_bank	"io blo ${pub_bank}"
launch central_bank	"io blo ${pub_bob}"

title "BOB to ALICE"
launch bob			"io trs ${pub_alice} 4"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo ${pub_alice}"

title "BOB APPROVE AND DE-APPROVE to ALICE"
launch bob			"io apr ${pub_alice} 2"
launch central_bank	"io get ${pub_bob}"
launch bob			"io apr ${pub_alice} 0"
launch central_bank	"io get ${pub_bob}"

title "BOB to ALICE (0pc)"
launch bob			"io trs ${pub_alice} 0"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo ${pub_alice}"

title "BOB to UNKNOWN (0pc)"
launch bob			"io trs unknown 0"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo unknown"

##################################################
###                  FAILURE                   ###
##################################################
mode=failure

title "BOB to ALICE (to much for bob)"
launch bob			"io trs ${pub_alice} 7"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo ${pub_alice}"

title "BOB to ALICE by ALICE (to much for bob to alice)"
launch bob			"io apr ${pub_alice} 10"
launch bob			"io get ${pub_bob}"
launch alice		"io trf ${pub_bob} ${pub_alice} 10"
launch central_bank	"io blo ${pub_alice}"
launch central_bank	"io blo ${pub_bob}"

title "BOB to BOB"
launch bob			"io trs ${pub_bob} 3"
launch central_bank	"io blo ${pub_bob}"

title "BOB to ALICE (negative amount)"
launch bob			"io trs ${pub_alice} -3"
launch central_bank	"io blo ${pub_bob}"
launch central_bank	"io blo ${pub_alice}"

title "CENTRALBANK to BOB by CENTRALBANK"
launch central_bank	"io apr ${pub_bob} 8"
launch central_bank	"io get ${pub_bank}"
launch central_bank	"io trf ${pub_bank} ${pub_bob} 8"
launch central_bank	"io blo ${pub_bank}"
launch central_bank	"io blo ${pub_bob}"

title "CENTRALBANK APPROVE CENTRALBANK"
launch central_bank	"io apr ${pub_bank} 10"
launch central_bank	"io apr ${pub_bob} 8"

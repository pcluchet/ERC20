#!/usr/bin/env bash

readonly BASE="peer chaincode invoke -o orderer.example.com:7050"
readonly TLS="--tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"

readonly fct=(	"totalSupply" "balanceOf" "allowance" "transfer" "approve"
				"transferFrom" "getState" "getPublicKey")
readonly usage=("io tsp"
				"io blo [address tokenOwner]"
				"io alo [address tokenOwner] [address spender]"
				"io trs [address to] [uint tokens] [publicKey]"
				"io apr [address spender] [uint tokens] [publicKey]"
				"io trf [address from] [address to] [uint tokens]"
				"io	get [key]"
				"io pub [flag silent]")

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
	local i=0
	local function_number=${#fct[@]}

	while [[ ${i} -lt ${function_number} ]]; do
		fctUsage ${i} ":\t"
		(( i++ ))
	done
}

# **************************************************************************** #
#			PRIVATE															   #
# **************************************************************************** #

function	get {
	if [ $1 ]; then
		echo "---------------> Get [$1] 🙈  <---------------"
		echo ""

		cmd=`$BASE $TLS -C ptwist -n ptwist -c "{\"Args\":[\"get\", \"$1\"]}"`
	else
		fctUsage 6 ": "
	fi
}

function	totalSupply {
	echo "---------------> Total Supply 💰  <---------------"
	echo ""

	cmd=`$BASE $TLS -C ptwist -n ptwist -c "{\"Args\":[\"totalSupply\"]}"`
}

function	balanceOf {
	if [ $1 ]; then
		echo "---------------> Balance of $1 💵  <---------------"
		echo ""
	
		cmd=`$BASE $TLS -C ptwist -n ptwist -c "{\"Args\":[\"balanceOf\", \"$1\"]}"`
	else
		fctUsage 1 ": "
	fi
}

function	allowance {
	if [ $1 ] && [ $2 ]; then
		echo "---------------> Allowance from $1 to $2 🤝  <---------------"
		echo ""

		cmd=`$BASE $TLS -C ptwist -n ptwist -c "{\"Args\":[\"allowance\", \"$1\", \"$2\"]}"`
	else
		fctUsage 2 ": "
	fi
}

function	transfer {
	if [ $1 ] && [ $2 ] && [ $3 ]; then
		echo "---------------> Transfer to $1 of $2 📲  <---------------"
		echo ""

		cmd=`$BASE $TLS -C ptwist -n ptwist -c "{\"Args\":[\"transfer\", \"$1\", \"$2\", \"$3\"]}"`
	else
		fctUsage 3 ": "
	fi
}

function	approve {
	if [ $1 ] && [ $2 ] && [ $3 ]; then
		echo "---------------> Approve from $1 of $2 👮  <---------------"
		echo ""

		cmd=`$BASE $TLS -C ptwist -n ptwist -c "{\"Args\":[\"approve\", \"$1\", \"$2\", \"$3\"]}"`
	else
		fctUsage 4 ": "
	fi
}

function	transferFrom {
	if [ $1 ] && [ $2 ] && [ $3 ]; then
		echo "---------------> TransferFrom from $1 to $2 of $3 🚀  <---------------"
		echo ""

		cmd=`$BASE $TLS -C ptwist -n ptwist -c "{\"Args\":[\"transferFrom\", \"$1\", \"$2\", \"$3\"]}"`
	else
		fctUsage 5 ": "
	fi
}

function	getPublicKey {
	if [[ -n $1 ]]; then
		if [[ $1 != "-s" ]] && [[ $1 != "--silent" ]]; then
			fctUsage 7 ": "
			exit 1
		fi
	else
		printf -- "---------------> Public key 👀 <---------------\n\n"
	fi
	# set "nullglob" shell option
	# globbing which does not match will result as empty string
	shopt -s nullglob

	publicKey=(
		$(openssl ec \
		-in		"${CORE_PEER_MSPCONFIGPATH}/keystore/"* \
		-pubout	2>&- \
		| tail -n 3 \
		| head -n 2))
	if [[ ${#publicKey[@]} -gt 0 ]]; then
		echo "${publicKey[0]}${publicKey[1]}"
	else
		if [[ -z ${CORE_PEER_MSPCONFIGPATH} ]]; then
			printf "error: CORE_PEER_MSPCONFIGPATH is not set.\n" >&2
		else
			printf "error: dammaged or non-existent msp config file in [%s]\n" \
				"${CORE_PEER_MSPCONFIGPATH}"
		fi
	fi
	shopt -u nullglob
}

# **************************************************************************** #
#			PUBLIC															   #
# **************************************************************************** #

case $1 in
	get)
		get $2 ;;
	tsp)
		totalSupply ;;
	blo)
		balanceOf $2 ;;
	alo)
		allowance $2 $3 ;;
	trs)
		transfer $2 $3 $4 ;;
	apr)
		approve $2 $3 $4 ;;
	trf)
		transferFrom $2 $3 $4 ;;
	pub)
		getPublicKey $2;;
	*)
		basicUsage ;;
esac

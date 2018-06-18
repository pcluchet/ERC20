package	main

import	"github.com/hyperledger/fabric/core/chaincode/shim"

type	SimpleAsset	struct {
}

type	UserInfos	struct {
		Amount		uint64
		Allowances	map[string]uint64
}

type	Events		struct {
		From		string
		To			string
		Value		uint64
}

type	Transaction	struct {
		From		string
		To			string
		Amount		uint64
		User		UserInfos
}

var		STUB shim.ChaincodeStubInterface
var		ledger_dev_keys map[string]bool = map[string]bool {
	"totalSupply":	true,
	"transfer":		true,
	"approval":		true}

const	centralBankTotalSupply	uint64 = 100000

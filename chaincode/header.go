package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

type SimpleAsset struct {
}

type AllowanceCouple struct {
	Spender string
	Amount  uint64
}

type AllowanceCouples []AllowanceCouple

type UserInfos struct {
	Amount     uint64
	Allowances AllowanceCouples
}

type Events struct {
	Owner		string
	Spender		string
	Value		uint64
}

var STUB shim.ChaincodeStubInterface

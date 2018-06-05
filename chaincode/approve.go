package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func generateJSON(infos UserInfos) ([]byte, error) {
	var	jsBytes		[]byte
	var	err			error

	jsBytes, err = json.Marshal(infos)
	if err != nil {
		return jsBytes, fmt.Errorf("Unable to parse given object to json")
	}
	return jsBytes, nil
}

func approve(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	// TODO :
	// [ ] real user implementation
	// [ ] check if spender is real/valid user
	// [ ] if amount is zero, delete allowance
	// [ ] better payload
	// [ ] Event
	var	err			error
	var	newAmount	uint64
	var userKey		string
	var userInfos	UserInfos
	var	jsBytes		[]byte

	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a 3 (last one is user)")
	}
	// GET USER INFOS
	userKey = args[2]
	userInfos, err = getUserInfos(stub, userKey)
	if err != nil {
		return "", err
	}
	// GET NEW APPROVED AMOUNT
	newAmount, err = strconv.ParseUint(args[1], 10, 64)
	if err != nil {
		return "", err
	}
	// SET NEW APPROVED AMOUNT
	userInfos.Allowances[args[0]] = newAmount

	fmt.Printf("%+v", userInfos)

	// CAST INFOS TO STRING
	jsBytes, err = generateJSON(userInfos)
	if err != nil {
		return "", err
	}
	// PUT STRINGIFIED INFOS TO LEDGER
	err = stub.PutState(userKey, jsBytes)
	if err != nil {
		return "", err
	}
	// SUCCESS
	return fmt.Sprintf("User added in the allowances "), nil
}

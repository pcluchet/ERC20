package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func getUserInfos(stub shim.ChaincodeStubInterface, userPublicKey string) (UserInfos, error) {
	var user UserInfos

	value, err := stub.GetState(userPublicKey)
	if err != nil {
		return user, fmt.Errorf("Failed to get asset: %s with error: %s", userPublicKey, err)
	}
	if value == nil {
		return user, fmt.Errorf("Asset not found: %s", userPublicKey)
	}

	b := bytes.NewReader(value)
	err = json.NewDecoder(b).Decode(&user)
	if err != nil {
		return user, fmt.Errorf("Failed to decode user json : %s", userPublicKey)
	}
	return user, nil
}

func balanceOf(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	// TODO :
	// [ ] User is a valid user, or check pubkey validity or whatever

	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	usrInfos, err := getUserInfos(stub, args[0])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", strconv.FormatUint(usrInfos.Amount, 10)), nil
}

//returns the amount of an allowance matching given spender in a allowanceouple list
func allowanceOfUser(allowanceList AllowanceCouples, spender string) (uint64, error) {

	for _, value := range allowanceList {
		if value.Spender == spender {
			return value.Amount, nil
		}
	}

	return 0, fmt.Errorf("Spender not found, maybe he is not allowed by given user")
}

//returns the index of a said spender in an allowance list, -1 if not found
func indexOfSpender(allowanceList AllowanceCouples, spender string) int {

	for key, value := range allowanceList {
		if value.Spender == spender {
			return key
		}
	}
	return -1
}

func allowance(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	// TODO :
	// [ ] User is a valid user, or check pubkey validity or whatever

	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a token owner public key, and a spender public key")
	}

	usrInfos, err := getUserInfos(stub, args[0])
	if err != nil {
		return "", err
	}

	value, err0 := allowanceOfUser(usrInfos.Allowances, args[1])
	if err0 != nil {
		return "", err0
	}

	return fmt.Sprintf("%s", strconv.FormatUint(value, 10)), nil
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

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

	return fmt.Sprintf("%s", usrInfos.Amount), nil
}

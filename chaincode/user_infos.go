package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (self UserInfos) Set(publicKey string) error {
	var ret		[]byte
	var err		error

	if ret, err = json.Marshal(self); err != nil {
		return err
	}
	if err = STUB.PutState(publicKey, ret); err != nil {
		return err
	}

	return nil
}

func getUserInfos(stub shim.ChaincodeStubInterface, userPublicKey string) (UserInfos, error) {
	var user UserInfos

	value, err := stub.GetState(userPublicKey)
	if err != nil {
		return user, fmt.Errorf("Failed to get asset: %s with error: %s", userPublicKey, err)
	}
	if value == nil {
		return user, fmt.Errorf("Asset not found: %s", userPublicKey)
	}
	if err = json.Unmarshal(value, &user); err != nil {
		return user, err
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
func allowanceOfUser(userInfos UserInfos, spender string) (uint64, error) {
	var	amount	uint64
	var	exist	bool

	amount, exist = userInfos.Allowances[spender]
	if exist == false {
		return 0, fmt.Errorf("Spender not found, maybe he is not allowed by given user")
	}
	return amount, nil
}

//returns the index of a said spender in an allowance list, -1 if not found
//func indexOfSpender(allowanceList AllowanceCouples, spender string) int {
//	return -1
//}

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

	value, err0 := allowanceOfUser(usrInfos, args[1])
	if err0 != nil {
		return "", err0
	}

	return fmt.Sprintf("%s", strconv.FormatUint(value, 10)), nil
}

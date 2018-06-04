package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func generateAllowanceCouple(spender string, amount string) (AllowanceCouple, error) {
	var ret AllowanceCouple
	ret.Spender = spender
	bigint, err := strconv.ParseUint(amount, 10, 64)
	if err != nil {
		return ret, fmt.Errorf("Unable to parse given int")
	}
	ret.Amount = bigint
	return ret, nil
}

func generateJSON(infos UserInfos) (string, error) {

	rejson, err := json.Marshal(infos)
	if err != nil {
		return string(rejson), fmt.Errorf("Unable to parse given object to json")
	}
	return string(rejson), nil
}

func approve(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	// TODO :
	// [ ] real user implementation
	// [ ] check if spender is real/valid user
	// [ ] if amount is zero, delete allowance
	// [ ] better payload
	// [ ] Event

	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a 3 (last one is user)")
	}

	var userkey string
	userkey = args[2]

	usrInfos, err := getUserInfos(stub, userkey)
	if err != nil {
		return "", err
	}

	//checking if spender is already in allowance list
	key := indexOfSpender(usrInfos.Allowances, args[0])
	if key != -1 {
		newamount, err1 := strconv.ParseUint(args[1], 10, 64)
		if err1 != nil {
			return "", err1
		}
		usrInfos.Allowances[key].Amount = newamount
	} else {
		newallowance, err0 := generateAllowanceCouple(args[0], args[1])
		if err0 != nil {
			return "", err0
		}
		usrInfos.Allowances = append(usrInfos.Allowances, newallowance)
	}

	fmt.Printf("%+v", usrInfos)

	rejson, err2 := generateJSON(usrInfos)
	if err != nil {
		return "", err2
	}

	err = stub.PutState(userkey, []byte(rejson))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("User added in the allowances "), nil
}

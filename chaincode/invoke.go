package main

import	"fmt"
import	"github.com/hyperledger/fabric/core/chaincode/shim"
import	"github.com/hyperledger/fabric/protos/peer"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var fct		string
	var argv	[]string
	var ret		string
	var err		error

	fmt.Println("---------------> Invoke <---------------")
	fct, argv = stub.GetFunctionAndParameters()
	STUB = stub

	switch fct {
	case "set":
		ret, err = set(argv)
	case "get":
		ret, err = get(argv)
	case "balanceOf":
		ret, err = balanceOf(argv)
	case "allowance":
		ret, err = allowance(argv)
	case "transfer":
		ret, err = transfer(argv)
	case "transferFrom":
		ret, err = transferFrom(argv)
	case "approve":
		ret, err = approve(argv)
	case "totalSupply":
		ret, err = totalSupply()
	default:
		err = fmt.Errorf("Illegal function called \n")
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(ret))
}

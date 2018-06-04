package main

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("---------------> Init <---------------")

	return shim.Success(nil)
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var fct string
	var argv []string
	var ret string
	var err error

	fct, argv = stub.GetFunctionAndParameters()
	fmt.Println("---------------> Invoke <---------------")
	fmt.Println()
	fmt.Println(fct)
	fmt.Println(argv)

	switch fct {
	case "set":
		ret, err = set(stub, argv)
	case "get":
		ret, err = get(stub, argv)
	case "balanceOf":
		ret, err = balanceOf(stub, argv)
	case "allowance":
		ret, err = allowance(stub, argv)
	default:
		err = fmt.Errorf("Illegal function called \n")
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(ret))
}

func main() {
	err := shim.Start(new(SimpleAsset))
	if err != nil {
		fmt.Fprint(os.Stderr, "Error starting Simple chaincode: %s", err)
		os.Exit(1)
	}
}

package main

import	"fmt"
import	"os"
import	"strconv"
import	"github.com/hyperledger/fabric/core/chaincode/shim"
import	"github.com/hyperledger/fabric/protos/peer"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("---------------> Init <---------------")
	var err				error
	var	bankString		string

	// SET CENTRAL BANK SUPPLY
	bankString = fmt.Sprintf("{\"amount\":%v,\"allowances\":{}}",
	/**/centralBankTotalSupply)
	err = stub.PutState(centralBankName, []byte(bankString))
	if err != nil {
		return shim.Error("Cannot set central bank")
	}
	// SET TOTAL SUPPLY
	err = stub.PutState("total_supply",
	/**/[]byte(strconv.FormatUint(centralBankTotalSupply, 10)))
	if err != nil {
		return shim.Error("Cannot set ledger total supply")
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var fct string
	var argv []string
	var ret string
	var err error

	fmt.Println("---------------> Invoke <---------------")
	fct, argv = stub.GetFunctionAndParameters()
	STUB = stub

	switch fct {
	case "set":
		ret, err = set(stub, argv)
	case "get":
		ret, err = get(stub, argv)
	case "balanceOf":
		ret, err = balanceOf(stub, argv)
	case "allowance":
		ret, err = allowance(stub, argv)
	case "transfer":
		ret, err = transfer(argv)
	case "transferFrom":
		ret, err = transferFrom(argv)
	case "approve":
		ret, err = approve(stub, argv)
	case "totalSupply":
		ret, err = totalSupply(stub)
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

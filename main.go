
package main

import "fmt"
import "os"
import "github.com/hyperledger/fabric/core/chaincode/shim"
import "github.com/hyperledger/fabric/protos/peer"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func	(t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("---------------> Init <---------------")

	return shim.Success(nil)
}

func	(t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var	fct		string
	var argv	[]string


	fct, argv = stub.GetFunctionAndParameters()
	fmt.Println("---------------> Invoke <---------------")

	switch fct {
		default:
			fmt.Fprintf(os.Stderr, "Illegal function called %s\n", fct)
	}

	return shim.Success([]byte(ret))
}

func	main() {
	err := shim.Start(new(SimpleAsset))
	if err != nil {
		fmt.Fprint(os.Stderr, "Error starting Simple chaincode: %s", err)
		os.Exit(1)
	}
}

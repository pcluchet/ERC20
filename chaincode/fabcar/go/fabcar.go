/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim"
import "github.com/hyperledger/fabric/protos/peer"
/*
import	"crypto/ecdsa"
import	"crypto/x509"
import	"encoding/pem"
import	"strings"
import	"github.com/hyperledger/fabric/core/chaincode/lib/cid"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

//Encode the given ecdsa key in pem format
func	pemEncodePublicKey(publicKey *ecdsa.PublicKey) string {
	var	x509EncodedPublicKey	[]byte
	var	pemEncodedPublicKey		[]byte

	x509EncodedPublicKey, _ = x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPublicKey = pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPublicKey})
	return string(pemEncodedPublicKey)
}

//return the public key of creator in pem format
func	getPemPublicKeyOfCreator(stub shim.ChaincodeStubInterface) (string, error) {
	var	err				error
	var	ecdsaPublicKey	*ecdsa.PublicKey
	var	cert			*x509.Certificate

	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		return "", fmt.Errorf("Error : %s", err)
	}
	ecdsaPublicKey = cert.PublicKey.(*ecdsa.PublicKey)
	return pemEncodePublicKey(ecdsaPublicKey), nil
}

//trim public key to remove newlines and begin/end tag
func	trimPemPubKey(key string) string {
	key = strings.Replace(key, "\n", "", -1)
	key = strings.Replace(key, "-----BEGIN PUBLIC KEY-----", "", -1)
	key = strings.Replace(key, "-----END PUBLIC KEY-----", "", -1)
	return key
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

//return the trimmed public key of creator in pem format
func	getPublicKey(stub shim.ChaincodeStubInterface) (string, error) {
	var	err		error
	var	spender	string

	spender, err = getPemPublicKeyOfCreator(stub)
	if err != nil {
		return "", fmt.Errorf("Cannot get creator of the transaction : %s", err)
	}
	spender = trimPemPubKey(spender)
	return spender, nil
}
*/

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// Get the args from the transaction proposal
	_, args := stub.GetFunctionAndParameters();
	fmt.Println(args);
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()
	test, _ := stub.GetCreator()

	fmt.Println(string(test))

	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
//	} //else if fn == "key" {
		//result, err = getPublicKey(stub)
	} else { // assume 'get' even if fn is nil
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
	}

	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return string(value), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

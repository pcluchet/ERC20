/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"crypto/x509"
	"encoding/pem"
	"github.com/golang/protobuf/proto"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
	//"github.com/hyperledger/fabric/vendor/github.com/golang/protobuf/proto"
)

/*
**
** Go to:
**   "https://stackoverflow.com/questions/45786023/how-to-retrieve-user-information-from-recent-version-of-hyperledger-fabric/45787041"
*/

func	get_creator(stub shim.ChaincodeStubInterface) (string, error) {
	// GetCreator returns marshaled serialized identity of the client
    serializedID, _ := stub.GetCreator()

    sId := &msp.SerializedIdentity{}
    err := proto.Unmarshal(serializedID, sId)
    if err != nil {
        return "", fmt.Errorf("Could not deserialize a SerializedIdentity, err %s", err)
    }

    bl, _ := pem.Decode(sId.IdBytes)
    if bl == nil {
        return "", fmt.Errorf("Failed to decode PEM structure")
    }
    cert, err := x509.ParseCertificate(bl.Bytes)
    if err != nil {
        return "", fmt.Errorf("Unable to parse certificate %s", err)
    }
	fmt.Println("----------------HAHAHAHAHAHA-----------------")
	fmt.Println(cert)
	fmt.Println(cert.Signature)
	fmt.Println(string(cert.Signature))
	fmt.Println(cert.PublicKey)
	return string(serializedID), nil
}

package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//Encode the given ecdsa key in pem format
func pem_encode_pubkey(publicKey *ecdsa.PublicKey) string {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	return string(pemEncodedPub)
}

//return the public key of creator in pem format
func getPemPublicKeyOfCreator(stub shim.ChaincodeStubInterface) (string, error) {

	cert, err := cid.GetX509Certificate(stub)
	if err != nil {
		return "", fmt.Errorf("Error : %s", err)
	}
	ecPublicKey := cert.PublicKey.(*ecdsa.PublicKey)
	//fmt.Println(ecPublicKey)
	//fmt.Printf("PUB : %x\n", ecdPublicKey)
	return pem_encode_pubkey(ecPublicKey), nil
}

//trim public key to remove newlines and begin/end tag
func trimPemPubKey(key string) string {
	key = strings.Replace(key, "\n", "", -1)
	key = strings.Replace(key, "-----BEGIN PUBLIC KEY-----", "", -1)
	key = strings.Replace(key, "-----END PUBLIC KEY-----", "", -1)
	return key
}

//return the trimmed public key of creator in pem format
func getTrimmedPubKeyOfCreator(stub shim.ChaincodeStubInterface) (string, error) {
	spender, er := getPemPublicKeyOfCreator(stub)
	if er != nil {
		return "", fmt.Errorf("Cannot get creator of the transaction : %s", er)
	}
	spender = trimPemPubKey(spender)
	return spender, nil

}

package main

import "fmt"
import "net/http"
import "os"
import "os/exec"

const	IP_ADDRESS = "192.168.1.58:8000"

func		ejbgekjrg(typeofTx string, id string, tx Request) string {
	var		base	string
	var		env		string
	var		command	string

	base = "docker exec alice bash -c "
	env = fmt.Sprintf("CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/MEDSOS.example.com/users/%s@MEDSOS.example.com/msp/ ", id)

	switch typeofTx {
		case "totalSupply":
			command = fmt.Sprintf("io totalSupply")
		case "balanceOf":
			command = fmt.Sprintf("io balanceOf %s", tx.Body["TokenOwner"])
		case "allowance":
			command = fmt.Sprintf("io allowance %s %s", tx.Body["TokenOwner"], tx.Body["Spender"])
		case "transfer":
			command = fmt.Sprintf("io transfer %s %s", tx.Body["To"], tx.Body["Tokens"])
		case "approve":
			command = fmt.Sprintf("io approve %s %s", tx.Body["Spender"], tx.Body["Tokens"])
		case "transferFrom":
			command = fmt.Sprintf("io transferFrom %s %s %s", tx.Body["From"], tx.Body["To"], tx.Body["Tokens"])
		case "publicKey":
			command = fmt.Sprintf("io publicKey --silent")
	}

	return base + "\"" + env + command + "\""
}
////////////////////////////////////////////////////////////////////////////////
///	PRIVATE	
////////////////////////////////////////////////////////////////////////////////

func	homepage(w http.ResponseWriter, req *http.Request) {
	var	tx		Request
	var command	string
	var b		[]byte
	var err		error

	if err = tx.Get(req); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	command = ejbgekjrg(tx.Body["Transaction"], tx.Body["Id"], tx)
	fmt.Printf("HOM [%s]\n", command)
	if b, err = exec.Command("bash", "-c", command).Output(); err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(string(b))
}

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC 
////////////////////////////////////////////////////////////////////////////////

func	main() {
	var err	error

	// Router
	http.HandleFunc("/", homepage)

	// Server
	if err = http.ListenAndServe(IP_ADDRESS, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

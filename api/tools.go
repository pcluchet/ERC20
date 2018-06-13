package main

import "fmt"
import "strings"
import "net"

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC 
////////////////////////////////////////////////////////////////////////////////

func	getIp() (string, error) {
	var addrs	[]net.Addr
	var ip		net.IP
	var err		error

	if addrs, err = net.InterfaceAddrs(); err != nil {
		return "", err
	}

	for _, value := range addrs {
		if ip, _, err = net.ParseCIDR(value.String()); err != nil {
			return "", err
		}
		if ip.IsLoopback() == false && ip.To4() != nil {
			return ip.String(), nil
		}
	}

	return "", fmt.Errorf("No ip found")
}

func parseStdout(stdout string) string {
	var index int

	index = strings.Index(stdout, "payload:")
	stdout = stdout[index+len("payload:\""):]
	stdout = strings.Split(stdout, "\n")[0]
	stdout = stdout[:(len(stdout) - 2)]
	stdout = strings.Replace(stdout, "\\", "", -1)

	return stdout
}

func		ejbgekjrg(typeofTx string, id string, tx Request) string {
	var		base	string
	var		env		string
	var		command	string

	base = "docker exec CLI bash -c "
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

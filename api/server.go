package main

import "fmt"
import "net/http"
import "os"
import "os/exec"

////////////////////////////////////////////////////////////////////////////////
///	PRIVATE
////////////////////////////////////////////////////////////////////////////////

func	homepage(w http.ResponseWriter, req *http.Request) {
	var err		error
	var	txType	string
	var	tx		Request
	var command	string
	var output	[]byte
	var	body	string

	fmt.Println("A")
	if err = tx.Get(req); err != nil {
		fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "500", err)
		return
	}
	fmt.Println(tx)

	fmt.Println("B")
	txType = tx.Body["Transaction"]
	command = ejbgekjrg(txType, tx.Body["Id"], tx)
	fmt.Printf("[%s]\n", command)
	if output, err = exec.Command("bash", "-c", command).Output(); err != nil {
		fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "500", err)
		return
	}

	fmt.Println("C")
	if txType == "publicKey" {
		body = parseStdoutForPubkey(string(output))
	} else {
		body = parseStdout(string(output))
	}

	fmt.Println("D")
	if txType == "listUsers" || txType == "whoOwesMe"{
		body, err = humanReadableKeys(body, txType)
		if err != nil {
			fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "500", err)
			return
		}
	}

	fmt.Println("E")
	fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "200", body)
}

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC
////////////////////////////////////////////////////////////////////////////////

func main() {
	var ip string
	var err error

	if ip, err = getIp(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	fmt.Printf("IP address: %s:8000\n", ip)

	http.HandleFunc("/", homepage)
	if err = http.ListenAndServe(ip+":8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

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

	if err = tx.Get(req); err != nil {
		fmt.Fprintln(w, err)
		return
	}

	txType = tx.Body["Transaction"]
	command = ejbgekjrg(txType, tx.Body["Id"], tx)
	if output, err = exec.Command("bash", "-c", command).Output(); err != nil {
		fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "500", err)
		return
	}

	body = parseStdout(string(output))

	if txType == "listUsers" || txType == "whoOwesMe"{
		body = humanReadableKeys(body, txType)
	}

	fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "200", body)
}

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC 
////////////////////////////////////////////////////////////////////////////////

func	main() {
	var	ip	string
	var err	error

	if ip, err = getIp(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	http.HandleFunc("/", homepage)
	if err = http.ListenAndServe(ip + ":8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

package main

import "fmt"
import "net/http"
import "os"
import "os/exec"

////////////////////////////////////////////////////////////////////////////////
///	PRIVATE
////////////////////////////////////////////////////////////////////////////////

func homepage(w http.ResponseWriter, req *http.Request) {
	var tx Request
	var command string
	var b []byte
	var err error

	if err = tx.Get(req); err != nil {
		fmt.Fprintln(w, err)
		return
	}

	command = ejbgekjrg(tx.Body["Transaction"], tx.Body["Id"], tx)
	if b, err = exec.Command("bash", "-c", command).Output(); err != nil {
		fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "500", err)
		return
	}

	if tx.Body["Transaction"] == "publicKey" {
		fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "200", parseStdoutForPubkey(string(b)))
	} else {
		fmt.Fprintf(w, "{\"result\":\"%s\",\"body\":\"%s\"}", "200", parseStdout(string(b)))
	}
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

	http.HandleFunc("/", homepage)
	if err = http.ListenAndServe(ip+":8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

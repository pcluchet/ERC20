package main

import "fmt"
import "net/http"
import "os"

const	IP_ADDRESS = "192.168.1.58:8000"

////////////////////////////////////////////////////////////////////////////////
///	PRIVATE	
////////////////////////////////////////////////////////////////////////////////

func	homepage(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
}

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC 
////////////////////////////////////////////////////////////////////////////////

func	main() {
	var err	error

	// Router
	http.HandleFunc("/", homepage)
	http.HandleFunc("/totalSupply", totalSupply)
	http.HandleFunc("/balanceOf", balanceOf)
	http.HandleFunc("/allowance", allowance)
	http.HandleFunc("/transfer", transfer)
	http.HandleFunc("/approve", approve)
	http.HandleFunc("/transferFrom", transferFrom)

	// Server
	if err = http.ListenAndServe(IP_ADDRESS, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

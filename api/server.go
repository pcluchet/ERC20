package main

import "fmt"
import "net/http"
import "os"

func	homepage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", req.RemoteAddr)
}

func	main() {
	var err	error

	// Router
	http.HandleFunc("/", homepage)

	// Server
	if err = http.ListenAndServe("192.168.1.58:8000", nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

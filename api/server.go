package main

import "fmt"
import "net/http"
import "os"
import "io/ioutil"
import "encoding/json"

const	IP_ADDRESS = "192.168.1.58:8000"

func (self *Request) Get(req *http.Request) {
	var b		[]byte
	var err		error

	if b, err = ioutil.ReadAll(req.Body); err != nil {
		fmt.Printf("ReadAll: %s\n", err)
	}
	if err = json.Unmarshal(b, &self.Body); err != nil {
		fmt.Printf("Unmarshal: %s\n", err)
	}
}

func	balanceOf(w http.ResponseWriter, req *http.Request) {
	var	tx	Request

	tx.Get(req)
	fmt.Println(tx.Body)
}

func	main() {
	var err	error

	// Router
	http.HandleFunc("/balanceOf", balanceOf)

	// Server
	if err = http.ListenAndServe(IP_ADDRESS, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

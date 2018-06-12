package main

import "fmt"
import "net/http"

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC 
////////////////////////////////////////////////////////////////////////////////

func	totalSupply(w http.ResponseWriter, req *http.Request) {
	var	tx	Request
	var err	error

	if err = tx.Get(req); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tx.Body)
}

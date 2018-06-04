package main

import "fmt"
import "os"
import "encoding/json"

const APPROVAL = "approvals"
const TRANSFER = "transfers"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func	Transfer(from string, to string, value uint64) {
	var err		error
	var ret		[]byte

	if ret, err = json.Marshal(Events{from, to, value}); err != nil {
		fmt.Fprintf(os.Stderr, "Error = %s\n", err)
	}
	if err = STUB.PutState(APPROVAL, ret); err != nil {
		fmt.Fprintf(os.Stderr, "Error = %s\n", err)
	}
}

func	Approval(owner string, spender string, value uint64) {
	var err		error
	var ret		[]byte

	if ret, err = json.Marshal(Events{owner, spender, value}); err != nil {
		fmt.Fprintf(os.Stderr, "Error = %s\n", err)
	}
	if err = STUB.PutState(TRANSFER, ret); err != nil {
		fmt.Fprintf(os.Stderr, "Error = %s\n", err)
	}
}

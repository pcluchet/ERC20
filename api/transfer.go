package	main

import	"strconv"
import	"fmt"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func	transfer(request Request) {
	var	err				error
	var	exist			bool
	var	from			string
	var	to				string
	var	amount_string	string
	var	amount			uint64

	from, exist = request.Header["Id"]
	if exist == false {
		fmt.Printf("error: cannot get user id.")
	}
	to, exist = request.Body["To"]
	if exist == false {
		fmt.Printf("error: cannot get user id.")
	}
	amount_string, exist = request.Body["Amount"]
	if exist == false {
		fmt.Printf("error: cannot get user id.")
	}
	amount, err = strconv.ParseUint(amount_string, 10, 64)
	if err != nil {
		fmt.Printf("error: cannot get user id.")
	}
	fmt.Printf("%s (%u)-> %s\n", from, to, amount)
}

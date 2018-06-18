package	main

import	"fmt"
import	"github.com/hyperledger/fabric/core/chaincode/shim"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func		loadUsers(iterator shim.StateQueryIteratorInterface) (string, error) {
	var		is_dev		bool
	var		users		string

	for iterator.HasNext() {
		result, iter_err := iterator.Next()
		if iter_err != nil {
			return "", fmt.Errorf("Cannot iter through users.")
		}
		_, is_dev = ledger_dev_keys[result.Key]
		if is_dev == true {
			continue
		}
		if users == "" {
			users = fmt.Sprintf("%s", result.Key)
		} else {
			users = fmt.Sprintf("%s, %s", users, result.Key)
		}
	}
	iterator.Close()
	return users, nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func		listUsers() (string, error) {
	var		err			error
	var		iterator	shim.StateQueryIteratorInterface

	iterator, err = STUB.GetStateByRange("", "")
	if err != nil {
		return "", fmt.Errorf("Cannot get keys iterator.")
	}
	return loadUsers(iterator)
}

package	main

import	"fmt"
import	"github.com/hyperledger/fabric/core/chaincode/shim"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func		loadApprovals(iterator shim.StateQueryIteratorInterface) (string, error) {
	var		approvals	string

	for iterator.HasNext() {
		result, iter_err := iterator.Next()
		if iter_err != nil {
			return "", fmt.Errorf("Cannot iter through users.")
		}
		if approvals == "" {
			approvals = fmt.Sprintf("%s:%s", result.Key, result.Value)
		} else {
			approvals = fmt.Sprintf("%s, %s:%s", approvals, result.Key, result.Value)
		}
	}
	iterator.Close()
	return approvals, nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func		whoOwesMe() (string, error) {
	var		err			error
	var		iterator	shim.StateQueryIteratorInterface

	iterator, err = STUB.GetQueryResult("{\"selector\":{\"owner\":\"tom\"}}")
	if err != nil {
		return "", fmt.Errorf("Cannot get query iterator.")
	}
	return loadApprovals(iterator)
}

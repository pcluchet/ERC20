package main

import "fmt"
import "strconv"


/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func	checkNil(argv []string) error {
	for index, _ := range argv {
		if argv[index] == "" {
			return fmt.Errorf("Parameters cannot be equal to nil")
		}
	}

	return nil
}

func		usage(typeofTx string) error {
	switch typeofTx {
		case "transfer":
			return fmt.Errorf("Transfer: [address to] [uint tokens] [publicKey]")
		case "transferFrom":
			return fmt.Errorf("TransferFrom: [address from] [address to] [uint tokens]")
	}

	return nil
}

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func	parseArgv(argv []string, typeofTx string) error {
	var err	error

	if len(argv) != 3 {
		return usage(typeofTx)
	}
	if err = checkNil(argv); err != nil {
		return err
	}
	if argv[0] == argv[2] {
		return fmt.Errorf("Cannot send money to yourself")
	}

	return nil
}

func	parseFund(str string, publicKey string) (uint64, error) {
	var amount uint64
	var user UserInfos
	var err	error

	if amount, err = strconv.ParseUint(str, 10, 64); err != nil {
		return 0, err
	}
	if amount == 0 {
		return 0, fmt.Errorf("Cannot send 0 value")
	}
	if user, err = getUserInfos(STUB, publicKey); err != nil {
		return 0, err
	}
	if amount > user.Amount {
		return 0, fmt.Errorf("Insufficent fund [%d > %d]", amount, user.Amount)
	}

	return amount, nil
}

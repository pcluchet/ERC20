package main

import "fmt"

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

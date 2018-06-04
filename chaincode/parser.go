package main

import "fmt"

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func	checkNil(argv []string) error {
	params := []string{"address to", "uint tokens", "publicKey"}

	for index, _ := range argv {
		if argv[index] == "" {
			return fmt.Errorf("[%s] cannot be equal to nil", params[index])
		}
	}

	return nil
}

func	parseArgv(argv []string, amount uint64) error {
	var err	error

	if len(argv) != 3 {
		return fmt.Errorf("Usage: [address to] [uint tokens] [publicKey]")
	}
	if err = checkNil(argv); err != nil {
		return err
	}
	if argv[0] == argv[2] {
		return fmt.Errorf("Cannot send money to yourself")
	}

	return nil
}

func	parseFund(amount uint64, publicKey string) error {
	var user UserInfos
	var err	error

	if amount == 0 {
		return fmt.Errorf("Cannot send 0 value")
	}
	if user, err = getUserInfos(STUB, publicKey); err != nil {
		return err
	}
	if amount > user.Amount {
		return fmt.Errorf("Insufficent fund [%d > %d]", amount, user.Amount)
	}

	return nil
}


/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func	parser(argv []string, amount uint64) error {
	var err	error

	if err = parseArgv(argv, amount); err != nil {
		return err
	}
	if err = parseFund(amount, argv[2]); err != nil {
		return err
	}

	return nil
}

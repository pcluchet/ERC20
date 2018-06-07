package main

import "fmt"
import "strconv"

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func	parseTransferFrom(argv []string) (uint64, error) {
	var amount uint64
	var user	UserInfos
	var err		error
	var prs		bool

	if user, err = getUserInfos(STUB, argv[0]); err != nil {
		return 0, err
	}
	if _, prs = user.Allowances[argv[1]]; prs == false {
		return 0, fmt.Errorf("[%s] has no right to transferFrom [%s]", argv[1], argv[0])
	}
	if amount, err = strconv.ParseUint(argv[2], 10, 64); err != nil {
		return 0, err
	}

	if amount == 0 {
		return 0, fmt.Errorf("Cannot send 0 value")
	}
	if amount > user.Allowances[argv[1]] {
		return 0, fmt.Errorf("Insufficent fund")
	}

	return amount, nil
}
/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func	transferFrom(argv []string) (string, error) {
	var amount	uint64
	var err		error

	if err = parseArgv(argv, "transferFrom"); err != nil {
		return "", err
	}
	if amount, err = parseTransferFrom(argv); err != nil {
		return "", err
	}
	if err = changeStateFrom(argv[0], argv[1], amount, _transferFrom); err != nil {
		return "", err
	}
	if err = changeStateTo(argv[1], amount); err != nil {
		return "", err
	}
	if err = event(argv[0], argv[1], amount, "transferFrom"); err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfull transaction from [%s] to [%s]", argv[0], argv[1]), nil
}

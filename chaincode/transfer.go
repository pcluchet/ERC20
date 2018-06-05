package main

import "fmt"
import "strconv"

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func	parseTransfer(str string, publicKey string) (uint64, error) {
	var amount uint64
	var user UserInfos
	var err	error

	if amount, err = strconv.ParseUint(str, 10, 64); err != nil {
		return 0, err
	}
	if amount == 0 {
		return 0, fmt.Errorf("Cannot send 0 value")
	}
	if user, err = getUserInfos(publicKey); err != nil {
		return 0, err
	}
	if amount > user.Amount {
		return 0, fmt.Errorf("Insufficent fund")
	}

	return amount, nil
}

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

// To do [Change authentification of argv[2] with getCreator->getPublicKey]

func transfer(argv []string) (string, error) {
	var err error
	var amount uint64

	if err = parseArgv(argv, "transfer"); err != nil {
		return "", err
	}
	if amount, err = parseTransfer(argv[1], argv[2]); err != nil {
		return "", err
	}
	if err = changeStateFrom(argv[2], argv[0], amount, _transfer); err != nil {
		return "", err
	}
	if err = changeStateTo(argv[0], amount); err != nil {
		return "", err
	}
	if err = event(argv[2], argv[0], amount, "transfer"); err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfull transaction from [%s] to [%s]", argv[2], argv[0]), nil
}

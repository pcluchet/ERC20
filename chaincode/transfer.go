package main

import "fmt"

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func changeStateTo(to string, amount uint64) error {
	var user UserInfos
	var err error

	if user, err = getUserInfos(STUB, to); err == nil {
		user.Amount += amount
	} else {
		user.Amount = amount
	}

	if err = user.Set(to); err != nil {
		return err
	}

	return nil
}

func changeStateFrom(from string, amount uint64) error {
	var user UserInfos
	var err error

	if user, err = getUserInfos(STUB, from); err != nil {
		return err
	}
	user.Amount -= amount

	if err = user.Set(from); err != nil {
		return err
	}

	return nil
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
	if amount, err = parseFund(argv[1], argv[2]); err != nil {
		return "", err
	}
	if err = changeStateFrom(argv[2], amount); err != nil {
		return "", err
	}
	if err = changeStateTo(argv[0], amount); err != nil {
		return "", err
	}
	if err = event(argv[2], argv[0], amount, "transfer"); err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfull transaction from [%s] to [%s]", argv[0], argv[2]), nil
}

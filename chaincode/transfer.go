package main

import (
	"fmt"
	"strconv"
)

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

func transfer(argv []string) (string, error) {
	var err error
	var amount uint64

	if len(argv) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a 3 (last one is user)")
	}

	if amount, err = strconv.ParseUint(argv[1], 10, 64); err != nil {
		return "", err
	}

	if err = parser(argv, amount); err != nil {
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

	return "", nil
}

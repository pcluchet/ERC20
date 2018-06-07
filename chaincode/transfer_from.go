package main

import "fmt"
import "strconv"

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func	(self Transaction) ParseTransferFrom() (Transaction, error) {
	var prs		bool
	var err		error

	if _, err = self.ParseTransfer(); err != nil {
		return Transaction{}, err
	}
	if _, prs = self.User.Allowances[self.To]; prs == false {
		return Transaction{}, fmt.Errorf("Permission Denied")
	}

	return self, nil
}

func	getTransferFrom(argv []string) (Transaction, error) {
	var amount		uint64
	var user		UserInfos
	var err			error

	if amount, err = strconv.ParseUint(argv[2], 10, 64); err != nil {
		return Transaction{}, err
	}
	if user, err = getUserInfos(argv[0]); err != nil {
		return Transaction{}, err
	}

	return (Transaction{argv[0], argv[1], amount, user}).ParseTransferFrom()
}

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func	transferFrom(argv []string) (string, error) {
	var tx		Transaction
	var err		error

	if err = parseArgv(argv, "transferFrom", 3); err != nil {
		return "", err
	}
	if tx, err = getTransferFrom(argv); err != nil {
		return "", err
	}

	if err = changeStateFrom(tx, _transferFrom); err != nil {
		return "", err
	}
	if err = changeStateTo(tx); err != nil {
		return "", err
	}
	if err = event(tx.From, tx.To, tx.Amount, "transferFrom"); err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfull transaction"), nil
}

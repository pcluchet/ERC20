package main

import	"fmt"
import	"strconv"

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func	(self Transaction) ParseApprove() error {
	if self.From == self.To {
		return fmt.Errorf("Cannot allow money to yourself")
	}
	if self.Amount > self.User.Amount {
		return fmt.Errorf("Insufficent fund")
	}

	return nil
}

func	getApprove(argv []string) (Transaction, error) {
	var publicKey	string
	var amount		uint64
	var user		UserInfos
	var err			error

	if publicKey, err = getPublicKey(); err != nil {
		return Transaction{}, err
	}
	if amount, err = strconv.ParseUint(argv[1], 10, 64); err != nil {
		return Transaction{}, err
	}
	if user, err = getUserInfos(publicKey); err != nil {
		return Transaction{}, err
	}

	if err = (Transaction{publicKey, argv[0], amount, user}).ParseApprove(); err != nil {
		return Transaction{}, err
	}
	return Transaction{publicKey, argv[0], amount, user}, nil

}

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func approve(argv []string) (string, error) {
	var tx		Transaction
	var err		error

	if err = parseArgv(argv, "approve", 2); err != nil {
		return "", err
	}
	if tx, err = getApprove(argv); err != nil {
		return "", err
	}

	if err = changeStateFrom(tx, _approve); err != nil {
		return "", err
	}
	if err = event(tx.From, tx.To, tx.Amount, "approval"); err != nil {
		return "", err
	}

	return fmt.Sprintf("Successfull approval"), nil
}

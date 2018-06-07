package main

import	"fmt"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func approve(argv []string) (string, error) {
	var tx		Transaction
	var err		error

	if err = parseArgv(argv, "approve", 2); err != nil {
		return "", err
	}
	if tx, err = getTransfer(argv); err != nil {
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

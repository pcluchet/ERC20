package main

import "fmt"

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

//function allowance(address tokenOwner, address spender)
//function approve(address spender, uint tokens)
//function transferFrom(address from, address to, uint tokens)
//transfer(address to, uint tokens)

func	transferFrom(argv []string) (string, error) {
	var user	UserInfos
	var err		error

	// ParseArgv
	if user, err = getUserInfos(STUB, argv[0]); err != nil {
		return "", err
	}
	if _, prs := user.Allowances[argv[1]]; prs == false {
		return "", fmt.Errorf("[%s] has no right to transferFrom [%s]", argv[1], argv[0])
	}
	// Check if value is <= thant allowed
	// Change State
	// Change allowed

	fmt.Println("OK")

	return "", nil
}

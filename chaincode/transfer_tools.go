package main

type TransferFCT func(ptr *UserInfos, to string, amount uint64)

/* ************************************************************************** */
/*		PRIVATE																  */
/* ************************************************************************** */

func	_transfer(ptr *UserInfos, to string, amount uint64) {
	(*ptr).Amount -= amount
}

func	_transferFrom(ptr *UserInfos, to string, amount uint64) {
	(*ptr).Amount -= amount
	(*ptr).Allowances[to] -= amount
}

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func changeStateTo(to string, amount uint64) error {
	var user UserInfos
	var err error

	if user, err = getUserInfos(STUB, to); err == nil {
		user.Amount += amount
	} else {
		user.Amount = amount
		user.Allowances = make(map[string]uint64)
	}

	if err = user.Set(to); err != nil {
		return err
	}

	return nil
}

func changeStateFrom(from string, to string, amount uint64, fct TransferFCT) error {
	var user UserInfos
	var err error

	if user, err = getUserInfos(STUB, from); err != nil {
		return err
	}
	fct(&user, to, amount)
	if err = user.Set(from); err != nil {
		return err
	}

	return nil
}

package main

import "encoding/json"

/* ************************************************************************** */
/*		PUBLIC																  */
/* ************************************************************************** */

func	event(from string, to string, value uint64, typeOfEvent string) error {
	var err		error
	var ret		[]byte

	if ret, err = json.Marshal(Events{from, to, value}); err != nil {
		return err
	}
	if err = STUB.PutState(typeOfEvent, ret); err != nil {
		return err
	}

	return nil
}

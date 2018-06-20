package	main

import	"fmt"
import	"os"
import	"io/ioutil"
import	"os/exec"
import	"regexp"
import	"encoding/json"

////////////////////////////////////////////////////////////////////////////////
/// STATIC FUNCTIONS
////////////////////////////////////////////////////////////////////////////////

func		getUserKey(path string) (string, error) {
	var		err		error
	var		command	string
	var		output	[]byte

	command = fmt.Sprintf("openssl ec -in \"%s\"* -pubout 2>&- | tail -n 3 | head -n 2 | tr -d '\n'", path)
	output, err = exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", fmt.Errorf("Cannot load user public key.")
	}
	if len(output) == 0 {
		return "", fmt.Errorf("Cannot load user public key.")
	}
	return string(output), nil
}

func		getUserName(path string) (string, error) {
	var		err		error
	var		reg		*regexp.Regexp
	var		name	[]byte

	reg, err = regexp.Compile("/[a-zA-Z0-9_]+@")
	if err != nil {
		return "", fmt.Errorf("Cannot load user public key: %s", err)
	}
	name = reg.Find([]byte(path))
	if len(name) == 0 {
		return "", fmt.Errorf("Cannot load user name.")
	}
	return string(name[1:len(name) - 1]), nil
}

func		loadUsers() (map[string]string, error) {
	var		err			error
	var		usersPath	string
	var		usersMap	map[string]string
	var		usersDir	[]os.FileInfo
	var		userKeyPath	string
	var		user		os.FileInfo
	var		userName	string
	var		userKey		string

	usersMap = make(map[string]string)
	usersPath = "../network/crypto-config/peerOrganizations/MEDSOS.example.com/users/"
	usersDir, err = ioutil.ReadDir(usersPath)
	if err != nil {
		return nil, fmt.Errorf("Cannot load users public key: %s", err)
	}
	for _, user = range usersDir {
		userKeyPath = usersPath + user.Name() + "/msp/keystore/"
		userKey, err = getUserKey(userKeyPath)
		if err != nil {
			return nil, err
		}
		userName, err = getUserName(userKeyPath)
		if err != nil {
			return nil, err
		}
		usersMap[userKey] = userName
	}
	return usersMap, nil
}

func		translateListUsers(output string, users map[string]string) (string, error) {
	var		err			error
	var		list		[]string
	var		index		int
	var		user		string
	var		userName	string
	var		isPresent	bool
	var		newOutput	[]byte

	err = json.Unmarshal([]byte(output), &list)
	if err != nil {
		return "", fmt.Errorf("Cannot get users list: %s", err)
	}
	for index, user = range list {
		userName, isPresent = users[user]
		if isPresent == true {
			list[index] = userName
		}
	}
	newOutput, err = json.Marshal(list)
	if err != nil {
		return "", fmt.Errorf("Cannot marshal users list: %s", err)
	}
	return string(newOutput), nil
}

func		translateWhoOwesMe(output string, users map[string]string) (string, error) {
	var		err			error
	var		allowances	map[string]uint64
	var		allowance	uint64
	var		userName	string
	var		user		string
	var		isPresent	bool
	var		newOutput	[]byte

	err = json.Unmarshal([]byte(output), &allowances)
	if err != nil {
		return "", fmt.Errorf("Cannot get allowances: %s", err)
	}
	for user, allowance = range allowances {
		userName, isPresent = users[user]
		if isPresent == true {
			delete(allowances, user)
			allowances[userName] = allowance
		}
	}
	newOutput, err = json.Marshal(allowances)
	if err != nil {
		return "", fmt.Errorf("Cannot marshal allowances: %s", err)
	}
	return string(newOutput), nil
}

////////////////////////////////////////////////////////////////////////////////
/// PUBLIC FUNCTION
////////////////////////////////////////////////////////////////////////////////

func		humanReadableKeys(output string, mode string) (string, error) {
	var		err		error
	var		users	map[string]string

	users, err = loadUsers()
	if err != nil {
		return "", fmt.Errorf("Cannot load users public key: %s", err)
	}
	fmt.Println(users)
	if mode == "listUsers" {
		return translateListUsers(output, users)
	} else if mode == "whoOwesMe" {
		return translateWhoOwesMe(output, users)
	}
	return "", fmt.Errorf("Unknown human readable translation mode [%s]", mode)
}

package main

import "fmt"
import "io/ioutil"
import "encoding/json"
import "net/http"
import "os/exec"
import "strings"

////////////////////////////////////////////////////////////////////////////////
///	PRIVATE
////////////////////////////////////////////////////////////////////////////////

func getPublicKey(tx Request, value string) string {
	var command string
	var b []byte
	var err error

	command = ejbgekjrg("publicKey", value, tx)
	if b, err = exec.Command("bash", "-c", command).Output(); err != nil {
		return ""
	}
	return strings.Trim(string(b), "\n")
}

func (self *Request) Public() error {
	var value string
	var prs bool
	var params = []string{"TokenOwner", "Spender", "From", "To"}

	for index, _ := range params {
		if value, prs = self.Body[params[index]]; prs == true {
			self.Body[params[index]] = getPublicKey(*self, value)
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC
////////////////////////////////////////////////////////////////////////////////

func (self *Request) Get(req *http.Request) error {
	var b []byte
	var err error

	if b, err = ioutil.ReadAll(req.Body); err != nil {
		return fmt.Errorf("ReadAll: %s", err)
	}
	if err = json.Unmarshal(b, &self.Body); err != nil {
		return fmt.Errorf("Unmarshal: %s", err)
	}

	return (*self).Public()
}

package main

import "fmt"
import "io/ioutil"
import "encoding/json"
import "net/http"

////////////////////////////////////////////////////////////////////////////////
///	PUBLIC 
////////////////////////////////////////////////////////////////////////////////

func (self *Request) Get(req *http.Request) error {
	var b		[]byte
	var err		error

	if b, err = ioutil.ReadAll(req.Body); err != nil {
		return fmt.Errorf("ReadAll: %s", err)
	}
	if err = json.Unmarshal(b, &self.Body); err != nil {
		return fmt.Errorf("Unmarshal: %s", err)
	}

	return nil
}

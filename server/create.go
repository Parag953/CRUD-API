package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func postUser(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newuser user
	json.Unmarshal(reqBody, &newuser)
	addUser(newuser.Name, newuser.Contact_no, newuser.DOB)

}

func addUser(name string, contact_no string, DOB string) error {

	res, err := db.Exec("insert into userinfo(name, contact_no , DOB) Values ( ? ,? ,? )", name, contact_no, DOB)
	if err != nil {
		return fmt.Errorf("can't add due to error: %s", err)
	}
	id, _ := res.LastInsertId()
	fmt.Println("Last inserted id is: ", id)

	return nil
}

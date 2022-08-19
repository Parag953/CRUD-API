package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var changedUser user
	json.Unmarshal(reqBody, &changedUser)

	err := UpdateUserById(Id, changedUser.Name, changedUser.Contact_no, changedUser.DOB)

	if err != nil {
		log.Fatal(err)
	}
}

func UpdateUserById(Id int, name string, contact_no string, DOB string) error {
	_, err := db.Exec("UPDATE userinfo SET name = ?, contact_no= ?, DOB= ? WHERE id = ?;", name, contact_no, DOB, Id)
	if err != nil {
		return fmt.Errorf("can't update due to error: %s", err)
	}

	return nil
}

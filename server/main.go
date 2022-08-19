package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type user struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Contact_no string `json:"contact_no"`
	DOB        string `json:"dob"`
}

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	deleteUserById(Id)

}
func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Users",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users/{id}", getUserById).Methods("GET")
	r.HandleFunc("/users", postUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")    // ----> To update a grocery
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE") // ----> Delete a grocery
	log.Fatal(http.ListenAndServe(":8080", r))

}

func UpdateUserById(Id int, name string, contact_no string, DOB string) error {
	_, err := db.Exec("UPDATE userinfo SET name = ?, contact_no= ?, DOB= ? WHERE id = ?;", name, contact_no, DOB, Id)
	if err != nil {
		return fmt.Errorf("can't update due to error: %s", err)
	}

	return nil
}

func deleteUserById(Id int) error {
	_, err := db.Exec("delete from userinfo where id=?", Id)
	if err != nil {
		return fmt.Errorf("can't delete due to error: %s", err)
	}
	fmt.Printf("User with Id=%d is deleted\n", Id)

	return nil
}

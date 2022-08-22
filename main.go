package main

import (
	"CRUD/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *sql.DB

// type user struct {
// 	Id         int    `json:"id"`
// 	Name       string `json:"name"`
// 	Contact_no string `json:"contact_no"`
// 	DOB        string `json:"dob"`
// }

func main() {
	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   "root",
	// 	Passwd: "",
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "Users",
	// }

	// // Get a database handle.
	// var err error
	// db, err = sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	// fmt.Println("Connected!")
	err := server.DbConnect()
	if err != nil {
		fmt.Println("Error while connecting to database. Error:", err)

	}
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users/{id}", server.GetUserById).Methods("GET")
	r.HandleFunc("/users", server.PostUser).Methods("POST")
	r.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")    // ----> To update a grocery
	r.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE") // ----> Delete a grocery
	log.Fatal(http.ListenAndServe(":8080", r))

}

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
// 	Contact_no string `json:"contactNo"`
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

	// read about all the http codes 2xx,3xx,4xx,5xx		:OK
	// difference between println and logger and all log levels
	// different types of request body in http post
	// soap vs rest
	// camelCase and snake_case
	err := server.DbConnect()
	if err != nil {
		fmt.Println("Error while connecting to database.", err)
	}
	log.Println("Database Connected")
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users/{id}", server.GetUserById).Methods("GET")
	r.HandleFunc("/users", server.PostUser).Methods("POST")
	r.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}

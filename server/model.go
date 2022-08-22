package server

import (
	"database/sql"
)

var db *sql.DB

type user struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Contact_no string `json:"contact_no"`
	DOB        string `json:"dob"`
}

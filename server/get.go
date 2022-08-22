package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	StrId := vars["id"]
	Id, er := strconv.Atoi(StrId)
	if er != nil {
		log.Fatal(er)
	}
	usr, err := UserById(Id)
	if err != nil {
		log.Fatal()
	}
	json.NewEncoder(w).Encode(usr)
}

func UserById(Id int) (*user, error) {
	// An albums slice to hold data from returned rows.
	var usr user

	row, err := db.Query("SELECT * FROM userinfo WHERE id = ?", Id)
	if err != nil {
		return nil, fmt.Errorf("userById %q: %v", Id, err)
	}
	defer row.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for row.Next() {
		if err := row.Scan(&usr.Id, &usr.Name, &usr.Contact_no, &usr.DOB); err != nil {
			return nil, fmt.Errorf("userById %q: %v", Id, err)
		}
	}
	return &usr, nil
}

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	StrId := params["id"]
	Id, _ := strconv.Atoi(StrId)

	usr, err := UserById(Id, w)
	if err != nil {
		return
	}
	if usr.Id == 0 {
		w.WriteHeader(404)
		return
	}
	json.NewEncoder(w).Encode(usr)
}

func UserById(Id int, w http.ResponseWriter) (*user, error) {
	var usr user
	row, err := db.Query("SELECT * FROM userinfo WHERE id = ?", Id)
	if err != nil {
		w.WriteHeader(500)
		return nil, fmt.Errorf("userById %q: %v", Id, err)
	}
	defer row.Close()
	for row.Next() {
		if err := row.Scan(&usr.Id, &usr.Name, &usr.Contact_no, &usr.DOB); err != nil {
			return nil, fmt.Errorf("userById %q: %v", Id, err)
		}
	}
	return &usr, nil
}

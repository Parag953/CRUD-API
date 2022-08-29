package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, _ := strconv.Atoi(vars["id"])
	deleteUserById(Id, w)

}

func deleteUserById(Id int, w http.ResponseWriter) error {

	_, err := db.Exec("delete from userinfo where id=?", Id)
	if err != nil {
		w.WriteHeader(500)
		return fmt.Errorf("can't delete due to error: %s", err)
	}
	fmt.Printf("User with Id=%d is deleted\n", Id)

	return nil
}

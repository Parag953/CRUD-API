package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostUser(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newuser user
	json.Unmarshal(reqBody, &newuser)
	//pingErr := db.Ping()
	// if pingErr != nil {
	// 	fmt.Println(pingErr)
	// 	w.WriteHeader(500)
	// 	return
	// }
	AddUser(newuser.Name, newuser.Contact_no, newuser.DOB, w)

}

func AddUser(name string, contact_no string, DOB string, w http.ResponseWriter) error {

	res, err := db.Exec("insert into userinfo(name, contactNo , dateOfBirth) Values ( ? ,? ,? )", name, contact_no, DOB)
	if err != nil {
		w.WriteHeader(500)
		return fmt.Errorf("can't add due to error: %s", err)
	}
	id, _ := res.LastInsertId()
	fmt.Println("Last inserted id is: ", id)

	return nil
}

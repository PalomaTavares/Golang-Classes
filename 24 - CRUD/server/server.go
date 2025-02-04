package server

import (
	db "crud/DB"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

// Create user on db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Failed to read request"))
	}

	var user user

	if error = json.Unmarshal(requestBody, &user); error != nil {
		w.Write([]byte("Error when converting user to struct"))
		return
	}
	dataBase, error := db.Connect()

	if error != nil {
		w.Write([]byte("Failed to connect to data base"))
		return
	}

	defer dataBase.Close()

	//PREPARE STATEMENT
	statement, error := dataBase.Prepare("insert into users (name, email) values (?, ?)")

	if error != nil {
		w.Write([]byte("Error on statement criation"))
		return
	}
	defer statement.Close()

	insert, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		w.Write([]byte("Error on statement execution"))
		return
	}

	idInserted, error := insert.LastInsertId()
	if error != nil {
		w.Write([]byte("Error on getting inserted id"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted successfully! id: %d", idInserted)))
}

// get all users on database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, error := db.Connect()
	if error != nil {
		w.Write([]byte("error on data base connection"))
		return
	}
	defer db.Close()

	lines, error := db.Query("select * from users")
	if error != nil {
		w.Write([]byte("Error when searching user"))
		return
	}

	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user

		if error := lines.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("error on user scan"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)

	if error := json.NewEncoder(w).Encode(users); error != nil {
		w.Write([]byte("Error on json conversion"))
	}
}

// get an especific user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error on conversion to int"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		w.Write([]byte("error on data base connection"))
		return
	}

	line, error := db.Query("select * from users where id = ?", ID)
	if error != nil {
		w.Write([]byte("error on user search"))
		return
	}

	var user user

	if line.Next() {
		if error := line.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("error on user scan"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(w).Encode(user); error != nil {
		w.Write([]byte("Error on json conversion"))
	}

	db.Close()
}

// update user
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("error on parameter convertion"))
		return
	}

	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("error while reading request body"))
		return
	}

	var user user
	if error := json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("error on converting user"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		w.Write([]byte("error on db connection"))
		return
	}

	defer db.Close()

	statement, error := db.Prepare("update users set name = ?, email = ? where id = ?")
	if error != nil {
		w.Write([]byte("error on statement creation"))
		return
	}

	defer statement.Close()

	if _, error := statement.Exec(user.Name, user.Email, ID); error != nil {
		w.Write([]byte("error on user update"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("error on covertion"))
		return
	}

	db, error := db.Connect()
	if error != nil {
		w.Write([]byte("error on connection to db"))
		return
	}

	defer db.Close()

	statement, error := db.Prepare("delete from users where id = ?")
	if error != nil {
		w.Write([]byte("error on statement creation"))
		return
	}

	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		w.Write([]byte("error on delete"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var userList []User

func main() {
	userList = append(userList, User{Id: 1, Name: "Balu"})
	userList = append(userList, User{Id: 2, Name: "Babu"})
    userList = append(userList, User{Id: 3, Name: "Naidu"})
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsersListHandler).Methods("GET")
	r.HandleFunc("/users/{userID}", GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users", CreateNewUserHandler).Methods("POST")
	r.HandleFunc("/users/{userID}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{userID}", DeleteUserHandler).Methods("DELETE")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))

}

// TODO: Create User
func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Fatal("Error creating user")
	}

	userList = append(userList, newUser)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
}

// TODO: Get User by ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(string(vars["userID"]), 0, 64)

	user, _ := json.Marshal(getUserByID(userList, int(userID)))

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(user)
}

// TODO: Get all Users
func GetUsersListHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(userList)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// TODO: Update User details
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(string(vars["userID"]), 0, 64)

	type UpdateUserRequest struct {
		Name string `json:"name"`
	}

	var newUserDetails UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&newUserDetails)
	if err != nil {
		log.Fatal("Error creating user")
	}

	UpdateUserDetails(&userList, int(userID), newUserDetails.Name)

	u, _ := json.Marshal(getUserByID(userList,int(userID)))
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(u)
}

// TODO: Delete a User
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.ParseInt(string(vars["userID"]), 0, 64)
	var temp []User
	for _, u := range userList {
		if u.Id != int(userID) {
			temp = append(temp, u)
		}
	}
	userList = temp
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:5004")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/user", CreateUser).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/users", GetAllUser).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{id}", GetUserById).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{id}", UpdateUser).Methods("OPTIONS", "PUT")
	myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("OPTIONS", "Delete")

	log.Fatal(http.ListenAndServe(":5004", myRouter))
}
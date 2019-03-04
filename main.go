package main

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get-employees", getEmployees).Methods("GET")
	router.HandleFunc("/post-employee", postEmployee).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 6969")
	log.Fatal(http.ListenAndServe(":6969", router))
}


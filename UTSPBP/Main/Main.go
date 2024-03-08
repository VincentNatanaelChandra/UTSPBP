package main

import (
	"UTSPBP/Controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getrooms", Controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/insertroom", Controller.InsertRoom).Methods("POST")
	router.HandleFunc("/detailrooms", Controller.GetDetailRooms).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" //menangani permintaan http yang masuk
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "root"
	dbName   = "articleCrud"
)

func main() {
	// membuat router baru
	r := mux.NewRouter()

	// mendefinisikan route http dengan router
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	// menjalankan server pada port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))

}
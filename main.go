package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux" 
	"github.com/Alifanoveliasukma/crud-article-simple/handlers"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "root"
	dbName   = "articleCrud"
)

func main() {

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal("Failed connection to database", err)
	}
	defer db.Close()
	// membuat router baru
	r := mux.NewRouter()

	// mendefinisikan route http dengan router
	r.HandleFunc("/article", handlers.CreateArticleHandler(db)).Methods("POST")
	r.HandleFunc("/article/{id}", handlers.GetArticleHandler(db)).Methods("GET")
	r.HandleFunc("/article/{id}", handlers.UpdateArticleHandler(db)).Methods("PUT")
	r.HandleFunc("/article/{id}", handlers.DeleteArticleHandler(db)).Methods("DELETE")

	// menjalankan server pada port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))

}
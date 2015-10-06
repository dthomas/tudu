package main

import (
	"log"
	"net/http"

	"bitbucket.org/derick/tudu/api"
	"bitbucket.org/derick/tudu/datastore"
)

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	datastore.DB, err = sqlx.Connect("postgres", "user=dthomas password=test1234 dbname=tudu_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer datastore.DB.Close()

	m := http.NewServeMux()
	m.Handle("/api/", http.StripPrefix("/api", api.Handler()))
	log.Println("Listening on PORT 3000")

	err = http.ListenAndServe(":3000", m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

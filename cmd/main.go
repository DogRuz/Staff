package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	// server settings
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	flag.Parse()
	connStr := "host=localhost user=postgres dbname=Employees sslmode=disable"
	db, err := Open(connStr)
	env := Env{db: db}
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/get_information_people", env.GetInformationPeople)
	log.Fatal(http.ListenAndServe(*addr, mux))
}


package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"net/http"
)

func (env *Env)GetInformationPeople(w http.ResponseWriter, r *http.Request) {
	// function for getting information about an employee
	if r.Method != http.MethodGet {w.WriteHeader(405); w.Write([]byte("Метод запрещен!")); return}
	id := r.URL.Query().Get("id")
	stmt, err := env.db.Prepare("select * from People where id = $1")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []Working
	for rows.Next() {
		p := Working{}
		err := rows.Scan(&p.Id, &p.Name,(*pq.StringArray)(&p.Languages), &p.Salary)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
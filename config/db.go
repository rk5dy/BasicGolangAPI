package config

import (
  "fmt"
   _ "github.com/lib/pq"
  "database/sql"
)

var DB *sql.DB

func init() {
  var err error

  DB, err = sql.Open("postgres", "postgres://admin:sample@postgres/firecontrol?sslmode=disable")
	if err != nil {
		panic(err)
	}
  
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

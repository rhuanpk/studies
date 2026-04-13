package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type tb struct {
	ID *int    `json:"id"`
	CL *string `json:"cl"`
	// CL sql.NullString `json:"cl"`
}

func main() {
	db, err := sql.Open(
		"mysql",
		"root:root@tcp(172.17.0.2:3306)/db?tls=skip-verify",
	)
	defer db.Close()
	if err != nil {
		log.Fatalln("error in sql open connection:", err)
	}

	db.SetConnMaxLifetime(1)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	if err := db.Ping(); err != nil {
		log.Fatalln("error in sql ping connection:", err)
	}

	data, err := db.Query(`select * from tb`)
	defer data.Close()
	if err != nil {
		log.Fatalln("error in sql query rows:", err)
	}

	var rows []*tb
	for data.Next() {
		// row := &tb{}
		row := new(tb)
		err := data.Scan(&row.ID, &row.CL)
		if err != nil {
			log.Println("error in sql scan row:", err)
		}
		rows = append(rows, row)
	}

	indented, err := json.MarshalIndent(rows, "", "\t")
	if err != nil {
		log.Fatalln("error in json marshal indent:", err)
	}
	println(string(indented))
}

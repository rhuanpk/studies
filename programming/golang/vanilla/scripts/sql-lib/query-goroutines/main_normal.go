package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const max = 1000

type tb struct {
	ID *int    `json:"id"`
	CL *string `json:"cl"`
}

func main() {
	defer func(start time.Time) {
		println("duration:", time.Since(start).String())
	}(time.Now())

	db, err := sql.Open(
		"mysql",
		"root:root@tcp(172.17.0.2:3306)/db?tls=skip-verify",
	)
	defer db.Close()
	if err != nil {
		log.Fatalln("error in sql open connection:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("error in sql ping connection:", err)
	}

	var rows []*tb
	for index := 0; index < max; index++ {
		data, err := db.Query(`select * from tb`)
		// defer data.Close()
		if err != nil {
			log.Fatalln("error in sql query rows:", err)
		}

		if index > (max - 2) {
			for data.Next() {
				row := new(tb)
				err := data.Scan(&row.ID, &row.CL)
				if err != nil {
					log.Println("error in sql scan row:", err)
				}
				rows = append(rows, row)
			}
		}

		// case this data.Close() are not here change the db.SetMaxOpenConns() value
		data.Close()
		println("finish -", index)
	}

	indented, err := json.MarshalIndent(rows, "", "\t")
	if err != nil {
		log.Fatalln("error in json marshal indent:", err)
	}
	println(string(indented))
}

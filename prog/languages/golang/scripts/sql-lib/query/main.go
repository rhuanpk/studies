package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type tb struct {
	id int
	cl sql.NullString
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

	var rows []tb
	for data.Next() {
		var row tb
		err := data.Scan(&row.id, &row.cl)
		if err != nil {
			log.Println("error in sql scan row:", err)
		}
		rows = append(rows, row)
	}

	for index, value := range rows {
		fmt.Printf(
			"index [%d]\n\t- id: %d\n\t- cl: %s\n",
			index,
			value.id,
			map[bool]any{
				true:  value.cl.String,
				false: "nil",
			}[value.cl.Valid],
		)
	}
	//fmt.Printf("%#v\n", rows)
}

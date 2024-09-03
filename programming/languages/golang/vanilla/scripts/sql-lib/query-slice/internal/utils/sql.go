package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// struct to represent database table
type tb struct {
	id int
	cl sql.NullString
}

// Query exec queries in database.
func Query(db *sql.DB, query string, args ...any) {
	// query data
	println("executing query")
	data, err := db.Query(query, args...)
	defer data.Close()
	if err != nil {
		var severity string
		if errors.Is(sql.ErrNoRows, err) {
			severity = "warning"
		} else {
			severity = "error"
		}
		log.Println(severity+" in db query rows:", err)
	}

	// iterate under data to scan it into a manipulable object
	var rows []tb
	for data.Next() {
		var row tb
		err := data.Scan(&row.id, &row.cl)
		if err != nil {
			log.Println("error in data scan row:", err)
		}
		rows = append(rows, row)
	}

	// iterate under rows to print they
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

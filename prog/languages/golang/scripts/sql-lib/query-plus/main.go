package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"back/pkg/utils"

	// [file:]sqlite.db[?cache=shared[&mode=memory]]
	_ "github.com/mattn/go-sqlite3"
)

// global db instance
var db *sql.DB

// struct to represent database table
type tb struct {
	id int
	cl sql.NullString
}

func init() {
	// necessary so that the db variable is not initialized with local scope
	var err error

	// open and validade the connection
	db, err = sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatalln("error in sql open connection:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalln("error in sql ping connection:", err)
	}

	// set connection parameters
	db.SetConnMaxLifetime(1)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	// create table
	_, err = db.Exec(`
		create table if not exists tb (
			id integer primary key autoincrement,
			cl text null
		)
	`)
	if err != nil {
		log.Fatalln("error in db create table:", err)
	}

	// insert values
	result, err := db.Exec(`
		insert into tb (cl)
		values
			('value 1'),
			(null)
	`)
	if err != nil {
		log.Fatalln("error in db insert values:", err)
	}
	utils.LastInsertedID(result)
	utils.AffectedRows(result)

	// insert with statement
	stmt, err := db.Prepare(`
		insert into tb (cl)
		values (?)
	`)
	if err != nil {
		log.Fatalln("error in stmt create:", err)
	}

	for _, value := range []any{"value 3", "value 4"} {
		if result, err = stmt.Exec(value); err != nil {
			log.Fatalln("error in stmt exec:", err)
		}
		utils.LastInsertedID(result)
		utils.AffectedRows(result)
	}
}

func main() {
	defer db.Close()
	println("database/sql")

	// query column names
	data, err := db.Query(`select * from tb limit 0`)
	defer data.Close()
	if err != nil {
		log.Fatalln("error in db query rows:", err)
	}

	columns, err := data.Columns()
	if err != nil {
		log.Fatalln("error in db get column names:", err)
	}
	fmt.Println("columns:", columns)

	// query data
	data, err = db.Query(`select * from tb`)
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

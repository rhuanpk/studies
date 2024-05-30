package main

import (
	"database/sql"
	"fmt"
	"log"

	internalUtils "dev/internal/utils"
	pkgUtils "dev/pkg/utils"

	// [file:]sqlite.db[?cache=shared[&mode=memory]]
	_ "github.com/mattn/go-sqlite3"
)

// global db instance
var db *sql.DB

func init() {
	// necessary so that the db variable is not initialized with local scope
	var err error

	// default message struct
	message := &pkgUtils.Message{
		Replace: pkgUtils.Merge,
		Before:  "\t- ",
	}

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
	println("creating table")
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
	println("inserting values")
	result, err := db.Exec(`
		insert into tb (cl)
		values
			('value 1'),
			(null)
	`)
	if err != nil {
		log.Fatalln("error in db insert values:", err)
	}

	pkgUtils.LastInsertedID(result, message)
	pkgUtils.AffectedRows(result, message)

	// insert with statement
	println("inserting values with statement")
	stmt, err := db.Prepare(`
		insert into tb (cl)
		values (?)
	`)
	if err != nil {
		log.Fatalln("error in stmt create:", err)
	}

	for index, value := range []any{"value 3", "value 4"} {
		println("exec", index)
		if result, err = stmt.Exec(value); err != nil {
			log.Fatalln("error in stmt exec:", err)
		}
		pkgUtils.LastInsertedID(result, message)
		pkgUtils.AffectedRows(result, message)
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

	// exec queries
	args := []any{2, 3}
	internalUtils.Query(db, `select * from tb`)
	internalUtils.Query(db, `select * from tb where id in (2,3)`)
	internalUtils.Query(db, `select * from tb where id in (?, ?)`, args...)
	internalUtils.Query(db, `select * from tb where id in ($1, $2)`, args...)
	internalUtils.Query(db, pkgUtils.BuildQuerySlice(`select * from tb where id in (???)`, args...), args...)
}

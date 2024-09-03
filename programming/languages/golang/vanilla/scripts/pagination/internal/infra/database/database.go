package database

import (
	"database/sql"
	"fmt"
	"log"

	// [file:]sqlite.db[?cache=shared[&mode=memory]]
	_ "github.com/mattn/go-sqlite3"
)

// TB represent database table.
type TB struct {
	ID *int            `json:"id,omitempty"`
	CL *sql.NullString `json:"cl,omitempty"`
}

// DB global database instance.
var DB *sql.DB

func init() {
	// necessary so that the db variable is not initialized with local scope
	var err error

	// open and validade the connection
	DB, err = sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatalln("error in sql open connection:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalln("error in sql ping connection:", err)
	}

	// create table
	_, err = DB.Exec(`
		create table if not exists tb (
			id integer primary key autoincrement,
			cl text null
		)
	`)
	if err != nil {
		log.Fatalln("error in db create table:", err)
	}

	// insert with statement
	stmt, err := DB.Prepare(`
		insert into tb (cl)
		values (?)
	`)
	if err != nil {
		log.Fatalln("error in stmt create:", err)
	}

	for index := 1; index <= 10; index++ {
		if _, err = stmt.Exec(fmt.Sprintf("value %d", index)); err != nil {
			log.Fatalln("error in stmt exec:", err)
		}
	}
}

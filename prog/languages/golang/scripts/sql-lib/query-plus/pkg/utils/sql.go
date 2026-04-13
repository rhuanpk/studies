package utils

import (
	"database/sql"
	"log"
	"strings"
)

// LastInsertedID has default message "last inserted id:".
func LastInsertedID(result sql.Result, message ...string) {
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		log.Fatalln("error in db get last inserted id:", err)
	}
	// println("last inserted id:", lastInsertedID)
	if len(message) < 1 {
		message = []string{"last inserted id:"}
	}
	println(strings.Join(message, " "), lastInsertedID)
}

// AffectedRows has default message "affected rows:".
func AffectedRows(result sql.Result, message ...string) {
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("error in db get affected rows:", err)
	}
	if len(message) < 1 {
		message = []string{"affected rows:"}
	}
	println(strings.Join(message, " "), affectedRows)
}

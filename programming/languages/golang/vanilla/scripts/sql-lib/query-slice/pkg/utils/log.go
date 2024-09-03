package utils

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

type replace string

// const of replace type
const (
	Merge  replace = "merge"
	Printf replace = "printf"
	Full   replace = "full"
)

// Message struct for messages.
type Message struct {
	standard []any
	Before   any
	After    any
	Format   string
	Args     []any
	Full     []any
	Replace  replace
}

func (m Message) String() string {
	switch m.Replace {
	case Merge:
		var args []any
		if m.Before != nil {
			args = append(args, m.Before)
		}
		args = append(args, m.standard...)
		if m.After != nil {
			args = append(args, m.After)
		}
		return fmt.Sprint(args...)
	case Printf:
		return fmt.Sprintf(m.Format, m.Args...)
	case Full:
		return fmt.Sprint(m.Full...)
	default:
		return fmt.Sprint(m.standard...)
	}
}

// LastInsertedID has default message "last inserted id:".
func LastInsertedID(result sql.Result, message *Message) {
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		log.Println("error in db get last inserted id:", err)
	}
	if reflect.ValueOf(message).IsZero() {
		message = new(Message)
	}
	message.standard = []any{"last inserted id: ", lastInsertedID}
	fmt.Println(message)
}

// AffectedRows has default message "affected rows:".
func AffectedRows(result sql.Result, message *Message) {
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error in db get affected rows:", err)
	}
	if reflect.ValueOf(message).IsZero() {
		message = new(Message)
	}
	message.standard = []any{"affected rows: ", affectedRows}
	fmt.Println(message)
}

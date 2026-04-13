package logger

import (
	"log"
	"os"
	"strings"
	"time"
)

// Logger is the logger struct.
type Logger struct {
	log    *log.Logger
	place  string
	Prefix string
}

// Print prints with standard log.Print().
func (l Logger) Print(args ...any) {
	l.log.Print(args...)
}

// Println prints with standard log.Println().
func (l Logger) Println(args ...any) {
	l.log.Println(args...)
}

// Printf prints with standard log.Printf().
func (l Logger) Printf(format string, args ...any) {
	l.log.Printf(format, args...)
}

// Fatalln prints with standard log.Fatalln().
func (l Logger) Fatalln(args ...any) {
	l.log.Fatalln(args...)
}

// Short prints a default message with itself defined place or passa a label to overwrite once.
func (l Logger) Short(label ...string) {
	var message string
	if len(label) < 1 {
		message = string(l.place)
	} else {
		message = label[0]
	}
	println("> " + strings.ToUpper(message) + "!")
}

// FullPrefix return manualy the full prefix of logger.
func (l Logger) FullPrefix() string {
	return l.Prefix + time.Now().Format("15:04:05") + " "
}

// NewLogger return a pointer to Logger, "place" param is the function where you call this one.
func NewLogger(place string) *Logger {
	logger := &Logger{
		place:  place,
		Prefix: "[" + place + "] ",
	}
	logger.log = log.New(
		os.Stderr,
		logger.Prefix,
		log.Ltime,
	)
	return logger
}

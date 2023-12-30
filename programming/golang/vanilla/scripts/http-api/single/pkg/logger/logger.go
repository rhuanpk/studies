package logger

import (
	"log"
	"os"
	"strings"
	"time"
)

type logger string

// Logger constants.
const (
	Main   logger = "main"
	Client logger = "client"
	Server logger = "server"
)

// Logger is the logger struct.
type Logger struct {
	log    *log.Logger
	place  logger
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

// Short prints a default message with itself defined place.
func (l Logger) Short() {
	println("> " + strings.ToUpper(string(l.place)) + "!")
}

// FullPrefix return manualy the full prefix of logger.
func (l Logger) FullPrefix() string {
	return l.Prefix + time.Now().Format("15:04:05") + " "
}

// NewLogger return a pointer to Logger
func NewLogger(place logger) *Logger {
	logger := &Logger{
		Prefix: "[" + string(place) + "] ",
	}
	logger.log = log.New(
		os.Stderr,
		logger.Prefix,
		log.Ltime,
	)
	logger.place = place
	return logger
}

package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/google/uuid"
)

// go playground: https://go.dev/play/p/FvFuQan3BpC
func main() {
	headers := []string{"uuid", "name", "yers_old", "email"}
	values := [][]string{
		{uuid.NewString(), "foo bar", "24", "foobar@email.any"},
		{uuid.NewString(), "boo baz", "42", "boobaz@email.any"},
	}
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(writer, strings.Join(headers, "\t"))
	for _, value := range values {
		fmt.Fprintln(writer, strings.Join(value, "\t"))
	}
	writer.Flush()
}

package structure

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

const defaultSpaces = "    "

func pretty(structures ...any) {
	for _, structure := range structures {
		structJSON, err := json.MarshalIndent(structure, "", defaultSpaces)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf(
			"%s: %s\n",
			reflect.TypeOf(structure).Elem().Name(),
			string(structJSON),
		)
	}
}

package structure

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func pretty(structures ...any) {
	for _, structure := range structures {
		structJSON, err := json.MarshalIndent(structure, "", defaultSpace)
		if err != nil {
			log.Println("error in json marshal indent:", err)
		}
		fmt.Printf(
			"%s: %s\n",
			reflect.TypeOf(structure).Elem().Name(),
			string(structJSON),
		)
	}
}

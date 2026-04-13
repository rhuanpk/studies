package structure

import (
	"fmt"
	"reflect"
)

// Assign assigns the fields of source (src) into destiny (dst).
func Assign(src, dst any) {
	pretty(dst, src)
	println("\n------------------------------\n")

	reflectValueDST := reflect.ValueOf(dst).Elem()
	reflectValueSRC := reflect.ValueOf(src).Elem()
	for index := 0; index < reflectValueDST.NumField(); index++ {
		assignField(&reflectValueSRC, &reflectValueDST, index)
	}

	println("\n------------------------------\n")
	pretty(dst, src)
}

func assignField(valueSRC, valueDST *reflect.Value, index int) {
	valueDSTFieldName := valueDST.Type().Field(index).Name

	fmt.Printf("index %d)\n", index)
	print(defaultSpace+"* dst field: ", valueDSTFieldName)

	if valueDSTField := valueDST.Field(index); valueDSTField.CanSet() {
		valueSRCField := valueSRC.FieldByName(valueDSTFieldName)
		// defer func() {
		// 	if recover() != nil {
		// 		println(": no such field in source")
		// 	}
		// }()
		// if valueSRCField.IsZero() {
		// 		println(": no such field in source")
		// 	return
		// }

		valueSRCFieldName, ok := valueSRC.Type().FieldByName(valueDSTFieldName)
		if !ok {
			println(": no such field in source")
			return
		}

		fmt.Printf(" (%s)\n", valueDSTField.Kind().String())
		fmt.Printf(
			defaultSpace+"* src field: %s (%s)\n",
			valueSRCFieldName.Name,
			valueSRCField.Kind().String(),
		)
		println(defaultSpace + "- can set: YES")

		defer func() {
			if recover() != nil {
				var msg string
				if valueDSTFieldType := valueDSTField.Type(); valueSRCField.CanConvert(valueDSTFieldType) {
					msg = "YES"
					valueDSTField.Set(valueSRCField.Convert(valueDSTFieldType))
				} else {
					msg = "NO"
				}
				println(defaultSpace+"- can convert:", msg)
			}
		}()

		valueDSTField.Set(valueSRCField.Elem())
	} else {
		println("\n" + defaultSpace + "- can set: NO")
	}
}

package structure

import (
	"fmt"
	"reflect"
)

// Assign assign the fields of dst into src.
func Assign(dst, src any) {
	pretty(dst, src)
	println("\n------------------------------\n")

	reflectValueDST := reflect.ValueOf(dst).Elem()
	reflectValueSRC := reflect.ValueOf(src).Elem()
	for index := 0; index < reflectValueDST.NumField(); index++ {
		assignField(&reflectValueDST, &reflectValueSRC, index)
	}

	println("\n------------------------------\n")
	pretty(dst, src)
}

func assignField(dst, src *reflect.Value, index int) {
	reflectValueDSTField := dst.Field(index)
	fmt.Printf("index %d)\n", index)
	if reflectValueDSTField.CanSet() {
		println(defaultSpaces + "- can set")
		reflectValueDSTFieldType := reflectValueDSTField.Type()
		reflectValueDSTFieldName := dst.Type().Field(index).Name
		reflectValueSRCField := src.FieldByName(reflectValueDSTFieldName)
		defer func() {
			if recover() != nil {
				if reflectValueSRCField.CanConvert(reflectValueDSTFieldType) {
					println(defaultSpaces + "- can convert")
					reflectValueDSTField.Set(reflectValueSRCField.Convert(reflectValueDSTFieldType))
				} else {
					fmt.Printf(
						defaultSpaces+"- can't convert '%s' -> '%s' (%s)!\n",
						reflectValueSRCField.Kind().String(),
						reflectValueDSTField.Kind().String(),
						reflectValueDSTFieldName,
					)
				}
			}
		}()
		reflectValueDSTField.Set(reflectValueSRCField.Elem())
	} else {
		println(defaultSpaces + "- can't set")
	}
}

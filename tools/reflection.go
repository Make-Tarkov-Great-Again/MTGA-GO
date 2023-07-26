package tools

import "reflect"

func GetStructField(s interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(s)
	curStruct := pointToStruct.Elem()

	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}

	curField := curStruct.FieldByName(fieldName)
	if !curField.IsValid() {
		panic("not found: " + fieldName)
	}
	return curField
}

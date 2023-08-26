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

func InitializeFields(elem reflect.Value, jsonMap map[string]interface{}) {
	elemType := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldType := elemType.Field(i)

		jsonTag := fieldType.Tag.Get("json")

		if value, ok := jsonMap[jsonTag]; ok {
			if value == nil {
				// Handle null value
				continue
			}

			if field.CanSet() {
				switch field.Kind() {
				case reflect.Slice:
					// Handle slice assignment
					if valSlice, ok := value.([]interface{}); ok {
						// Convert []interface{} to []string
						stringSlice := make([]string, len(valSlice))
						for i, v := range valSlice {
							stringSlice[i] = v.(string)
						}
						field.Set(reflect.ValueOf(stringSlice))
					}
				case reflect.Struct:
					// Handle nested struct assignment
					if valMap, ok := value.(map[string]interface{}); ok {
						InitializeFields(field, valMap)
					}
				default:
					// Handle other basic types
					field.Set(reflect.ValueOf(value).Convert(field.Type()))
				}
			}
		}
	}
}

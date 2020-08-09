package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func StructToStringMap(v interface{}) *map[string]string {
	structValue := reflect.ValueOf(v)

	if structValue.IsNil() {
		return nil
	}

	stringMap := make(map[string]string, 0)
	structValue = structValue.Elem()

	for i := 0; i < structValue.NumField(); i++ {
		str := fieldToStringPointer(structValue.Field(i))

		if str != nil {
			stringMap[structValue.Type().Field(i).Name] = *str
		}
	}

	return &stringMap
}

func fieldToStringPointer(value reflect.Value) *string {
	var v interface{}

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil
		}

		v = value.Elem().Interface()
	} else {
		v = value.Interface()
	}

	switch vType := v.(type) {
	case string:
		return String(v.(string))
	case int:
		return String(strconv.Itoa(v.(int)))
	case bool:
		return String(strconv.FormatBool(v.(bool)))
	case time.Time:
		return String(v.(time.Time).Format(time.RFC3339))
	default:
		fmt.Printf("Type %v is not supported\n", vType)
		return nil
	}
}

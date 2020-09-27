package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

type mappedURLValue struct {
	queryParamSuffix *string
	value            *string
}

func StructToURLValues(v interface{}) *url.Values {
	structValue := reflect.ValueOf(v)

	if structValue.IsNil() {
		return nil
	}

	values := url.Values{}
	structValue = structValue.Elem()

	for i := 0; i < structValue.NumField(); i++ {
		slice := fieldToMappedURLValueSlice(structValue.Field(i))

		for _, urlValue := range slice {
			if urlValue.value != nil {
				name := structValue.Type().Field(i).Name
				if urlValue.queryParamSuffix != nil {
					name = name + "." + *urlValue.queryParamSuffix
				}
				values.Add(name, *urlValue.value)
			}
		}
	}

	return &values
}

func fieldToMappedURLValueSlice(value reflect.Value) []mappedURLValue {
	var v reflect.Value

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil
		}

		v = value.Elem()
	} else {
		v = value
	}

	switch v.Kind() {
	case reflect.Slice:
		slice := make([]mappedURLValue, 0)
		for i := 0; i < v.Len(); i++ {
			slice = append(slice, mappedURLValue{value: fieldToStringPointer(v.Index(i).Interface())})
		}
		return slice
	case reflect.Map:
		slice := make([]mappedURLValue, 0)
		for _, k := range v.MapKeys() {
			slice = append(slice, mappedURLValue{queryParamSuffix: String(k.Interface().(string)), value: fieldToStringPointer(v.MapIndex(k).Interface())})
		}
		return slice
	default:
		return []mappedURLValue{{value: fieldToStringPointer(v.Interface())}}
	}
}

func fieldToStringPointer(v interface{}) *string {
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

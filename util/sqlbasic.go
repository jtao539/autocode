package util

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// Entity2DTO 将字段类型为sql.NullXXX的实体 转换为字段类型为普通类型的DTO，转换过程以DTO为主。 tags 为需要跳过转换的字段, 常用于某个字段Entity和Dto字段类型不一致
func Entity2DTO(a interface{}, b interface{}, tags ...string) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	tb := reflect.TypeOf(b)
	ovb := vb.Elem()
	otb := tb.Elem()
	for i := 0; i < ovb.NumField(); i++ {
		field := otb.Field(i)
		if tag := field.Tag.Get("json"); len(tags) > 0 && containArray(tag, tags) {
			continue
		}
		switch ovb.Field(i).Kind() {
		case reflect.String:
			sValue := va.FieldByName(field.Name)
			if sValue.Kind() == reflect.Struct {
				for j := 0; j < sValue.NumField(); j++ {
					if sValue.Field(j).Kind() == reflect.String {
						ovb.FieldByName(field.Name).SetString(sValue.Field(j).String())
					}
				}
			}
		case reflect.Int:
			sValue := va.FieldByName(field.Name)
			if sValue.Kind() == reflect.Struct {
				for j := 0; j < sValue.NumField(); j++ {
					if sValue.Field(j).Kind() == reflect.Int32 || sValue.Field(j).Kind() == reflect.Int {
						ovb.FieldByName(field.Name).SetInt(sValue.Field(j).Int())
					}
				}
			}
		case reflect.Float64:
			sValue := va.FieldByName(field.Name)
			if sValue.Kind() == reflect.Struct {
				for j := 0; j < sValue.NumField(); j++ {
					if sValue.Field(j).Kind() == reflect.Float64 {
						ovb.FieldByName(field.Name).SetFloat(sValue.Field(j).Float())
					}
				}
			}
		}
	}
}

// DTO2Entity 将字段类型为普通类型的DTO 转换为字段类型为sql.NullXXX的实体，转换过程以Entity为主
func DTO2Entity(a interface{}, b interface{}, tags ...string) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	tb := reflect.TypeOf(b)
	ovb := vb.Elem()
	otb := tb.Elem()
	for i := 0; i < ovb.NumField(); i++ {
		field := otb.Field(i)
		if tag := field.Tag.Get("json"); len(tags) > 0 && containArray(tag, tags) {
			continue
		}
		if ovb.Field(i).Kind() == reflect.Struct {
			fieldA := va.FieldByName(field.Name)
			switch fieldA.Kind() {
			case reflect.String:
				if fieldA.String() == "" {
					ovb.FieldByName(field.Name).FieldByName("Valid").SetBool(false)
				} else {
					ovb.FieldByName(field.Name).FieldByName("Valid").SetBool(true)
					ovb.FieldByName(field.Name).FieldByName("String").SetString(fieldA.String())
				}
			case reflect.Int:
				if fieldA.Int() == 0 {
					ovb.FieldByName(field.Name).FieldByName("Valid").SetBool(false)
				} else {
					ovb.FieldByName(field.Name).FieldByName("Valid").SetBool(true)
					ovb.FieldByName(field.Name).FieldByName("Int32").SetInt(fieldA.Int())
				}
			case reflect.Float64:
				if fieldA.Float() == 0 {
					ovb.FieldByName(field.Name).FieldByName("Valid").SetBool(false)
				} else {
					ovb.FieldByName(field.Name).FieldByName("Valid").SetBool(true)
					ovb.FieldByName(field.Name).FieldByName("Float64").SetFloat(fieldA.Float())
				}
			}
		}
	}
}

func IntToNullInt32(a int) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(a), Valid: true}
}

func StringToNullString(a string) sql.NullString {
	return sql.NullString{String: a, Valid: true}
}

func anythingToSlice(a interface{}) []interface{} {
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		result := make([]interface{}, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = v.Index(i).Interface()
			t := reflect.TypeOf(result[i])
			fmt.Println("t = ", t)
		}
		return result
	default:
		panic("not supported")
	}
}

func GetLength(a interface{}) int {
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return v.Len()
	default:
		panic("not supported")
	}
}

func ArrayToString(array []string) string {
	var result string
	for i := 0; i < len(array); i++ {
		result += array[i] + ","
	}
	if strings.Contains(result, ",") {
		result = result[:strings.LastIndex(result, ",")]
	}
	return result
}

func containArray(tagName string, args []string) bool {
	for i := 0; i < len(args); i++ {
		if tagName == args[i] {
			return true
		}
	}
	return false
}

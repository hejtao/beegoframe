package dao

import (
	"errors"
	"reflect"
	"strings"
)

var (
	errDeleteWithEmptyCond             = errors.New("delete with empty condition")
	errQueryWithEmptyCond              = errors.New("query with empty condition")
	errQueryRowNotFound                = errors.New("query row not found")
	errUpdateWithEmptyCond             = errors.New("update with empty condition")
	errUpdateWithEmptyData             = errors.New("update with empty data")
	errInsertWithEmptyData             = errors.New("insert with empty data")
	errQueryRowWithWrongContainerType  = errors.New("wrong container type, need struct/map pointer")
	errQueryRowsWithWrongContainerType = errors.New("wrong container type, need slice pointer")
)

func checkPtr(container interface{}) error {
	if reflect.TypeOf(container).Kind() != reflect.Ptr {
		return errQueryRowWithWrongContainerType
	}
	return nil
}

func checkSlicePtr(container interface{}) error {
	if typ := reflect.TypeOf(container); typ.Kind() != reflect.Ptr {
		return errQueryRowsWithWrongContainerType
	} else if typ.Elem().Kind() != reflect.Slice {
		return errQueryRowsWithWrongContainerType
	}
	return nil
}

//func unfold(arg interface{})[]interface{}{
//	v := reflect.ValueOf(arg)
//	if v.Kind() == reflect.Slice{
//
//	}
//}

// v can be struct, *struct, []struct, []*struct, *[]*struct
func getSelectFields(v interface{}) []string {
	t := reflect.TypeOf(v)
	fields := getFields(t)
	if len(fields) == 0 {
		fields = append(fields, "*")
	}
	return fields
}

func getFields(t reflect.Type) []string {
	var fields []string
	// slice ptr?
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// slice ?
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	// struct ptr ?
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		fieldTag := t.Field(i).Tag.Get("field")
		if fieldTag == "" {
			fieldTag = jsonTag
		} else {
			fieldTag = fieldTag + " " + jsonTag
		}
		aliasTag := t.Field(i).Tag.Get("alias")
		if aliasTag == "" {
			aliasTag = "t0"
		}
		aliasTag = strings.ToUpper(aliasTag)
		fieldTag = aliasTag + "." + fieldTag
		if fieldTag != "" {
			fields = append(fields, fieldTag)
			continue
		}
		t2 := t.Field(i).Type
		fields = append(fields, getFields(t2)...)
	}
	return fields
}

// v can be struct, *struct, []struct, []*struct, *[]*struct
func getJsonTags(v interface{}) []string {
	t := reflect.TypeOf(v)
	return getTags(t)
}

func getTags(t reflect.Type) []string {
	var tags []string
	// slice ptr?
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// slice ?
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	// struct ptr ?
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		if jsonTag != "" {
			tags = append(tags, jsonTag)
			continue
		}
		t2 := t.Field(i).Type
		tags = append(tags, getFields(t2)...)
	}
	return tags
}

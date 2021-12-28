package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"

	"beegoframe/app/model"
)

func main() {
	system := "index" // change to admin for the back desk administration system
	projectName := "beegoframe"
	entity := model.Reader{}

	// app/dao/xx.go
	genDao(projectName, entity)

	// app/system/index/internal/service/xx.go
	genService(system, projectName, entity)

	// app/system/index/internal/controller/xx.go
	genController(system, projectName, entity)
}

const (
	ProjectName     = "{PROJECT_NAME}"
	EntityName      = "{ENTITY_NAME}"
	LowerEntityName = "{LOWER_ENTITY_NAME}"
	DaoFieldsDef    = "{DAO_FIELDS_DEF}"
	DaoFieldsInit   = "{DAO_FIELDS_INIT}"
)

func genDao(projectName string, entity interface{}) {
	entityName := getEntityName(entity)
	lowerEntityName := lowerFirst(entityName)
	initCode := getDaoInitCode(entity)
	defCode := getDaoFieldsDefCode(entity)
	content := readFile("./dao.txt")
	content = strings.ReplaceAll(content, ProjectName, projectName)
	content = strings.ReplaceAll(content, EntityName, entityName)
	content = strings.ReplaceAll(content, LowerEntityName, lowerEntityName)
	content = strings.ReplaceAll(content, DaoFieldsInit, initCode)
	content = strings.ReplaceAll(content, DaoFieldsDef, defCode)
	file := fmt.Sprintf("../../app/dao/%s.go", toSnake(entityName))
	_, err := os.Stat(file)
	if err == nil {
		fmt.Println(fmt.Sprintf("%s internal dao file exits", entityName))
		return
	}
	f, _ := os.Create(file)
	defer f.Close()
	f.WriteString(content)
}

func genService(system, projectName string, entity interface{}) {
	entityName := getEntityName(entity)
	lowerEntityName := lowerFirst(entityName)
	content := readFile("./service.txt")
	content = strings.ReplaceAll(content, ProjectName, projectName)
	content = strings.ReplaceAll(content, EntityName, entityName)
	content = strings.ReplaceAll(content, LowerEntityName, lowerEntityName)
	file := fmt.Sprintf("../../app/system/%s/internal/service/%s.go", system, toSnake(entityName))
	_, err := os.Stat(file)
	if err == nil {
		fmt.Println(fmt.Sprintf("index %s service file exits", entityName))
		return
	}
	f, _ := os.Create(file)
	defer f.Close()
	f.WriteString(content)
}

func genController(system, projectName string, entity interface{}) {
	entityName := getEntityName(entity)
	lowerEntityName := lowerFirst(entityName)
	content := readFile("./controller.txt")
	content = strings.ReplaceAll(content, ProjectName, projectName)
	content = strings.ReplaceAll(content, EntityName, entityName)
	content = strings.ReplaceAll(content, LowerEntityName, lowerEntityName)
	file := fmt.Sprintf("../../app/system/%s/internal/controller/%s.go", system, toSnake(entityName))
	_, err := os.Stat(file)
	if err == nil {
		fmt.Println(fmt.Sprintf("index %s controller file exits", entityName))
		return
	}
	f, _ := os.Create(file)
	defer f.Close()
	f.WriteString(content)
}

func getDaoFieldsDefCode(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}
	res := getDaoFieldsContentCode(t)
	return fmt.Sprintf(`
type %sFields struct {
%s
}`, lowerFirst(t.Name()), res[:len(res)-1])
}

func getDaoFieldsContentCode(t reflect.Type) string {
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
		return ""
	}
	var res string
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		if name != "Base" {
			res += fmt.Sprintf(`	%s internal.Field
`, name)
			continue
		}
		t2 := t.Field(i).Type
		res += getDaoFieldsContentCode(t2)
	}
	return res
}

func getDaoInitCode(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}
	res := getDaoFieldsInitCode(t)

	return fmt.Sprintf(`
var %s = &%sDao{
	setter:internal.DB.QueryTable((*model.%s)(nil)),
	table: "%s",
	Fields: &%sFields{
%s
	},
}`, t.Name(), lowerFirst(t.Name()), t.Name(), toSnake(t.Name()), lowerFirst(t.Name()), res[:len(res)-1])
}

func getDaoFieldsInitCode(t reflect.Type) string {
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
		return ""
	}
	var res string
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		if name != "Base" {
			res += fmt.Sprintf(`		%s: "%s",
`, name, toSnake(name))
			continue
		}
		t2 := t.Field(i).Type
		res += getDaoFieldsInitCode(t2)
	}
	return res
}

func getEntityName(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}
	return t.Name()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnake(camel string) string {
	snake := matchFirstCap.ReplaceAllString(camel, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func lowerFirst(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

func readFile(path string) string {
	file, _ := os.Open(path)
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	return string(content)
}

package openapi

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"log"
)

func getString(m map[string]interface{}, fieldName string) (value string, ok bool) {
	_, ok = m[fieldName]
	if !ok {
		return "", false
	}
	value, ok = m[fieldName].(string)
	return
}

func getStringOrDefault(m map[string]interface{}, fieldName string, def string) (value string) {
	value, ok := getString(m, fieldName)
	if !ok {
		return def
	}
	return value
}

func mustGetString(m map[string]interface{}, fieldName string) (value string) {
	value, ok := getString(m, fieldName)
	if !ok {
		log.Panicf("tried to get field named: \"%s\" but did not exist or was not a string", fieldName)
	}
	return value
}

func getBool( m map[string]interface{}, fieldName string, value *optional.Bool) (err error) {
	valueString := getStringOrDefault(m, fieldName, "false")
	switch valueString {
	case "false":
		value.Set(false)
		return nil
	case "true":
		value.Set(true)
		return nil
	default:
		value.Unset()
		return fmt.Errorf("boolean field: \"%s\" must be true or false, but got \"%s\"", fieldName, valueString)
	}
}

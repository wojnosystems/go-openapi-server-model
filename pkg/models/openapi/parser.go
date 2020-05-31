package openapi

import (
	"errors"
	"io"
)
import "gopkg.in/yaml.v2"


func ParseModels( reader io.Reader ) (models SchemaModels, err error) {
	models = NewSchemaModels()
	var ok bool
	yaml := yaml.NewDecoder(reader)

	var document map[string]interface{}
	err = yaml.Decode(&document)
	if err != nil {
		return
	}
	components, err := extractStringMap(document, "components", errors.New("no components"))
	if err != nil {
		return
	}
	schemas, err := extractStringMap(components, "schemas", errors.New("no schemas"))
	if err != nil {
		return
	}

	for modelName := range schemas {
		body, err := extractStringMap(schemas, modelName, nil)
		if err != nil {
			return
		}
		models[modelName], err = parseModel(modelName,body)
	}

	return
}

func extractStringMap( field map[string]interface{}, key string, errIfMissing error ) (castType map[string]interface{}, err error) {
	var ok bool
	var extractedField interface{}
	if extractedField, ok = field[key]; !ok {
		err = errIfMissing
		return
	}
	castType, ok = extractedField.(map[string]interface{})
	if !ok {
		return nil, errors.New("unable to cast to map[string]interface{}")
	}
	return castType, nil
}

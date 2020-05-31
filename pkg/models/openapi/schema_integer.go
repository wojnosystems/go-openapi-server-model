package openapi

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"openapi-go-server-model/pkg/string_utils"
	"strconv"
)

type SchemaInteger struct {
	modelBase
	Format           optional.String
	Minimum          optional.Int64
	Maximum          optional.Int64
	ExclusiveMinimum optional.Bool
	ExclusiveMaximum optional.Bool
	MultipleOf       optional.Int64
}

func parseSchemaInteger(modelName string, modelBody map[string]interface{} ) (model Modeler, err error) {
	intModel := &SchemaInteger{
		modelBase: modelBase{
			name: modelName,
			description: getStringOrDefault(modelBody, "description", ""),
		},
		MultipleOf: optional.NewInt64(),
	}
	intModel.Format, err = intModel.parseFormat(modelBody)
	if err != nil {
		return
	}
	err = getBool(modelBody, "exclusiveMaximum", &intModel.ExclusiveMaximum)
	if err != nil {
		return
	}
	err = getBool(modelBody, "exclusiveMinimum", &intModel.ExclusiveMinimum)
	if err != nil {
		return
	}
	err = parseInt64Field(&intModel.Minimum, modelBody, "minimum")
	if err != nil {
		return
	}
	err = parseInt64Field(&intModel.Maximum, modelBody, "maximum")
	if err != nil {
		return
	}
	err = parseInt64Field(&intModel.MultipleOf, modelBody, "multipleOf")
	if err != nil {
		return
	}

	if violations := intModel.Validate(); len(violations) != 0 {
		model = intModel
	}

	return
}

func (s SchemaInteger) parseFormat(modelBody map[string]interface{}) (format optional.String, err error) {
	extractedValue := getStringOrDefault(modelBody, "format", "")
	if !string_utils.IsBlank(&extractedValue) {
		format.Set(extractedValue)
	}
	return
}

func parseInt64Field(field *optional.Int64, modelBody map[string]interface{}, fieldName string ) (err error) {
	extractedValue := getStringOrDefault(modelBody, fieldName, "")

	if !string_utils.IsBlank(&extractedValue) {
		var value int64
		value, err = strconv.ParseInt(extractedValue, 10, 64)
		if err != nil {
			return
		}
		field.Set(value)
	}
	return
}

func (s SchemaInteger) Validate() (violations constraint.Violations) {

}

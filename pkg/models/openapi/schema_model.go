package openapi

import (
	"fmt"
)

type Modeler interface{
	Name() string
	Description() string
}

func parseModel( schemas map[string]interface{}, modelName string, modelBody map[string]interface{}) (model Modeler, err error) {
	if hasRef(modelBody) {
		return resolveRef(schemas, modelBody["$ref"])
	}

	switch parseSchemaType(modelBody["type"]) {
	case SchemaTypeObject:
		return parseObject(modelName, modelBody)
	case SchemaTypeInteger:
		return parseSchemaInteger(modelName, modelBody)
	default:
		return nil, fmt.Errorf("not a valid schema type: \"%s\"", modelBody["type"])
	}
}

func hasRef(model map[string]interface{}) bool {
	_, ok := model["$ref"]
	return ok
}

func parseSchemaType( objectType interface{} ) (schemaType SchemaTypeEnum) {
	stringVal, ok := objectType.(string)
	if !ok {
		return SchemaTypeInvalid
	}
	switch stringVal {
	case SchemaTypeObject:
		return SchemaTypeObject
	case SchemaTypeString:
		return SchemaTypeString
	case SchemaTypeInteger:
		return SchemaTypeInteger
	case SchemaTypeNumber:
		return SchemaTypeNumber
	default:
		return SchemaTypeInvalid
	}
}

func resolveRef( schemas map[string]interface{}, refValue interface{} ) (Modeler, error) {
	return nil, nil
}
package openapi

type SchemaObject struct {
	modelBase
	properties map[string]Modeler
}

func parseObject(name string, body map[string]interface{}) (Modeler, error) {
	return nil, nil
}
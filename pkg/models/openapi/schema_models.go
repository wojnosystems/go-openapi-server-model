package openapi

type SchemaModels map[string]Modeler

func NewSchemaModels() SchemaModels {
	return make(map[string]Modeler)
}

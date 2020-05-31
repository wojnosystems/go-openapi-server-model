package openapi

type modelBase struct {
	name string
	description string
}

func (m modelBase) Name() string {
	return m.name
}

func (m modelBase) Description() string {
	return m.description
}
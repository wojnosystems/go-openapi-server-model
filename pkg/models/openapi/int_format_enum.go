package openapi

type IntFormatType uint8

const (
	IntFormatTypeInvalid IntFormatType = iota
	IntFormatTypeUnset
	IntFormatTypeUnknown
	IntFormatTypeInt32
	IntFormatTypeInt64
)

func ParseIntFormatType( value string ) IntFormatType {
	switch value {
	case "":
		return IntFormatTypeUnset
	case "int32":
		return IntFormatTypeInt32
	case "int64":
		return IntFormatTypeInt64
	default:
		return IntFormatTypeUnknown
	}
}

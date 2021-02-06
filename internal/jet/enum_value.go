package jet

type enumValue struct {
	expressionInterfaceImpl
	stringInterfaceImpl

	name string
}

// NewEnumValue creates new named enum value
func NewEnumValue(name string) StringExpression {
	enumValue := &enumValue{name: name}

	enumValue.expressionInterfaceImpl.parent = enumValue
	enumValue.stringInterfaceImpl.parent = enumValue

	return enumValue
}

func (e enumValue) Serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	out.insertConstantArgument(e.name)
}

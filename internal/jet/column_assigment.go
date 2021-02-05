package jet

// ColumnAssigment is interface wrapper around column assigment
type ColumnAssigment interface {
	Serializer
	isColumnAssigment()
}

type columnAssigmentImpl struct {
	column     ColumnSerializer
	expression Expression
}

func NewColumnAssigment(column ColumnSerializer, expression Expression) *columnAssigmentImpl {
	return &columnAssigmentImpl{
		column:     column,
		expression: expression,
	}
}

func (a columnAssigmentImpl) isColumnAssigment() {}

func (a columnAssigmentImpl) Serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	a.column.Serialize(statement, out, ShortName.WithFallTrough(options)...)
	out.WriteString("=")
	a.expression.Serialize(statement, out, FallTrough(options)...)
}

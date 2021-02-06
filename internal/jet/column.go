// Modeling of columns

package jet

// Column is common column interface for all types of columns.
type Column interface {
	Name() string
	TableName() string

	SetTableName(table string)
	SetSubQuery(subQuery SelectTable)
	defaultAlias() string
}

// ColumnSerializer is interface for all serializable columns
type ColumnSerializer interface {
	Serializer
	Column
}

// ColumnExpression interface
type ColumnExpression interface {
	Column
	Expression
}

// ColumnExpressionImpl is base type for sql columns.
type columnExpressionImpl struct {
	expressionInterfaceImpl

	name      string
	tableName string

	subQuery SelectTable
}

// NewColumnExpression creates new ColumnExpressionImpl
func NewColumnExpression(name string, tableName string, parent ColumnExpression) *columnExpressionImpl {
	bc := &columnExpressionImpl{
		name:      name,
		tableName: tableName,
	}

	if parent != nil {
		bc.expressionInterfaceImpl.parent = parent
	} else {
		bc.expressionInterfaceImpl.parent = bc
	}

	return bc
}

// Name returns name of the column
func (c *columnExpressionImpl) Name() string {
	return c.name
}

// TableName returns column table name
func (c *columnExpressionImpl) TableName() string {
	return c.tableName
}

func (c *columnExpressionImpl) SetTableName(table string) {
	c.tableName = table
}

func (c *columnExpressionImpl) SetSubQuery(subQuery SelectTable) {
	c.subQuery = subQuery
}

func (c *columnExpressionImpl) defaultAlias() string {
	if c.tableName != "" {
		return c.tableName + "." + c.name
	}

	return c.name
}

func (c *columnExpressionImpl) fromImpl(subQuery SelectTable) Projection {
	newColumn := NewColumnExpression(c.name, c.tableName, nil)
	newColumn.SetSubQuery(subQuery)

	return newColumn
}

func (c *columnExpressionImpl) serializeForOrderBy(statement StatementType, out *SQLBuilder) {
	if statement == SetStatementType {
		// set Statement (UNION, EXCEPT ...) can reference only select projections in order by clause
		out.WriteAlias(c.defaultAlias()) //always quote
		return
	}

	c.Serialize(statement, out)
}

func (c columnExpressionImpl) serializeForProjection(statement StatementType, out *SQLBuilder) {
	c.Serialize(statement, out)

	out.WriteString("AS")
	out.WriteAlias(c.defaultAlias())
}

func (c columnExpressionImpl) Serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	if c.subQuery != nil {
		out.WriteIdentifier(c.subQuery.Alias())
		out.WriteByte('.')
		out.WriteIdentifier(c.defaultAlias())
	} else {
		if c.tableName != "" && !serializeOptionsContain(options, ShortName) {
			out.WriteIdentifier(c.tableName)
			out.WriteByte('.')
		}

		out.WriteIdentifier(c.name)
	}
}

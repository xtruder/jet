package jet

// SelectTable is interface for SELECT sub-queries
type SelectTable interface {
	Serializer
	Alias() string
	AllColumns() ProjectionList
}

type selectTableImpl struct {
	selectStmt SerializerStatement
	alias      string
}

// NewSelectTable func
func NewSelectTable(selectStmt SerializerStatement, alias string) SelectTable {
	selectTable := &selectTableImpl{selectStmt: selectStmt, alias: alias}
	return selectTable
}

func (s selectTableImpl) Alias() string {
	return s.alias
}

func (s selectTableImpl) AllColumns() ProjectionList {
	statementWithProjections, ok := s.selectStmt.(HasProjections)
	if !ok {
		return ProjectionList{}
	}

	projectionList := statementWithProjections.projections().fromImpl(s)
	return projectionList.(ProjectionList)
}

func (s selectTableImpl) Serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	s.selectStmt.Serialize(statement, out)

	out.WriteString("AS")
	out.WriteIdentifier(s.alias)
}

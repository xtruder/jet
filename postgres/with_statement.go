package postgres

import "github.com/go-jet/jet/internal/jet"

// CommonTableExpression contains information about a CTE.
type CommonTableExpression struct {
	readableTableInterfaceImpl
	jet.CommonTableExpression
}

// CommonTableExpressionDefinition defines CTE.
type CommonTableExpressionDefinition = jet.CommonTableExpressionDefinition

// WITH function creates new WITH statement from list of common table expressions
func WITH(cte ...CommonTableExpressionDefinition) func(statement jet.Statement) Statement {
	return jet.WITH(Dialect, cte...)
}

// CTE creates new named CommonTableExpression
func CTE(name string) CommonTableExpression {
	cte := CommonTableExpression{
		readableTableInterfaceImpl: readableTableInterfaceImpl{},
		CommonTableExpression:      jet.CTE(name),
	}

	cte.parent = &cte

	return cte
}

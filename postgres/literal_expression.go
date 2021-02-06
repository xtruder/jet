package postgres

import (
	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/lib/pq"
)

type arrayLiteral struct {
	arrayInterfaceImpl
	jet.LiteralExpression
}

// newArrayLiteral creates new array literal expression
func newArrayLiteral(values interface{}) ArrayExpression {
	arrayLiteral := arrayLiteral{}
	arrayLiteral.LiteralExpression = jet.Literal(pq.Array(values))
	arrayLiteral.arrayInterfaceImpl.parent = &arrayLiteral

	return &arrayLiteral
}

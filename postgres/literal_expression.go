package postgres

import (
	"encoding/json"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/utils"
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

type jsonLiteral struct {
	jsonInterfaceImpl
	jet.LiteralExpression
}

func newJsonLiteral(value interface{}) JSONExpression {
	data, err := json.Marshal(value)
	utils.PanicOnError(err)

	jsonLiteral := jsonLiteral{}
	jsonLiteral.LiteralExpression = jet.Literal(string(data))
	jsonLiteral.jsonInterfaceImpl.parent = &jsonLiteral

	return &jsonLiteral
}

type jsonbLiteral struct {
	jsonbInterfaceImpl
	jet.LiteralExpression
}

func newJsonbLiteral(value interface{}) JSONBExpression {
	data, err := json.Marshal(value)
	utils.PanicOnError(err)

	jsonbLiteral := jsonbLiteral{}
	jsonbLiteral.LiteralExpression = jet.Literal(string(data))
	jsonbLiteral.jsonbInterfaceImpl.parent = &jsonbLiteral

	return &jsonbLiteral
}

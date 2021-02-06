package mysql

import (
	"encoding/json"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/go-jet/jet/v2/internal/utils"
)

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

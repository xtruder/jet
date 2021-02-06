package mysql

import "github.com/go-jet/jet/v2/internal/jet"

type jsonFunc struct {
	jet.Func
	jsonInterfaceImpl
}

func newJSONFunc(name string, expressions ...Expression) JSONExpression {
	jsonFunc := &jsonFunc{}

	jsonFunc.Func = *jet.NewFunc(name, expressions, jsonFunc)
	jsonFunc.jsonInterfaceImpl.parent = jsonFunc

	return jsonFunc
}

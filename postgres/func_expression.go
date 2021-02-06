package postgres

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

type jsonbFunc struct {
	jet.Func
	jsonbInterfaceImpl
}

func newJSONBFunc(name string, expressions ...Expression) JSONBExpression {
	jsonFunc := &jsonbFunc{}

	jsonFunc.Func = *jet.NewFunc(name, expressions, jsonFunc)
	jsonFunc.jsonbInterfaceImpl.parent = jsonFunc

	return jsonFunc
}

package postgres

import (
	"database/sql/driver"
	"strconv"
	"strings"
	"time"

	"github.com/go-jet/jet/v2/internal/jet"
	"github.com/lib/pq"
)

// Dialect is implementation of postgres dialect for SQL Builder serialisation.
var Dialect = newDialect()

func newDialect() jet.Dialect {
	operatorSerializeOverrides := map[string]jet.SerializeOverride{}
	operatorSerializeOverrides[jet.StringRegexpLikeOperator] = postgresREGEXPLIKEoperator
	operatorSerializeOverrides[jet.StringNotRegexpLikeOperator] = postgresNOTREGEXPLIKEoperator
	operatorSerializeOverrides["CAST"] = postgresCAST

	dialectParams := jet.DialectParams{
		Name:                       "PostgreSQL",
		PackageName:                "postgres",
		OperatorSerializeOverrides: operatorSerializeOverrides,
		AliasQuoteChar:             '"',
		IdentifierQuoteChar:        '"',
		ArgumentPlaceholder: func(ord int) string {
			return "$" + strconv.Itoa(ord)
		},
		ReservedWords:   reservedWords,
		ArgToStringFunc: argToString,
	}

	return jet.NewDialect(dialectParams)
}

func argToString(value interface{}) string {
	switch bindVal := value.(type) {
	case time.Time:
		return stringQuote(string(pq.FormatTimestamp(bindVal)))
	case *pq.StringArray,
		*pq.BoolArray,
		*pq.ByteaArray,
		*pq.Float32Array,
		*pq.Float64Array,
		*pq.Int32Array,
		*pq.Int64Array:
		// reuse driver.Valuer to convert to array value
		val, err := bindVal.(driver.Valuer).Value()
		if err != nil {
			panic(err)
		}

		return stringQuote(val.(string))
	}

	return ""
}

func stringQuote(value string) string {
	return `'` + strings.Replace(value, "'", "''", -1) + `'`
}

func postgresCAST(expressions ...jet.Serializer) jet.SerializerFunc {
	return func(statement jet.StatementType, out *jet.SQLBuilder, options ...jet.SerializeOption) {
		if len(expressions) < 2 {
			panic("jet: invalid number of expressions for operator")
		}

		expression := expressions[0]

		litExpr, ok := expressions[1].(jet.LiteralExpression)

		if !ok {
			panic("jet: cast invalid cast type")
		}

		castType, ok := litExpr.Value().(string)

		if !ok {
			panic("jet: cast type is not string")
		}

		expression.Serialize(statement, out, options...)
		out.WriteString("::" + castType)
	}
}

func postgresREGEXPLIKEoperator(expressions ...jet.Serializer) jet.SerializerFunc {
	return func(statement jet.StatementType, out *jet.SQLBuilder, options ...jet.SerializeOption) {
		if len(expressions) < 2 {
			panic("jet: invalid number of expressions for operator")
		}

		expressions[0].Serialize(statement, out, options...)

		caseSensitive := false

		if len(expressions) >= 3 {
			if stringLiteral, ok := expressions[2].(jet.LiteralExpression); ok {
				caseSensitive = stringLiteral.Value().(bool)
			}
		}

		if caseSensitive {
			out.WriteString("~")
		} else {
			out.WriteString("~*")
		}

		expressions[1].Serialize(statement, out, options...)
	}
}

func postgresNOTREGEXPLIKEoperator(expressions ...jet.Serializer) jet.SerializerFunc {
	return func(statement jet.StatementType, out *jet.SQLBuilder, options ...jet.SerializeOption) {
		if len(expressions) < 2 {
			panic("jet: invalid number of expressions for operator")
		}

		expressions[0].Serialize(statement, out, options...)

		caseSensitive := false

		if len(expressions) >= 3 {
			if stringLiteral, ok := expressions[2].(jet.LiteralExpression); ok {
				caseSensitive = stringLiteral.Value().(bool)
			}
		}

		if caseSensitive {
			out.WriteString("!~")
		} else {
			out.WriteString("!~*")
		}

		expressions[1].Serialize(statement, out, options...)
	}
}

var reservedWords = []string{
	"ALL",
	"ANALYSE",
	"ANALYZE",
	"AND",
	"ANY",
	"ARRAY",
	"AS",
	"ASC",
	"ASYMMETRIC",
	"BOTH",
	"CASE",
	"CAST",
	"CHECK",
	"COLLATE",
	"COLUMN",
	"CONSTRAINT",
	"CREATE",
	"CURRENT_CATALOG",
	"CURRENT_DATE",
	"CURRENT_ROLE",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER",
	"DEFAULT",
	"DEFERRABLE",
	"DESC",
	"DISTINCT",
	"DO",
	"ELSE",
	"END",
	"EXCEPT",
	"FALSE",
	"FETCH",
	"FOR",
	"FOREIGN",
	"FROM",
	"GRANT",
	"GROUP",
	"HAVING",
	"IN",
	"INITIALLY",
	"INTERSECT",
	"INTO",
	"LATERAL",
	"LEADING",
	"LIMIT",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"NOT",
	"NULL",
	"OFFSET",
	"ON",
	"ONLY",
	"OR",
	"ORDER",
	"PLACING",
	"PRIMARY",
	"REFERENCES",
	"RETURNING",
	"SELECT",
	"SESSION_USER",
	"SOME",
	"SYMMETRIC",
	"TABLE",
	"THEN",
	"TO",
	"TRAILING",
	"TRUE",
	"UNION",
	"UNIQUE",
	"USER",
	"USING",
	"VARIADIC",
	"WHEN",
	"WHERE",
	"WINDOW",
	"WITH",
}

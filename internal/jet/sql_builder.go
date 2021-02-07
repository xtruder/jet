package jet

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/go-jet/jet/v2/internal/utils"
	"github.com/go-jet/jet/v2/qrm"
)

// SerializeOption type
type SQLBuilderOption int

// Serialize options
const (
	SQLBuilderOptPretty SQLBuilderOption = iota
	SQLBuilderOptDebug
)

func sqlBuilderOptionsContain(options []SQLBuilderOption, option SQLBuilderOption) bool {
	for _, opt := range options {
		if opt == option {
			return true
		}
	}

	return false
}

// SQLBuilder generates output SQL
type SQLBuilder struct {
	Dialect Dialect
	Buff    bytes.Buffer
	Args    []interface{}

	lastChar byte
	ident    int

	Debug  bool
	Pretty bool
}

func NewSQLBuilder(dialect Dialect, options ...SQLBuilderOption) *SQLBuilder {
	builder := &SQLBuilder{Dialect: dialect}

	if sqlBuilderOptionsContain(options, SQLBuilderOptPretty) {
		builder.Pretty = true
	}

	if sqlBuilderOptionsContain(options, SQLBuilderOptDebug) {
		builder.Debug = true
	}

	return builder
}

const defaultIdent = 5

// IncreaseIdent adds ident or defaultIdent number of spaces to each new line
func (s *SQLBuilder) IncreaseIdent(ident ...int) {
	if len(ident) > 0 {
		s.ident += ident[0]
	} else {
		s.ident += defaultIdent
	}
}

// DecreaseIdent removes ident or defaultIdent number of spaces for each new line
func (s *SQLBuilder) DecreaseIdent(ident ...int) {
	toDecrease := defaultIdent

	if len(ident) > 0 {
		toDecrease = ident[0]
	}

	if s.ident < toDecrease {
		s.ident = 0
	}

	s.ident -= toDecrease
}

// WriteProjections func
func (s *SQLBuilder) WriteProjections(statement StatementType, projections []Projection) {
	s.IncreaseIdent()
	SerializeProjectionList(statement, projections, s)
	s.DecreaseIdent()
}

// NewLine adds new line to output SQL
func (s *SQLBuilder) NewLine() {
	s.write([]byte{'\n'})
	s.write(bytes.Repeat([]byte{' '}, s.ident))
}

func (s *SQLBuilder) Space() {
	s.write([]byte{' '})
}

func (s *SQLBuilder) write(data []byte) {
	if len(data) == 0 {
		return
	}

	if !isPreSeparator(s.lastChar) && !isPostSeparator(data[0]) && s.Buff.Len() > 0 {
		s.Buff.WriteByte(' ')
	}

	s.Buff.Write(data)
	s.lastChar = data[len(data)-1]
}

func isPreSeparator(b byte) bool {
	return b == ' ' || b == '.' || b == ',' || b == '(' || b == '\n' || b == ':'
}

func isPostSeparator(b byte) bool {
	return b == ' ' || b == '.' || b == ',' || b == ')' || b == '\n' || b == ':'
}

// WriteAlias is used to add alias to output SQL
func (s *SQLBuilder) WriteAlias(str string) {
	aliasQuoteChar := string(s.Dialect.AliasQuoteChar())
	s.WriteString(aliasQuoteChar + str + aliasQuoteChar)
}

// WriteString writes sting to output SQL
func (s *SQLBuilder) WriteString(str string) {
	s.write([]byte(str))
}

// WriteIdentifier adds identifier to output SQL
func (s *SQLBuilder) WriteIdentifier(name string, alwaysQuote ...bool) {
	if s.shouldQuote(name, alwaysQuote...) {
		identQuoteChar := string(s.Dialect.IdentifierQuoteChar())
		s.WriteString(identQuoteChar + name + identQuoteChar)
	} else {
		s.WriteString(name)
	}
}

func (s *SQLBuilder) shouldQuote(name string, alwaysQuote ...bool) bool {
	return s.Dialect.IsReservedWord(name) || shouldQuoteIdentifier(name) || len(alwaysQuote) > 0
}

// WriteByte writes byte to output SQL
func (s *SQLBuilder) WriteByte(b byte) {
	s.write([]byte{b})
}

func (s *SQLBuilder) finalize() (string, []interface{}) {
	if s.Pretty {
		return s.Buff.String() + ";\n", s.Args
	} else {
		return s.Buff.String() + ";", s.Args
	}
}

func (s *SQLBuilder) insertConstantArgument(arg interface{}) {
	if val := s.Dialect.ArgToString(arg); val != "" {
		s.WriteString(val)
		return
	}

	s.WriteString(argToString(arg))
}

func (s *SQLBuilder) insertParametrizedArgument(arg interface{}) {
	if s.Debug {
		s.insertConstantArgument(arg)
		return
	}

	s.Args = append(s.Args, arg)
	argPlaceholder := s.Dialect.ArgumentPlaceholder()(len(s.Args))

	s.WriteString(argPlaceholder)
}

// default argToStr implementation
func argToString(value interface{}) string {
	if utils.IsNil(value) {
		return "NULL"
	}

	switch bindVal := value.(type) {
	case bool:
		if bindVal {
			return "TRUE"
		}
		return "FALSE"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return integerTypesToString(bindVal)

	case float32:
		return strconv.FormatFloat(float64(bindVal), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(bindVal), 'f', -1, 64)

	case string:
		return stringQuote(bindVal)
	case []byte:
		return stringQuote(string(bindVal))
	case time.Time:
		return stringQuote(bindVal.Format("2006-01-02 15:04:05.999999999Z07:00"))
	case qrm.JSON:
		return stringQuote(string(bindVal))
	case *qrm.JSON:
		return stringQuote(string(*bindVal))
	default:
		if strBindValue, ok := bindVal.(fmt.Stringer); ok {
			return stringQuote(strBindValue.String())
		}

		panic(fmt.Sprintf("jet: %T type can not be used as SQL query parameter", value))
	}
}

func integerTypesToString(value interface{}) string {
	switch bindVal := value.(type) {
	case int:
		return strconv.FormatInt(int64(bindVal), 10)
	case uint:
		return strconv.FormatUint(uint64(bindVal), 10)
	case int8:
		return strconv.FormatInt(int64(bindVal), 10)
	case uint8:
		return strconv.FormatUint(uint64(bindVal), 10)
	case int16:
		return strconv.FormatInt(int64(bindVal), 10)
	case uint16:
		return strconv.FormatUint(uint64(bindVal), 10)
	case int32:
		return strconv.FormatInt(int64(bindVal), 10)
	case uint32:
		return strconv.FormatUint(uint64(bindVal), 10)
	case int64:
		return strconv.FormatInt(bindVal, 10)
	case uint64:
		return strconv.FormatUint(bindVal, 10)
	}
	panic("jet: Unsupported integer type: " + reflect.TypeOf(value).String())
}

func shouldQuoteIdentifier(identifier string) bool {
	_, err := strconv.ParseInt(identifier, 10, 64)

	if err == nil { // if it is a number we should quote it
		return true
	}

	// check if contains non ascii characters
	for _, c := range identifier {
		if unicode.IsNumber(c) || c == '_' {
			continue
		}
		if c > unicode.MaxASCII || !unicode.IsLetter(c) || unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func stringQuote(value string) string {
	return `'` + strings.Replace(value, "'", "''", -1) + `'`
}

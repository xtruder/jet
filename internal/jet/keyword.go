package jet

const (
	// DEFAULT is jet equivalent of SQL DEFAULT
	DEFAULT Keyword = "DEFAULT"
)

// Keyword type
type Keyword string

func (k Keyword) Serialize(statement StatementType, out *SQLBuilder, options ...SerializeOption) {
	out.WriteString(string(k))
}

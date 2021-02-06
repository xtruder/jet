package postgres

import "testing"

func TestOperatorARRAY(t *testing.T) {
	assertSerialize(t, ARRAY(table2ColStr, String("test")), "ARRAY[ table2.col_str,$1 ]", "test")
}

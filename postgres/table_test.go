package postgres

import (
	"testing"
)

func TestJoinNilInputs(t *testing.T) {
	assertSerializeErr(t, table2.INNER_JOIN(nil, table1ColBool.EQ(table2ColBool)),
		"jet: right hand side of join operation is nil table")
	assertSerializeErr(t, table2.INNER_JOIN(table1, nil),
		"jet: join condition is nil")
}

func TestINNER_JOIN(t *testing.T) {
	assertRecordSerialize(t, table1.
		INNER_JOIN(table2, table1ColInt.EQ(table2ColInt)))
	assertRecordSerialize(t, table1.
		INNER_JOIN(table2, table1ColInt.EQ(table2ColInt)).
		INNER_JOIN(table3, table1ColInt.EQ(table3ColInt)))
	assertRecordSerialize(t, table1.
		INNER_JOIN(table2, table1ColInt.EQ(Int(1))).
		INNER_JOIN(table3, table1ColInt.EQ(Int(2))))
}

func TestLEFT_JOIN(t *testing.T) {
	assertRecordSerialize(t, table1.
		LEFT_JOIN(table2, table1ColInt.EQ(table2ColInt)))
	assertRecordSerialize(t, table1.
		LEFT_JOIN(table2, table1ColInt.EQ(table2ColInt)).
		LEFT_JOIN(table3, table1ColInt.EQ(table3ColInt)))
	assertRecordSerialize(t, table1.
		LEFT_JOIN(table2, table1ColInt.EQ(Int(1))).
		LEFT_JOIN(table3, table1ColInt.EQ(Int(2))))
}

func TestRIGHT_JOIN(t *testing.T) {
	assertRecordSerialize(t, table1.
		RIGHT_JOIN(table2, table1ColInt.EQ(table2ColInt)))
	assertRecordSerialize(t, table1.
		RIGHT_JOIN(table2, table1ColInt.EQ(table2ColInt)).
		RIGHT_JOIN(table3, table1ColInt.EQ(table3ColInt)))
	assertRecordSerialize(t, table1.
		RIGHT_JOIN(table2, table1ColInt.EQ(Int(1))).
		RIGHT_JOIN(table3, table1ColInt.EQ(Int(2))))
}

func TestFULL_JOIN(t *testing.T) {
	assertRecordSerialize(t, table1.
		FULL_JOIN(table2, table1ColInt.EQ(table2ColInt)))
	assertRecordSerialize(t, table1.
		FULL_JOIN(table2, table1ColInt.EQ(table2ColInt)).
		FULL_JOIN(table3, table1ColInt.EQ(table3ColInt)))
	assertRecordSerialize(t, table1.
		FULL_JOIN(table2, table1ColInt.EQ(Int(1))).
		FULL_JOIN(table3, table1ColInt.EQ(Int(2))))
}

func TestCROSS_JOIN(t *testing.T) {
	assertRecordSerialize(t, table1.
		CROSS_JOIN(table2))
	assertRecordSerialize(t, table1.
		CROSS_JOIN(table2).
		CROSS_JOIN(table3))
}

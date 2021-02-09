package postgres

import (
	"testing"

	"github.com/go-jet/jet/v2/internal/testutils"
)

func TestTO_JSON(t *testing.T) {
	testutils.SerializerTest{Test: TO_JSON(table2ColStr)}.Assert(t, Dialect)
}

func TestTO_JSONB(t *testing.T) {
	testutils.SerializerTest{Test: TO_JSONB(table2ColStr)}.Assert(t, Dialect)
}

func TestARRAY_TO_JSON(t *testing.T) {
	testutils.SerializerTest{Test: ARRAY_TO_JSON(table2ColArray)}.Assert(t, Dialect)
}

func TestROW_TO_JSON(t *testing.T) {
	testutils.SerializerTest{Test: ROW_TO_JSON("table2")}.Assert(t, Dialect)
}

func TestJSON_BUILD_ARRAY(t *testing.T) {
	testutils.SerializerTest{
		Test: JSON_BUILD_ARRAY(String("val1"), table2ColInt)}.Assert(t, Dialect)
}

func TestJSONB_BUILD_ARRAY(t *testing.T) {
	testutils.SerializerTest{
		Test: JSONB_BUILD_ARRAY(String("val1"), table2ColInt)}.Assert(t, Dialect)
}

func TestJSON_BUILD_OBJECT(t *testing.T) {
	testutils.SerializerTest{
		Test: JSON_BUILD_OBJECT(String("key"), table2ColJSON)}.Assert(t, Dialect)

}

func TestJSONB_BUILD_OBJECT(t *testing.T) {
	testutils.SerializerTest{
		Test: JSONB_BUILD_OBJECT(String("key"), table2ColJSON)}.Assert(t, Dialect)
}

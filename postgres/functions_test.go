package postgres

import "testing"

func TestTO_JSON(t *testing.T) {
	assertSerialize(t, TO_JSON(table2ColStr), "to_json(table2.col_str)")
}

func TestTO_JSONB(t *testing.T) {
	assertSerialize(t, TO_JSONB(table2ColStr), "to_jsonb(table2.col_str)")
}

func TestARRAY_TO_JSON(t *testing.T) {
	assertSerialize(t, ARRAY_TO_JSON(table2ColArray), "array_to_json(table2.col_array)")
}

func TestROW_TO_JSON(t *testing.T) {
	assertSerialize(t, ROW_TO_JSON("table2"), `row_to_json(table2)`)
}

func TestJSON_BUILD_ARRAY(t *testing.T) {
	assertSerialize(t,
		JSON_BUILD_ARRAY(String("val1"), table2ColInt),
		`json_build_array($1::text, table2.col_int)`, "val1")
}

func TestJSONB_BUILD_ARRAY(t *testing.T) {
	assertSerialize(t,
		JSONB_BUILD_ARRAY(String("val1"), table2ColInt),
		`jsonb_build_array($1::text, table2.col_int)`, "val1")
}

func TestJSON_BUILD_OBJECT(t *testing.T) {
	assertSerialize(t,
		JSON_BUILD_OBJECT(String("key"), table2ColJSON),
		`json_build_object($1::text, table2.col_json)`, "key")
}

func TestJSONB_BUILD_OBJECT(t *testing.T) {
	assertSerialize(t,
		JSONB_BUILD_OBJECT(String("key"), table2ColJSON),
		`jsonb_build_object($1::text, table2.col_json)`, "key")
}

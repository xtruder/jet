package mysql

import (
	"encoding/json"
	"testing"
)

func TestJSONExpression_EXTRACT(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).EXTRACT(String("$.id")),
		"(table2.col_json -> ?)", "$.id")
	assertSerialize(t, JSONExp(table2ColJSON).EXTRACT(String("$[0]"), String("$[1]")),
		"JSON_EXTRACT(?, ?)", "$[0]", "$[1]")
}

func TestJSONExpression_EXTRACT_UNQUOTE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).EXTRACT_UNQUOTE(String("$.id")),
		"(table2.col_json ->> ?)", "$.id")
}

func TestJSONExpression_UNQUOTE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).UNQUOTE(), "JSON_UNQUOTE(table2.col_json)")
}

func TestJSONExpression_ARRAY_APPEND(t *testing.T) {
	val := []string{"val1", "val2"}
	json, _ := json.Marshal(val)

	assertSerialize(t, JSONExp(table2ColJSON).ARRAY_APPEND(String("$.id"), JSON(val)),
		"JSON_ARRAY_APPEND(table2.col_json, ?, ?)", "$.id", string(json))
	assertSerialize(t,
		JSONExp(table2ColJSON).ARRAY_APPEND(String("$.id"), JSON(val), String("$.key"), JSON(val)),
		"JSON_ARRAY_APPEND(table2.col_json, ?, ?, ?, ?)", "$.id", string(json), "$.key", string(json))
}

func TestJSONExpression_ARRAY_INSERT(t *testing.T) {
	val := []string{"val1", "val2"}
	json, _ := json.Marshal(val)

	assertSerialize(t, JSONExp(table2ColJSON).ARRAY_INSERT(String("$.id"), JSON(val)),
		"JSON_ARRAY_INSERT(table2.col_json, ?, ?)", "$.id", string(json))
	assertSerialize(t,
		JSONExp(table2ColJSON).ARRAY_INSERT(String("$.id"), JSON(val), String("$.key"), JSON(val)),
		"JSON_ARRAY_INSERT(table2.col_json, ?, ?, ?, ?)", "$.id", string(json), "$.key", string(json))
}

func TestJSONExpression_CONTAINS(t *testing.T) {
	val := []string{"val1", "val2"}
	json, _ := json.Marshal(val)

	assertSerialize(t, JSONExp(table2ColJSON).CONTAINS(JSON(val), String("$.id")),
		"JSON_CONTAINS(table2.col_json, ?, ?)", string(json), "$.id")
}

func TestJSONExpression_CONTAINS_PATH(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).CONTAINS_PATH(false, String("$.id")),
		"JSON_CONTAINS_PATH(table2.col_json, 'one', ?)", "$.id")
	assertSerialize(t, JSONExp(table2ColJSON).CONTAINS_PATH(true, String("$.key1"), String("$.key2")),
		"JSON_CONTAINS_PATH(table2.col_json, 'all', ?, ?)", "$.key1", "$.key2")
}

func TestJSONExpression_DEPTH(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).DEPTH(), "JSON_DEPTH(table2.col_json)")
}

func TestJSONExpression_KEYS(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).KEYS(), "JSON_KEYS(table2.col_json)")
	assertSerialize(t, JSONExp(table2ColJSON).KEYS(String("$.id")),
		"JSON_KEYS(table2.col_json, ?)", "$.id")
}

func TestJSONExpression_LENGTH(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).LENGTH(), "JSON_LENGTH(table2.col_json)")
	assertSerialize(t, JSONExp(table2ColJSON).LENGTH(String("$.id")),
		"JSON_LENGTH(table2.col_json, ?)", "$.id")
}

func TestJSONExpression_MERGE_PATCH(t *testing.T) {
	val := []string{"val1", "val2"}
	json, _ := json.Marshal(val)

	assertSerialize(t, JSONExp(table2ColJSON).MERGE_PATCH(JSON(val)),
		"JSON_MERGE_PATCH(table2.col_json, ?)", string(json))
}

func TestJSONExpression_MERGE_PRESERVE(t *testing.T) {
	val := []string{"val1", "val2"}
	json, _ := json.Marshal(val)

	assertSerialize(t, JSONExp(table2ColJSON).MERGE_PRESERVE(JSON(val)),
		"JSON_MERGE_PRESERVE(table2.col_json, ?)", string(json))
}

func TestJSONExpression_OVERLAPS(t *testing.T) {
	val := []string{"val1", "val2"}
	json, _ := json.Marshal(val)

	assertSerialize(t, JSONExp(table2ColJSON).OVERLAPS(JSON(val)),
		"JSON_OVERLAPS(table2.col_json, ?)", string(json))
}

func TestJSONExpression_REMOVE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).REMOVE(String("$.id")),
		"JSON_REMOVE(table2.col_json, ?)", "$.id")
	assertSerialize(t, JSONExp(table2ColJSON).REMOVE(String("$.key1"), String("$.key2")),
		"JSON_REMOVE(table2.col_json, ?, ?)", "$.key1", "$.key2")
}

func TestJSONExpression_REPLACE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).REPLACE(String("$.id"), String("value")),
		"JSON_REPLACE(table2.col_json, ?, ?)", "$.id", "value")
	assertSerialize(t, JSONExp(table2ColJSON).REPLACE(String("$.key1"), String("value"), String("$.key2"), Int(2)),
		"JSON_REPLACE(table2.col_json, ?, ?, ?, ?)", "$.key1", "value", "$.key2", int64(2))
}

func TestJSONExpression_SCHEMA_VALID(t *testing.T) {
	schema := map[string]string{"type": "string"}
	json, _ := json.Marshal(schema)

	assertSerialize(t, JSONExp(table2ColJSON).SCHEMA_VALID(JSON(schema)),
		"JSON_SCHEMA_VALID(table2.col_json, ?)", string(json))
}

func TestJSONExpression_SCHEMA_VALIDATION_REPORT(t *testing.T) {
	schema := map[string]string{"type": "string"}
	json, _ := json.Marshal(schema)

	assertSerialize(t, JSONExp(table2ColJSON).SCHEMA_VALIDATION_REPORT(JSON(schema)),
		"JSON_SCHEMA_VALIDATION_REPORT(table2.col_json, ?)", string(json))
}

func TestJSONExpression_SEARCH(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).SEARCH(false, String("abc")),
		"JSON_SEARCH(table2.col_json, 'one', ?, ?)", "abc", "")
	assertSerialize(t,
		JSONExp(table2ColJSON).SEARCH(true, String("10"), String("$[1]"), String("$[2]")),
		"JSON_SEARCH(table2.col_json, 'all', ?, ?, ?, ?)", "10", "", "$[1]", "$[2]")
}

func TestJSONExpression_SET_VALUE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).SET_VALUE(String("$.id"), String("abc")),
		"JSON_SET(table2.col_json, ?, ?)", "$.id", "abc")
	assertSerialize(t,
		JSONExp(table2ColJSON).SET_VALUE(String("$.key1"), String("value"), String("$.key2"), Int(2)),
		"JSON_SET(table2.col_json, ?, ?, ?, ?)", "$.key1", "value", "$.key2", int64(2))
}

func TestJSONExpression_TYPE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).TYPE(), "JSON_TYPE(table2.col_json)")
}

func TestJSONExpression_EQ(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).EQ(String("value")),
		`(table2.col_json = CAST(? AS JSON))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).EQ(JSON(map[string]string{"key": "value"})),
		`(table2.col_json = ?)`, `{"key":"value"}`)
}

func TestJSONExpression_NOT_EQ(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).NOT_EQ(String("value")),
		`(table2.col_json != CAST(? AS JSON))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).NOT_EQ(JSON(map[string]string{"key": "value"})),
		`(table2.col_json != ?)`, `{"key":"value"}`)
}

func TestJSONExpression_LT(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).LT(String("value")),
		`(table2.col_json < CAST(? AS JSON))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).LT(JSON(map[string]string{"key": "value"})),
		`(table2.col_json < ?)`, `{"key":"value"}`)
}

func TestJSONExpression_LT_EQ(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).LT_EQ(String("value")),
		`(table2.col_json <= CAST(? AS JSON))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).LT_EQ(JSON(map[string]string{"key": "value"})),
		`(table2.col_json <= ?)`, `{"key":"value"}`)
}

func TestJSONExpression_GT(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).GT(String("value")),
		`(table2.col_json > CAST(? AS JSON))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).GT(JSON(map[string]string{"key": "value"})),
		`(table2.col_json > ?)`, `{"key":"value"}`)
}

func TestJSONExpression_GT_EQ(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).GT_EQ(String("value")),
		`(table2.col_json >= CAST(? AS JSON))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).GT_EQ(JSON(map[string]string{"key": "value"})),
		`(table2.col_json >= ?)`, `{"key":"value"}`)
}

func TestJSONExpression_IS_DISTINCT_FROM(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).IS_DISTINCT_FROM(String("value")),
		`(NOT(table2.col_json <=> ?))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).IS_DISTINCT_FROM(JSON(map[string]string{"key": "value"})),
		`(NOT(table2.col_json <=> ?))`, `{"key":"value"}`)
}

func TestJSONExpression_IS_NOT_DISTINCT_FROM(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).IS_NOT_DISTINCT_FROM(String("value")),
		`(table2.col_json <=> ?)`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).IS_NOT_DISTINCT_FROM(JSON(map[string]string{"key": "value"})),
		`(table2.col_json <=> ?)`, `{"key":"value"}`)
}

func TestJSONExpression_CONCAT(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).CONCAT(String("value")),
		`(CONCAT(table2.col_json, ?))`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).CONCAT(JSON(map[string]string{"key": "value"})),
		`(CONCAT(table2.col_json, ?))`, `{"key":"value"}`)
}

func TestJSONExpression_LIKE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).LIKE(String("value")),
		`(table2.col_json LIKE ?)`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).LIKE(JSON(map[string]string{"key": "value"})),
		`(table2.col_json LIKE ?)`, `{"key":"value"}`)
}

func TestJSONExpression_NOT_LIKE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).NOT_LIKE(String("value")),
		`(table2.col_json NOT LIKE ?)`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).NOT_LIKE(JSON(map[string]string{"key": "value"})),
		`(table2.col_json NOT LIKE ?)`, `{"key":"value"}`)
}

func TestJSONExpression_REGEXP_LIKE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).REGEXP_LIKE(String("value")),
		`(table2.col_json REGEXP ?)`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).REGEXP_LIKE(JSON(map[string]string{"key": "value"})),
		`(table2.col_json REGEXP ?)`, `{"key":"value"}`)
}

func TestJSONExpression_NOT_REGEXP_LIKE(t *testing.T) {
	assertSerialize(t, JSONExp(table2ColJSON).NOT_REGEXP_LIKE(String("value")),
		`(table2.col_json NOT REGEXP ?)`, "value")
	assertSerialize(t, JSONExp(table2ColJSON).NOT_REGEXP_LIKE(JSON(map[string]string{"key": "value"})),
		`(table2.col_json NOT REGEXP ?)`, `{"key":"value"}`)
}

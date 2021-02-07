package qrm

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSON []byte

func (j *JSON) Value() (driver.Value, error) {
	return *j, nil
}

func (j *JSON) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		if !json.Valid([]byte(src)) {
			return fmt.Errorf("Scan: invalid json: %s", string(src))
		}

		*j = []byte(src)

	case []byte:
		if !json.Valid(src) {
			return fmt.Errorf("Scan: invalid json: %s", string(src))
		}

		*j = src

	default:
		return fmt.Errorf("Scan: unable to scan type %T into JSON", src)
	}

	return nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}

	return j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}

	*j = append((*j)[0:0], data...)

	return nil
}

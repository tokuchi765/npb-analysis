package sqlwrapper

import (
	"database/sql"
	"encoding/json"
)

// NullFloat64 JSON拡張したNullFloat64
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON json変換処理
func (f NullFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Float64)
}

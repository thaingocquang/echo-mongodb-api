package testutil

import (
	"bytes"
	"encoding/json"
	"io"
)

type Verb struct {
	Method string
	Url    string
}

// BodyToReader ...
func BodyToReader(i interface{}) io.Reader {
	bodyMarshal, _ := json.Marshal(i)
	return bytes.NewReader(bodyMarshal)
}

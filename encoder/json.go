package encoder

import (
	"encoding/json"
)

var (
	JSONEncoder = new(jsonEncoder)
)

type jsonEncoder struct{}

func (e *jsonEncoder) Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

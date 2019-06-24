package assertive

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONSchemaInvalid(t *testing.T) {
	body := ioutil.NopCloser(bytes.NewBufferString(`{"foo":"baz","number": -1}`))
	res := &http.Response{Body: body}
	match := `{
  "title": "Example Schema",
  "type": "object",
  "properties": {
    "foo": {
      "type": "string"
    },
    "bar": {
      "type": "string"
    },
    "number": {
      "description": "A Number",
      "type": "integer",
      "minimum": 0
    }
  },
  "required": ["foo", "bar"]
}`
	err := JSONSchema(match)(res)
	assert.Error(t, err, "bar: bar is required")
	assert.Error(t, err, "number: number must be greater than or equal to 0")
}

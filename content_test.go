package assertive

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ItemsObject struct {
	Items []int
}

func TestContentIsValid(t *testing.T) {
	body := ioutil.NopCloser(bytes.NewBufferString(`[1,2,3]`))
	res := &http.Response{Body: body}

	match := "[1,2,3]"

	err := BodyContentMatches(match)(res)

	assert.Nil(t, err)
}

func TestContentIsInValid(t *testing.T) {
	body := ioutil.NopCloser(bytes.NewBufferString(`[1,2,3]`))
	res := &http.Response{Body: body}

	match := "[1,2,3,4]"

	err := BodyContentMatches(match)(res)

	assert.EqualError(t, err, "Expected Body to be [1,2,3,4] but was [1,2,3]")
}

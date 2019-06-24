package assertive

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func readBodyJSON(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	// Re-fill body reader stream after reading it
	res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, err
}

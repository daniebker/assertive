package assertive

import (
	"fmt"
	"net/http"
)

// ReturnsBody chexks the returned body matches the expected body
func (assertion *Assertion) ReturnsBody(content string) *Assertion {
	assertion.asserts = append(assertion.asserts, BodyContentMatches(content))
	return assertion
}

// BodyContentMatches Checks JSONCotent matches expected content
func BodyContentMatches(content string) Assert {
	return func(res *http.Response) error {
		buf, err := readBodyJSON(res)
		if err != nil {
			return err
		}

		if string(buf) != content {
			err = fmt.Errorf("Expected Body to be %s but was %s", content, string(buf))
			return err
		}

		return nil
	}
}

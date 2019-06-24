package assertive

import (
	"fmt"
	"net/http"
	"regexp"
)

// ReturnsContentTypeJSON checks the returned content type was JSON
func (assertion *Assertion) ReturnsContentTypeJSON() *Assertion {
	assertion.asserts = append(assertion.asserts, TypeEqual("application/json"))
	return assertion
}

// TypeEqual checks the response type equals the expected type
func TypeEqual(kind string) Assert {
	return func(res *http.Response) error {
		header := res.Header.Get("Content-Type")
		if match, _ := regexp.MatchString(kind, header); !match {
			return fmt.Errorf("Unexpected content type: '%s' should match '%s'", kind, header)
		}
		return nil
	}
}

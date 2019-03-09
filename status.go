package assertive

import (
	"fmt"
	"net/http"
)

// ReturnStatus {code} checks the response is the expected code.
func (assertion *Assertion) ReturnStatus(code int) *Assertion {
	assertion.asserts = append(assertion.asserts, StatusEqual(code))
	return assertion
}

// ReturnStatus200 checks the response was 200
func (assertion *Assertion) ReturnStatus200() *Assertion {
	assertion.asserts = append(assertion.asserts, StatusEqual(200))
	return assertion
}

// StatusEqual checks the response status equals the expected code
func StatusEqual(code int) Assert {
	return func(res *http.Response) error {
		if res.StatusCode != code {
			return fmt.Errorf("Expected status code: %d, got %d", code, res.StatusCode)
		}
		return nil
	}
}

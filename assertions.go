package assertive

import (
	"net/http"
	"testing"
)

// Assert the required interface for assertion functions
type Assert func(*http.Response) error

// Assertion struct to wrap asserions
type Assertion struct {
	api       *API
	testSuite *testing.T
	asserts   []Assert
}

// Should binds test context to api
func (api *API) Should(testSuite *testing.T) *Assertion {
	assertion := &Assertion{}
	assertion.testSuite = testSuite
	assertion.api = api
	return assertion
}

// Assert runs the assertions on the request
func (assertion *Assertion) Assert() error {
	res, err := assertion.api.DoRequest()

	if err != nil {
		assertion.testSuite.Error(err)
		return err
	}

	for _, assert := range assertion.asserts {
		err = assert(res)
		if err != nil {
			assertion.testSuite.Error(err)
			return err
		}
	}

	return nil
}

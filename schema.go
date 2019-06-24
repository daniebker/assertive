package assertive

import (
	"fmt"
	"net/http"

	"github.com/xeipuuv/gojsonschema"
)

// ReturnsSchemaMatching asserts the schema matches the one provided
func (assertion *Assertion) ReturnsSchemaMatching(schema string) *Assertion {
	assertion.asserts = append(assertion.asserts, JSONSchema(schema))
	return assertion
}

// JSONSchema validates the response body against a schema
func JSONSchema(schema string) Assert {
	return func(res *http.Response) error {
		buf, err := readBodyJSON(res)
		if err != nil {
			return err
		}

		loader := gojsonschema.NewStringLoader(schema)
		bodyLoader := gojsonschema.NewStringLoader(string(buf))

		result, err := gojsonschema.Validate(loader, bodyLoader)

		if err != nil {
			return err
		}

		if !result.Valid() {
			msg := "JSON document is not valid for the following reasons:\n"
			for _, detail := range result.Errors() {
				msg += fmt.Sprintf("\t- %s\n", detail)
			}
			return fmt.Errorf(msg)
		}

		return nil
	}
}

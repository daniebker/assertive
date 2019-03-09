# Assertive, GO.

A go package for making Fluent Assertions on APIs. 

## Example 

```go
	assertive.NewAPI("http://example.com").
		Get("/api/v1/Some/Path").
		WithHeader("Authorization", suite.JwtToken).
		Should(suite.T()).
		ReturnStatus200().
		Assert()
```
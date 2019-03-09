# Assertive

A go package for making Fluent Assertions on web APIs. 

## Example 

```go
	assertive.NewAPI("http://example.com").
		Get("/api/v1/Some/Path").
		WithHeader("Authorization", suite.JwtToken).
		Should(suite.T()).
		ReturnStatus200().
		Assert()
```

## DISCLAIMER

This is still very much a work in progress. Expect things to change with every commit until something stable is released. 
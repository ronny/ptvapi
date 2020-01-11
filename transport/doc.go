/*

Package transport provides a custom transport that automatically calculates
the request signature as required by the PTV Timetable API v3, and includes
both the `devid` and `signature` query param in all requests to the API.

## Example

``` go
	// Create a transport that will be wrapped by `SigningTransport`, here
	// we just create a default http transport from the
	// github.com/go-openapi/runtime/client package
	underlyingTransport := httptransport.New(
		apiclient.DefaultHost,
		apiclient.DefaultBasePath,
		apiclient.DefaultSchemes,
	)

	// Wrap the underlying transport and also supply devid and apikey that you
	// get from PTV
	transport := transport.NewSigningTransport(underlyingTransport,
		os.Getenv("DEVID"), os.Getenv("APIKEY"))

	// Create a new `client` (from github.com/ronny/ptvapi/v3/client pkg) with the
	// custom transport
	client := apiclient.New(transport, strfmt.Default)

	// Use `client` as you would normally, without `devid` and `signature` params
	resp, err := client.Routes.RoutesOneOrMoreRoutes(&routes.RoutesOneOrMoreRoutesParams{
		Context: context.Background(),
	})
```

See `examples/get-routes/` for a complete example of how to use this package.

*/
package transport

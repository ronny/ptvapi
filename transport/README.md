# transport
`import "github.com/ronny/ptvapi/transport"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
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

	// Create a new `client` (from github.com/ronny/ptvapi/client pkg) with the
	// custom transport
	client := apiclient.New(transport, strfmt.Default)

	// Use `client` as you would normally, without `devid` and `signature` params
	resp, err := client.Routes.RoutesOneOrMoreRoutes(&routes.RoutesOneOrMoreRoutesParams{
		Context: context.Background(),
	})

```

See `examples/get-routes/` for a complete example of how to use this package.

## <a name="pkg-imports">Imported Packages</a>

- [github.com/go-openapi/runtime](https://godoc.org/github.com/go-openapi/runtime)
- [github.com/go-openapi/strfmt](https://godoc.org/github.com/go-openapi/strfmt)

## <a name="pkg-index">Index</a>
* [func Signature(apiKey, requestURI string) (string, error)](#Signature)
* [type SigningTransport](#SigningTransport)
  * [func NewSigningTransport(transport runtime.ClientTransport, devID string, apiKey string) \*SigningTransport](#NewSigningTransport)
  * [func (t \*SigningTransport) Submit(op \*runtime.ClientOperation) (interface{}, error)](#SigningTransport.Submit)

#### <a name="pkg-files">Package files</a>
[doc.go](./doc.go) [signature.go](./signature.go) [transport.go](./transport.go)

## <a name="Signature">func</a> [Signature](./signature.go#L9)
``` go
func Signature(apiKey, requestURI string) (string, error)
```

## <a name="SigningTransport">type</a> [SigningTransport](./transport.go#L13-L16)
``` go
type SigningTransport struct {
    Transport runtime.ClientTransport
    AuthInfo  runtime.ClientAuthInfoWriter
}

```
SigningTransport wraps another transport (`Transport`) to automatically add
the `devid` and `signature` query params to every outgoing request as
required by the PTV Timetable API v3.

The `signature` query param is calculated from the request URI and `APIKey`.

### <a name="NewSigningTransport">func</a> [NewSigningTransport](./transport.go#L20)
``` go
func NewSigningTransport(transport runtime.ClientTransport, devID string, apiKey string) *SigningTransport
```

### <a name="SigningTransport.Submit">func</a> (\*SigningTransport) [Submit](./transport.go#L42)
``` go
func (t *SigningTransport) Submit(op *runtime.ClientOperation) (interface{}, error)
```

- - -
Generated by [godoc2ghmd](https://github.com/iflix/godoc2ghmd)
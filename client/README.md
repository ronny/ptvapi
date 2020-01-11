# client
`import "github.com/ronny/ptvapi/client"`

* [Overview](#pkg-overview)
* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-imports">Imported Packages</a>

- [github.com/go-openapi/runtime](https://godoc.org/github.com/go-openapi/runtime)
- [github.com/go-openapi/runtime/client](https://godoc.org/github.com/go-openapi/runtime/client)
- [github.com/go-openapi/strfmt](https://godoc.org/github.com/go-openapi/strfmt)
- [github.com/ronny/ptvapi/client/departures](./departures)
- [github.com/ronny/ptvapi/client/directions](./directions)
- [github.com/ronny/ptvapi/client/disruptions](./disruptions)
- [github.com/ronny/ptvapi/client/outlets](./outlets)
- [github.com/ronny/ptvapi/client/patterns](./patterns)
- [github.com/ronny/ptvapi/client/route_types](./route_types)
- [github.com/ronny/ptvapi/client/routes](./routes)
- [github.com/ronny/ptvapi/client/runs](./runs)
- [github.com/ronny/ptvapi/client/search](./search)
- [github.com/ronny/ptvapi/client/stops](./stops)

## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type PTVTimetableAPIV3](#PTVTimetableAPIV3)
  * [func New(transport runtime.ClientTransport, formats strfmt.Registry) \*PTVTimetableAPIV3](#New)
  * [func NewHTTPClient(formats strfmt.Registry) \*PTVTimetableAPIV3](#NewHTTPClient)
  * [func NewHTTPClientWithConfig(formats strfmt.Registry, cfg \*TransportConfig) \*PTVTimetableAPIV3](#NewHTTPClientWithConfig)
  * [func (c \*PTVTimetableAPIV3) SetTransport(transport runtime.ClientTransport)](#PTVTimetableAPIV3.SetTransport)
* [type TransportConfig](#TransportConfig)
  * [func DefaultTransportConfig() \*TransportConfig](#DefaultTransportConfig)
  * [func (cfg \*TransportConfig) WithBasePath(basePath string) \*TransportConfig](#TransportConfig.WithBasePath)
  * [func (cfg \*TransportConfig) WithHost(host string) \*TransportConfig](#TransportConfig.WithHost)
  * [func (cfg \*TransportConfig) WithSchemes(schemes []string) \*TransportConfig](#TransportConfig.WithSchemes)

#### <a name="pkg-files">Package files</a>
[doc.go](./doc.go) [ptv_timetable_api_v3_client.go](./ptv_timetable_api_v3_client.go) 

## <a name="pkg-constants">Constants</a>
``` go
const (
    // DefaultHost is the default Host
    // found in Meta (info) section of spec file
    DefaultHost string = "timetableapi.ptv.vic.gov.au"
    // DefaultBasePath is the default BasePath
    // found in Meta (info) section of spec file
    DefaultBasePath string = "/"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var Default = NewHTTPClient(nil)
```
Default PTV timetable API v3 HTTP client.

``` go
var DefaultSchemes = []string{"http", "https"}
```
DefaultSchemes are the default schemes found in Meta (info) section of spec file

## <a name="PTVTimetableAPIV3">type</a> [PTVTimetableAPIV3](./ptv_timetable_api_v3_client.go#L132-L154)
``` go
type PTVTimetableAPIV3 struct {
    Departures *departures.Client

    Directions *directions.Client

    Disruptions *disruptions.Client

    Outlets *outlets.Client

    Patterns *patterns.Client

    RouteTypes *route_types.Client

    Routes *routes.Client

    Runs *runs.Client

    Search *search.Client

    Stops *stops.Client

    Transport runtime.ClientTransport
}

```
PTVTimetableAPIV3 is a client for PTV timetable API v3

### <a name="New">func</a> [New](./ptv_timetable_api_v3_client.go#L60)
``` go
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PTVTimetableAPIV3
```
New creates a new PTV timetable API v3 client

### <a name="NewHTTPClient">func</a> [NewHTTPClient](./ptv_timetable_api_v3_client.go#L42)
``` go
func NewHTTPClient(formats strfmt.Registry) *PTVTimetableAPIV3
```
NewHTTPClient creates a new PTV timetable API v3 HTTP client.

### <a name="NewHTTPClientWithConfig">func</a> [NewHTTPClientWithConfig](./ptv_timetable_api_v3_client.go#L48)
``` go
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PTVTimetableAPIV3
```
NewHTTPClientWithConfig creates a new PTV timetable API v3 HTTP client,
using a customizable transport config.

### <a name="PTVTimetableAPIV3.SetTransport">func</a> (\*PTVTimetableAPIV3) [SetTransport](./ptv_timetable_api_v3_client.go#L157)
``` go
func (c *PTVTimetableAPIV3) SetTransport(transport runtime.ClientTransport)
```
SetTransport changes the transport on the client and all its subresources

## <a name="TransportConfig">type</a> [TransportConfig](./ptv_timetable_api_v3_client.go#L104-L108)
``` go
type TransportConfig struct {
    Host     string
    BasePath string
    Schemes  []string
}

```
TransportConfig contains the transport related info,
found in the meta section of the spec file.

### <a name="DefaultTransportConfig">func</a> [DefaultTransportConfig](./ptv_timetable_api_v3_client.go#L94)
``` go
func DefaultTransportConfig() *TransportConfig
```
DefaultTransportConfig creates a TransportConfig with the
default settings taken from the meta section of the spec file.

### <a name="TransportConfig.WithBasePath">func</a> (\*TransportConfig) [WithBasePath](./ptv_timetable_api_v3_client.go#L119)
``` go
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig
```
WithBasePath overrides the default basePath,
provided by the meta section of the spec file.

### <a name="TransportConfig.WithHost">func</a> (\*TransportConfig) [WithHost](./ptv_timetable_api_v3_client.go#L112)
``` go
func (cfg *TransportConfig) WithHost(host string) *TransportConfig
```
WithHost overrides the default host,
provided by the meta section of the spec file.

### <a name="TransportConfig.WithSchemes">func</a> (\*TransportConfig) [WithSchemes](./ptv_timetable_api_v3_client.go#L126)
``` go
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig
```
WithSchemes overrides the default schemes,
provided by the meta section of the spec file.

- - -
Generated by [godoc2ghmd](https://github.com/iflix/godoc2ghmd)
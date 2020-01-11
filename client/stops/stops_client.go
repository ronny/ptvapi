// Code generated by go-swagger; DO NOT EDIT.

package stops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new stops API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stops API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
StopsStopDetails views facilities at a specific stop metro and v line stations only
*/
func (a *Client) StopsStopDetails(params *StopsStopDetailsParams) (*StopsStopDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopsStopDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Stops_StopDetails",
		Method:             "GET",
		PathPattern:        "/v3/stops/{stop_id}/route_type/{route_type}",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopsStopDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopsStopDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Stops_StopDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopsStopsByGeolocation views all stops near a specific location
*/
func (a *Client) StopsStopsByGeolocation(params *StopsStopsByGeolocationParams) (*StopsStopsByGeolocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopsStopsByGeolocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Stops_StopsByGeolocation",
		Method:             "GET",
		PathPattern:        "/v3/stops/location/{latitude},{longitude}",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopsStopsByGeolocationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopsStopsByGeolocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Stops_StopsByGeolocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopsStopsForRoute views all stops on a specific route
*/
func (a *Client) StopsStopsForRoute(params *StopsStopsForRouteParams) (*StopsStopsForRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopsStopsForRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Stops_StopsForRoute",
		Method:             "GET",
		PathPattern:        "/v3/stops/route/{route_id}/route_type/{route_type}",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopsStopsForRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopsStopsForRouteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Stops_StopsForRoute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
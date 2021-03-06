// Code generated by go-swagger; DO NOT EDIT.

package disruptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new disruptions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for disruptions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DisruptionsGetAllDisruptions views all disruptions for all route types
*/
func (a *Client) DisruptionsGetAllDisruptions(params *DisruptionsGetAllDisruptionsParams) (*DisruptionsGetAllDisruptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisruptionsGetAllDisruptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Disruptions_GetAllDisruptions",
		Method:             "GET",
		PathPattern:        "/v3/disruptions",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisruptionsGetAllDisruptionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisruptionsGetAllDisruptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Disruptions_GetAllDisruptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisruptionsGetDisruptionByID views a specific disruption
*/
func (a *Client) DisruptionsGetDisruptionByID(params *DisruptionsGetDisruptionByIDParams) (*DisruptionsGetDisruptionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisruptionsGetDisruptionByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Disruptions_GetDisruptionById",
		Method:             "GET",
		PathPattern:        "/v3/disruptions/{disruption_id}",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisruptionsGetDisruptionByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisruptionsGetDisruptionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Disruptions_GetDisruptionById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisruptionsGetDisruptionModes gets all disruption modes
*/
func (a *Client) DisruptionsGetDisruptionModes(params *DisruptionsGetDisruptionModesParams) (*DisruptionsGetDisruptionModesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisruptionsGetDisruptionModesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Disruptions_GetDisruptionModes",
		Method:             "GET",
		PathPattern:        "/v3/disruptions/modes",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisruptionsGetDisruptionModesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisruptionsGetDisruptionModesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Disruptions_GetDisruptionModes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisruptionsGetDisruptionsByRoute views all disruptions for a particular route
*/
func (a *Client) DisruptionsGetDisruptionsByRoute(params *DisruptionsGetDisruptionsByRouteParams) (*DisruptionsGetDisruptionsByRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisruptionsGetDisruptionsByRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Disruptions_GetDisruptionsByRoute",
		Method:             "GET",
		PathPattern:        "/v3/disruptions/route/{route_id}",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisruptionsGetDisruptionsByRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisruptionsGetDisruptionsByRouteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Disruptions_GetDisruptionsByRoute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisruptionsGetDisruptionsByRouteAndStop views all disruptions for a particular route and stop
*/
func (a *Client) DisruptionsGetDisruptionsByRouteAndStop(params *DisruptionsGetDisruptionsByRouteAndStopParams) (*DisruptionsGetDisruptionsByRouteAndStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisruptionsGetDisruptionsByRouteAndStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Disruptions_GetDisruptionsByRouteAndStop",
		Method:             "GET",
		PathPattern:        "/v3/disruptions/route/{route_id}/stop/{stop_id}",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisruptionsGetDisruptionsByRouteAndStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisruptionsGetDisruptionsByRouteAndStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Disruptions_GetDisruptionsByRouteAndStop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisruptionsGetDisruptionsByStop views all disruptions for a particular stop
*/
func (a *Client) DisruptionsGetDisruptionsByStop(params *DisruptionsGetDisruptionsByStopParams) (*DisruptionsGetDisruptionsByStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisruptionsGetDisruptionsByStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Disruptions_GetDisruptionsByStop",
		Method:             "GET",
		PathPattern:        "/v3/disruptions/stop/{stop_id}",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisruptionsGetDisruptionsByStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisruptionsGetDisruptionsByStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Disruptions_GetDisruptionsByStop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package route_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/v3/models"
)

// RouteTypesGetRouteTypesReader is a Reader for the RouteTypesGetRouteTypes structure.
type RouteTypesGetRouteTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteTypesGetRouteTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteTypesGetRouteTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRouteTypesGetRouteTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRouteTypesGetRouteTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRouteTypesGetRouteTypesOK creates a RouteTypesGetRouteTypesOK with default headers values
func NewRouteTypesGetRouteTypesOK() *RouteTypesGetRouteTypesOK {
	return &RouteTypesGetRouteTypesOK{}
}

/*RouteTypesGetRouteTypesOK handles this case with default header values.

All route types (i.e. identifiers of transport modes) and their names.
*/
type RouteTypesGetRouteTypesOK struct {
	Payload *models.V3RouteTypesResponse
}

func (o *RouteTypesGetRouteTypesOK) Error() string {
	return fmt.Sprintf("[GET /v3/route_types][%d] routeTypesGetRouteTypesOK  %+v", 200, o.Payload)
}

func (o *RouteTypesGetRouteTypesOK) GetPayload() *models.V3RouteTypesResponse {
	return o.Payload
}

func (o *RouteTypesGetRouteTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3RouteTypesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteTypesGetRouteTypesBadRequest creates a RouteTypesGetRouteTypesBadRequest with default headers values
func NewRouteTypesGetRouteTypesBadRequest() *RouteTypesGetRouteTypesBadRequest {
	return &RouteTypesGetRouteTypesBadRequest{}
}

/*RouteTypesGetRouteTypesBadRequest handles this case with default header values.

Invalid Request
*/
type RouteTypesGetRouteTypesBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *RouteTypesGetRouteTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/route_types][%d] routeTypesGetRouteTypesBadRequest  %+v", 400, o.Payload)
}

func (o *RouteTypesGetRouteTypesBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RouteTypesGetRouteTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouteTypesGetRouteTypesForbidden creates a RouteTypesGetRouteTypesForbidden with default headers values
func NewRouteTypesGetRouteTypesForbidden() *RouteTypesGetRouteTypesForbidden {
	return &RouteTypesGetRouteTypesForbidden{}
}

/*RouteTypesGetRouteTypesForbidden handles this case with default header values.

Access Denied
*/
type RouteTypesGetRouteTypesForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *RouteTypesGetRouteTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/route_types][%d] routeTypesGetRouteTypesForbidden  %+v", 403, o.Payload)
}

func (o *RouteTypesGetRouteTypesForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RouteTypesGetRouteTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

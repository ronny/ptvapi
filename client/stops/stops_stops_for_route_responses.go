// Code generated by go-swagger; DO NOT EDIT.

package stops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/models"
)

// StopsStopsForRouteReader is a Reader for the StopsStopsForRoute structure.
type StopsStopsForRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopsStopsForRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopsStopsForRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopsStopsForRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopsStopsForRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopsStopsForRouteOK creates a StopsStopsForRouteOK with default headers values
func NewStopsStopsForRouteOK() *StopsStopsForRouteOK {
	return &StopsStopsForRouteOK{}
}

/*StopsStopsForRouteOK handles this case with default header values.

All stops on the specified route.
*/
type StopsStopsForRouteOK struct {
	Payload *models.V3StopsOnRouteResponse
}

func (o *StopsStopsForRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteOK  %+v", 200, o.Payload)
}

func (o *StopsStopsForRouteOK) GetPayload() *models.V3StopsOnRouteResponse {
	return o.Payload
}

func (o *StopsStopsForRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3StopsOnRouteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopsStopsForRouteBadRequest creates a StopsStopsForRouteBadRequest with default headers values
func NewStopsStopsForRouteBadRequest() *StopsStopsForRouteBadRequest {
	return &StopsStopsForRouteBadRequest{}
}

/*StopsStopsForRouteBadRequest handles this case with default header values.

Invalid Request
*/
type StopsStopsForRouteBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *StopsStopsForRouteBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteBadRequest  %+v", 400, o.Payload)
}

func (o *StopsStopsForRouteBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *StopsStopsForRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopsStopsForRouteForbidden creates a StopsStopsForRouteForbidden with default headers values
func NewStopsStopsForRouteForbidden() *StopsStopsForRouteForbidden {
	return &StopsStopsForRouteForbidden{}
}

/*StopsStopsForRouteForbidden handles this case with default header values.

Access Denied
*/
type StopsStopsForRouteForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *StopsStopsForRouteForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/stops/route/{route_id}/route_type/{route_type}][%d] stopsStopsForRouteForbidden  %+v", 403, o.Payload)
}

func (o *StopsStopsForRouteForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *StopsStopsForRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package disruptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/models"
)

// DisruptionsGetDisruptionsByRouteReader is a Reader for the DisruptionsGetDisruptionsByRoute structure.
type DisruptionsGetDisruptionsByRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisruptionsGetDisruptionsByRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisruptionsGetDisruptionsByRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisruptionsGetDisruptionsByRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisruptionsGetDisruptionsByRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDisruptionsGetDisruptionsByRouteOK creates a DisruptionsGetDisruptionsByRouteOK with default headers values
func NewDisruptionsGetDisruptionsByRouteOK() *DisruptionsGetDisruptionsByRouteOK {
	return &DisruptionsGetDisruptionsByRouteOK{}
}

/*DisruptionsGetDisruptionsByRouteOK handles this case with default header values.

All disruption information (if any exists) for the specified route.
*/
type DisruptionsGetDisruptionsByRouteOK struct {
	Payload *models.V3DisruptionsResponse
}

func (o *DisruptionsGetDisruptionsByRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/disruptions/route/{route_id}][%d] disruptionsGetDisruptionsByRouteOK  %+v", 200, o.Payload)
}

func (o *DisruptionsGetDisruptionsByRouteOK) GetPayload() *models.V3DisruptionsResponse {
	return o.Payload
}

func (o *DisruptionsGetDisruptionsByRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3DisruptionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisruptionsGetDisruptionsByRouteBadRequest creates a DisruptionsGetDisruptionsByRouteBadRequest with default headers values
func NewDisruptionsGetDisruptionsByRouteBadRequest() *DisruptionsGetDisruptionsByRouteBadRequest {
	return &DisruptionsGetDisruptionsByRouteBadRequest{}
}

/*DisruptionsGetDisruptionsByRouteBadRequest handles this case with default header values.

Invalid Request
*/
type DisruptionsGetDisruptionsByRouteBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *DisruptionsGetDisruptionsByRouteBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/disruptions/route/{route_id}][%d] disruptionsGetDisruptionsByRouteBadRequest  %+v", 400, o.Payload)
}

func (o *DisruptionsGetDisruptionsByRouteBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DisruptionsGetDisruptionsByRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisruptionsGetDisruptionsByRouteForbidden creates a DisruptionsGetDisruptionsByRouteForbidden with default headers values
func NewDisruptionsGetDisruptionsByRouteForbidden() *DisruptionsGetDisruptionsByRouteForbidden {
	return &DisruptionsGetDisruptionsByRouteForbidden{}
}

/*DisruptionsGetDisruptionsByRouteForbidden handles this case with default header values.

Access Denied
*/
type DisruptionsGetDisruptionsByRouteForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *DisruptionsGetDisruptionsByRouteForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/disruptions/route/{route_id}][%d] disruptionsGetDisruptionsByRouteForbidden  %+v", 403, o.Payload)
}

func (o *DisruptionsGetDisruptionsByRouteForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DisruptionsGetDisruptionsByRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
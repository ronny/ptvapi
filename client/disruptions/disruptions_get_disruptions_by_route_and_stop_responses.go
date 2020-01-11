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

// DisruptionsGetDisruptionsByRouteAndStopReader is a Reader for the DisruptionsGetDisruptionsByRouteAndStop structure.
type DisruptionsGetDisruptionsByRouteAndStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisruptionsGetDisruptionsByRouteAndStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisruptionsGetDisruptionsByRouteAndStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisruptionsGetDisruptionsByRouteAndStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisruptionsGetDisruptionsByRouteAndStopForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDisruptionsGetDisruptionsByRouteAndStopOK creates a DisruptionsGetDisruptionsByRouteAndStopOK with default headers values
func NewDisruptionsGetDisruptionsByRouteAndStopOK() *DisruptionsGetDisruptionsByRouteAndStopOK {
	return &DisruptionsGetDisruptionsByRouteAndStopOK{}
}

/*DisruptionsGetDisruptionsByRouteAndStopOK handles this case with default header values.

All disruption information (if any exists) for the specified route and stop.
*/
type DisruptionsGetDisruptionsByRouteAndStopOK struct {
	Payload *models.V3DisruptionsResponse
}

func (o *DisruptionsGetDisruptionsByRouteAndStopOK) Error() string {
	return fmt.Sprintf("[GET /v3/disruptions/route/{route_id}/stop/{stop_id}][%d] disruptionsGetDisruptionsByRouteAndStopOK  %+v", 200, o.Payload)
}

func (o *DisruptionsGetDisruptionsByRouteAndStopOK) GetPayload() *models.V3DisruptionsResponse {
	return o.Payload
}

func (o *DisruptionsGetDisruptionsByRouteAndStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3DisruptionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisruptionsGetDisruptionsByRouteAndStopBadRequest creates a DisruptionsGetDisruptionsByRouteAndStopBadRequest with default headers values
func NewDisruptionsGetDisruptionsByRouteAndStopBadRequest() *DisruptionsGetDisruptionsByRouteAndStopBadRequest {
	return &DisruptionsGetDisruptionsByRouteAndStopBadRequest{}
}

/*DisruptionsGetDisruptionsByRouteAndStopBadRequest handles this case with default header values.

Invalid Request
*/
type DisruptionsGetDisruptionsByRouteAndStopBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *DisruptionsGetDisruptionsByRouteAndStopBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/disruptions/route/{route_id}/stop/{stop_id}][%d] disruptionsGetDisruptionsByRouteAndStopBadRequest  %+v", 400, o.Payload)
}

func (o *DisruptionsGetDisruptionsByRouteAndStopBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DisruptionsGetDisruptionsByRouteAndStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisruptionsGetDisruptionsByRouteAndStopForbidden creates a DisruptionsGetDisruptionsByRouteAndStopForbidden with default headers values
func NewDisruptionsGetDisruptionsByRouteAndStopForbidden() *DisruptionsGetDisruptionsByRouteAndStopForbidden {
	return &DisruptionsGetDisruptionsByRouteAndStopForbidden{}
}

/*DisruptionsGetDisruptionsByRouteAndStopForbidden handles this case with default header values.

Access Denied
*/
type DisruptionsGetDisruptionsByRouteAndStopForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *DisruptionsGetDisruptionsByRouteAndStopForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/disruptions/route/{route_id}/stop/{stop_id}][%d] disruptionsGetDisruptionsByRouteAndStopForbidden  %+v", 403, o.Payload)
}

func (o *DisruptionsGetDisruptionsByRouteAndStopForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DisruptionsGetDisruptionsByRouteAndStopForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
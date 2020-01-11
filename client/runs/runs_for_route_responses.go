// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/v3/models"
)

// RunsForRouteReader is a Reader for the RunsForRoute structure.
type RunsForRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsForRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsForRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunsForRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRunsForRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunsForRouteOK creates a RunsForRouteOK with default headers values
func NewRunsForRouteOK() *RunsForRouteOK {
	return &RunsForRouteOK{}
}

/*RunsForRouteOK handles this case with default header values.

All trip/service run details for the specified route ID.
*/
type RunsForRouteOK struct {
	Payload *models.V3RunsResponse
}

func (o *RunsForRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/runs/route/{route_id}][%d] runsForRouteOK  %+v", 200, o.Payload)
}

func (o *RunsForRouteOK) GetPayload() *models.V3RunsResponse {
	return o.Payload
}

func (o *RunsForRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3RunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsForRouteBadRequest creates a RunsForRouteBadRequest with default headers values
func NewRunsForRouteBadRequest() *RunsForRouteBadRequest {
	return &RunsForRouteBadRequest{}
}

/*RunsForRouteBadRequest handles this case with default header values.

Invalid Request
*/
type RunsForRouteBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *RunsForRouteBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/runs/route/{route_id}][%d] runsForRouteBadRequest  %+v", 400, o.Payload)
}

func (o *RunsForRouteBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RunsForRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsForRouteForbidden creates a RunsForRouteForbidden with default headers values
func NewRunsForRouteForbidden() *RunsForRouteForbidden {
	return &RunsForRouteForbidden{}
}

/*RunsForRouteForbidden handles this case with default header values.

Access Denied
*/
type RunsForRouteForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *RunsForRouteForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/runs/route/{route_id}][%d] runsForRouteForbidden  %+v", 403, o.Payload)
}

func (o *RunsForRouteForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RunsForRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

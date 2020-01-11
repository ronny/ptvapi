// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/models"
)

// RunsForRunAndRouteTypeReader is a Reader for the RunsForRunAndRouteType structure.
type RunsForRunAndRouteTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsForRunAndRouteTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsForRunAndRouteTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunsForRunAndRouteTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRunsForRunAndRouteTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunsForRunAndRouteTypeOK creates a RunsForRunAndRouteTypeOK with default headers values
func NewRunsForRunAndRouteTypeOK() *RunsForRunAndRouteTypeOK {
	return &RunsForRunAndRouteTypeOK{}
}

/*RunsForRunAndRouteTypeOK handles this case with default header values.

The trip/service run details for the run ID and route type specified.
*/
type RunsForRunAndRouteTypeOK struct {
	Payload *models.V3RunResponse
}

func (o *RunsForRunAndRouteTypeOK) Error() string {
	return fmt.Sprintf("[GET /v3/runs/{run_id}/route_type/{route_type}][%d] runsForRunAndRouteTypeOK  %+v", 200, o.Payload)
}

func (o *RunsForRunAndRouteTypeOK) GetPayload() *models.V3RunResponse {
	return o.Payload
}

func (o *RunsForRunAndRouteTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3RunResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsForRunAndRouteTypeBadRequest creates a RunsForRunAndRouteTypeBadRequest with default headers values
func NewRunsForRunAndRouteTypeBadRequest() *RunsForRunAndRouteTypeBadRequest {
	return &RunsForRunAndRouteTypeBadRequest{}
}

/*RunsForRunAndRouteTypeBadRequest handles this case with default header values.

Invalid Request
*/
type RunsForRunAndRouteTypeBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *RunsForRunAndRouteTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/runs/{run_id}/route_type/{route_type}][%d] runsForRunAndRouteTypeBadRequest  %+v", 400, o.Payload)
}

func (o *RunsForRunAndRouteTypeBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RunsForRunAndRouteTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsForRunAndRouteTypeForbidden creates a RunsForRunAndRouteTypeForbidden with default headers values
func NewRunsForRunAndRouteTypeForbidden() *RunsForRunAndRouteTypeForbidden {
	return &RunsForRunAndRouteTypeForbidden{}
}

/*RunsForRunAndRouteTypeForbidden handles this case with default header values.

Access Denied
*/
type RunsForRunAndRouteTypeForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *RunsForRunAndRouteTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/runs/{run_id}/route_type/{route_type}][%d] runsForRunAndRouteTypeForbidden  %+v", 403, o.Payload)
}

func (o *RunsForRunAndRouteTypeForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RunsForRunAndRouteTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
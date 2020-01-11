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

// RunsForRouteAndRouteTypeReader is a Reader for the RunsForRouteAndRouteType structure.
type RunsForRouteAndRouteTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunsForRouteAndRouteTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunsForRouteAndRouteTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunsForRouteAndRouteTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRunsForRouteAndRouteTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunsForRouteAndRouteTypeOK creates a RunsForRouteAndRouteTypeOK with default headers values
func NewRunsForRouteAndRouteTypeOK() *RunsForRouteAndRouteTypeOK {
	return &RunsForRouteAndRouteTypeOK{}
}

/*RunsForRouteAndRouteTypeOK handles this case with default header values.

All trip/service run details for the specified route ID and route type.
*/
type RunsForRouteAndRouteTypeOK struct {
	Payload *models.V3RunsResponse
}

func (o *RunsForRouteAndRouteTypeOK) Error() string {
	return fmt.Sprintf("[GET /v3/runs/route/{route_id}/route_type/{route_type}][%d] runsForRouteAndRouteTypeOK  %+v", 200, o.Payload)
}

func (o *RunsForRouteAndRouteTypeOK) GetPayload() *models.V3RunsResponse {
	return o.Payload
}

func (o *RunsForRouteAndRouteTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3RunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsForRouteAndRouteTypeBadRequest creates a RunsForRouteAndRouteTypeBadRequest with default headers values
func NewRunsForRouteAndRouteTypeBadRequest() *RunsForRouteAndRouteTypeBadRequest {
	return &RunsForRouteAndRouteTypeBadRequest{}
}

/*RunsForRouteAndRouteTypeBadRequest handles this case with default header values.

Invalid Request
*/
type RunsForRouteAndRouteTypeBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *RunsForRouteAndRouteTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/runs/route/{route_id}/route_type/{route_type}][%d] runsForRouteAndRouteTypeBadRequest  %+v", 400, o.Payload)
}

func (o *RunsForRouteAndRouteTypeBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RunsForRouteAndRouteTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunsForRouteAndRouteTypeForbidden creates a RunsForRouteAndRouteTypeForbidden with default headers values
func NewRunsForRouteAndRouteTypeForbidden() *RunsForRouteAndRouteTypeForbidden {
	return &RunsForRouteAndRouteTypeForbidden{}
}

/*RunsForRouteAndRouteTypeForbidden handles this case with default header values.

Access Denied
*/
type RunsForRouteAndRouteTypeForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *RunsForRouteAndRouteTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/runs/route/{route_id}/route_type/{route_type}][%d] runsForRouteAndRouteTypeForbidden  %+v", 403, o.Payload)
}

func (o *RunsForRouteAndRouteTypeForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *RunsForRouteAndRouteTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

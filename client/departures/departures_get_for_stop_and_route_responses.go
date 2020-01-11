// Code generated by go-swagger; DO NOT EDIT.

package departures

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/v3/models"
)

// DeparturesGetForStopAndRouteReader is a Reader for the DeparturesGetForStopAndRoute structure.
type DeparturesGetForStopAndRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeparturesGetForStopAndRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeparturesGetForStopAndRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeparturesGetForStopAndRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeparturesGetForStopAndRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeparturesGetForStopAndRouteOK creates a DeparturesGetForStopAndRouteOK with default headers values
func NewDeparturesGetForStopAndRouteOK() *DeparturesGetForStopAndRouteOK {
	return &DeparturesGetForStopAndRouteOK{}
}

/*DeparturesGetForStopAndRouteOK handles this case with default header values.

Service departures from the specified stop for the specified route (and route type); departures are timetabled and real-time (if applicable).
*/
type DeparturesGetForStopAndRouteOK struct {
	Payload *models.V3DeparturesResponse
}

func (o *DeparturesGetForStopAndRouteOK) Error() string {
	return fmt.Sprintf("[GET /v3/departures/route_type/{route_type}/stop/{stop_id}/route/{route_id}][%d] departuresGetForStopAndRouteOK  %+v", 200, o.Payload)
}

func (o *DeparturesGetForStopAndRouteOK) GetPayload() *models.V3DeparturesResponse {
	return o.Payload
}

func (o *DeparturesGetForStopAndRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3DeparturesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeparturesGetForStopAndRouteBadRequest creates a DeparturesGetForStopAndRouteBadRequest with default headers values
func NewDeparturesGetForStopAndRouteBadRequest() *DeparturesGetForStopAndRouteBadRequest {
	return &DeparturesGetForStopAndRouteBadRequest{}
}

/*DeparturesGetForStopAndRouteBadRequest handles this case with default header values.

Invalid Request
*/
type DeparturesGetForStopAndRouteBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *DeparturesGetForStopAndRouteBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/departures/route_type/{route_type}/stop/{stop_id}/route/{route_id}][%d] departuresGetForStopAndRouteBadRequest  %+v", 400, o.Payload)
}

func (o *DeparturesGetForStopAndRouteBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DeparturesGetForStopAndRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeparturesGetForStopAndRouteForbidden creates a DeparturesGetForStopAndRouteForbidden with default headers values
func NewDeparturesGetForStopAndRouteForbidden() *DeparturesGetForStopAndRouteForbidden {
	return &DeparturesGetForStopAndRouteForbidden{}
}

/*DeparturesGetForStopAndRouteForbidden handles this case with default header values.

Access Denied
*/
type DeparturesGetForStopAndRouteForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *DeparturesGetForStopAndRouteForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/departures/route_type/{route_type}/stop/{stop_id}/route/{route_id}][%d] departuresGetForStopAndRouteForbidden  %+v", 403, o.Payload)
}

func (o *DeparturesGetForStopAndRouteForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *DeparturesGetForStopAndRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

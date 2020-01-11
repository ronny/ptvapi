// Code generated by go-swagger; DO NOT EDIT.

package patterns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/models"
)

// PatternsGetPatternByRunReader is a Reader for the PatternsGetPatternByRun structure.
type PatternsGetPatternByRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatternsGetPatternByRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatternsGetPatternByRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatternsGetPatternByRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatternsGetPatternByRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatternsGetPatternByRunOK creates a PatternsGetPatternByRunOK with default headers values
func NewPatternsGetPatternByRunOK() *PatternsGetPatternByRunOK {
	return &PatternsGetPatternByRunOK{}
}

/*PatternsGetPatternByRunOK handles this case with default header values.

The stopping pattern of the specified trip/service run and route type.
*/
type PatternsGetPatternByRunOK struct {
	Payload *models.V3StoppingPattern
}

func (o *PatternsGetPatternByRunOK) Error() string {
	return fmt.Sprintf("[GET /v3/pattern/run/{run_id}/route_type/{route_type}][%d] patternsGetPatternByRunOK  %+v", 200, o.Payload)
}

func (o *PatternsGetPatternByRunOK) GetPayload() *models.V3StoppingPattern {
	return o.Payload
}

func (o *PatternsGetPatternByRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3StoppingPattern)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatternsGetPatternByRunBadRequest creates a PatternsGetPatternByRunBadRequest with default headers values
func NewPatternsGetPatternByRunBadRequest() *PatternsGetPatternByRunBadRequest {
	return &PatternsGetPatternByRunBadRequest{}
}

/*PatternsGetPatternByRunBadRequest handles this case with default header values.

Invalid Request
*/
type PatternsGetPatternByRunBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *PatternsGetPatternByRunBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/pattern/run/{run_id}/route_type/{route_type}][%d] patternsGetPatternByRunBadRequest  %+v", 400, o.Payload)
}

func (o *PatternsGetPatternByRunBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *PatternsGetPatternByRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatternsGetPatternByRunForbidden creates a PatternsGetPatternByRunForbidden with default headers values
func NewPatternsGetPatternByRunForbidden() *PatternsGetPatternByRunForbidden {
	return &PatternsGetPatternByRunForbidden{}
}

/*PatternsGetPatternByRunForbidden handles this case with default header values.

Access Denied
*/
type PatternsGetPatternByRunForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *PatternsGetPatternByRunForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/pattern/run/{run_id}/route_type/{route_type}][%d] patternsGetPatternByRunForbidden  %+v", 403, o.Payload)
}

func (o *PatternsGetPatternByRunForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *PatternsGetPatternByRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

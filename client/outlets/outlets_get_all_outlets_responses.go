// Code generated by go-swagger; DO NOT EDIT.

package outlets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/v3/models"
)

// OutletsGetAllOutletsReader is a Reader for the OutletsGetAllOutlets structure.
type OutletsGetAllOutletsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutletsGetAllOutletsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOutletsGetAllOutletsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOutletsGetAllOutletsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOutletsGetAllOutletsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOutletsGetAllOutletsOK creates a OutletsGetAllOutletsOK with default headers values
func NewOutletsGetAllOutletsOK() *OutletsGetAllOutletsOK {
	return &OutletsGetAllOutletsOK{}
}

/*OutletsGetAllOutletsOK handles this case with default header values.

Ticket outlets.
*/
type OutletsGetAllOutletsOK struct {
	Payload *models.V3OutletResponse
}

func (o *OutletsGetAllOutletsOK) Error() string {
	return fmt.Sprintf("[GET /v3/outlets][%d] outletsGetAllOutletsOK  %+v", 200, o.Payload)
}

func (o *OutletsGetAllOutletsOK) GetPayload() *models.V3OutletResponse {
	return o.Payload
}

func (o *OutletsGetAllOutletsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3OutletResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutletsGetAllOutletsBadRequest creates a OutletsGetAllOutletsBadRequest with default headers values
func NewOutletsGetAllOutletsBadRequest() *OutletsGetAllOutletsBadRequest {
	return &OutletsGetAllOutletsBadRequest{}
}

/*OutletsGetAllOutletsBadRequest handles this case with default header values.

Invalid Request
*/
type OutletsGetAllOutletsBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *OutletsGetAllOutletsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/outlets][%d] outletsGetAllOutletsBadRequest  %+v", 400, o.Payload)
}

func (o *OutletsGetAllOutletsBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *OutletsGetAllOutletsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutletsGetAllOutletsForbidden creates a OutletsGetAllOutletsForbidden with default headers values
func NewOutletsGetAllOutletsForbidden() *OutletsGetAllOutletsForbidden {
	return &OutletsGetAllOutletsForbidden{}
}

/*OutletsGetAllOutletsForbidden handles this case with default header values.

Access Denied
*/
type OutletsGetAllOutletsForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *OutletsGetAllOutletsForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/outlets][%d] outletsGetAllOutletsForbidden  %+v", 403, o.Payload)
}

func (o *OutletsGetAllOutletsForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *OutletsGetAllOutletsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ronny/ptvapi/v3/models"
)

// SearchSearchReader is a Reader for the SearchSearch structure.
type SearchSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchSearchOK creates a SearchSearchOK with default headers values
func NewSearchSearchOK() *SearchSearchOK {
	return &SearchSearchOK{}
}

/*SearchSearchOK handles this case with default header values.

Stops, routes and myki ticket outlets that contain the search term (note: stops and routes are ordered by route_type by default).
*/
type SearchSearchOK struct {
	Payload *models.V3SearchResult
}

func (o *SearchSearchOK) Error() string {
	return fmt.Sprintf("[GET /v3/search/{search_term}][%d] searchSearchOK  %+v", 200, o.Payload)
}

func (o *SearchSearchOK) GetPayload() *models.V3SearchResult {
	return o.Payload
}

func (o *SearchSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3SearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSearchBadRequest creates a SearchSearchBadRequest with default headers values
func NewSearchSearchBadRequest() *SearchSearchBadRequest {
	return &SearchSearchBadRequest{}
}

/*SearchSearchBadRequest handles this case with default header values.

Invalid Request
*/
type SearchSearchBadRequest struct {
	Payload *models.V3ErrorResponse
}

func (o *SearchSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/search/{search_term}][%d] searchSearchBadRequest  %+v", 400, o.Payload)
}

func (o *SearchSearchBadRequest) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *SearchSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchSearchForbidden creates a SearchSearchForbidden with default headers values
func NewSearchSearchForbidden() *SearchSearchForbidden {
	return &SearchSearchForbidden{}
}

/*SearchSearchForbidden handles this case with default header values.

Access Denied
*/
type SearchSearchForbidden struct {
	Payload *models.V3ErrorResponse
}

func (o *SearchSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/search/{search_term}][%d] searchSearchForbidden  %+v", 403, o.Payload)
}

func (o *SearchSearchForbidden) GetPayload() *models.V3ErrorResponse {
	return o.Payload
}

func (o *SearchSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V3ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

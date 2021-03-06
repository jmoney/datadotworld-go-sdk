// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// GetRequestsReader is a Reader for the GetRequests structure.
type GetRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRequestsOK creates a GetRequestsOK with default headers values
func NewGetRequestsOK() *GetRequestsOK {
	return &GetRequestsOK{}
}

/* GetRequestsOK describes a response with status code 200, with default header values.

Successfully retrieved dataset
*/
type GetRequestsOK struct {
	Payload *models.PaginatedResultsDto
}

func (o *GetRequestsOK) Error() string {
	return fmt.Sprintf("[GET /requests/owner/{agentid}][%d] getRequestsOK  %+v", 200, o.Payload)
}
func (o *GetRequestsOK) GetPayload() *models.PaginatedResultsDto {
	return o.Payload
}

func (o *GetRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedResultsDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestsBadRequest creates a GetRequestsBadRequest with default headers values
func NewGetRequestsBadRequest() *GetRequestsBadRequest {
	return &GetRequestsBadRequest{}
}

/* GetRequestsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetRequestsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetRequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /requests/owner/{agentid}][%d] getRequestsBadRequest  %+v", 400, o.Payload)
}
func (o *GetRequestsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestsUnauthorized creates a GetRequestsUnauthorized with default headers values
func NewGetRequestsUnauthorized() *GetRequestsUnauthorized {
	return &GetRequestsUnauthorized{}
}

/* GetRequestsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRequestsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /requests/owner/{agentid}][%d] getRequestsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetRequestsUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestsForbidden creates a GetRequestsForbidden with default headers values
func NewGetRequestsForbidden() *GetRequestsForbidden {
	return &GetRequestsForbidden{}
}

/* GetRequestsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRequestsForbidden struct {
	Payload *models.ErrorMessage
}

func (o *GetRequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /requests/owner/{agentid}][%d] getRequestsForbidden  %+v", 403, o.Payload)
}
func (o *GetRequestsForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestsNotFound creates a GetRequestsNotFound with default headers values
func NewGetRequestsNotFound() *GetRequestsNotFound {
	return &GetRequestsNotFound{}
}

/* GetRequestsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetRequestsNotFound struct {
	Payload *models.ErrorMessage
}

func (o *GetRequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /requests/owner/{agentid}][%d] getRequestsNotFound  %+v", 404, o.Payload)
}
func (o *GetRequestsNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestsInternalServerError creates a GetRequestsInternalServerError with default headers values
func NewGetRequestsInternalServerError() *GetRequestsInternalServerError {
	return &GetRequestsInternalServerError{}
}

/* GetRequestsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetRequestsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /requests/owner/{agentid}][%d] getRequestsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRequestsInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// GetProjectsByOwnerReader is a Reader for the GetProjectsByOwner structure.
type GetProjectsByOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsByOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsByOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsByOwnerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsByOwnerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsByOwnerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectsByOwnerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsByOwnerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsByOwnerOK creates a GetProjectsByOwnerOK with default headers values
func NewGetProjectsByOwnerOK() *GetProjectsByOwnerOK {
	return &GetProjectsByOwnerOK{}
}

/* GetProjectsByOwnerOK describes a response with status code 200, with default header values.

Successfully retrieved project
*/
type GetProjectsByOwnerOK struct {
	Payload *models.PaginatedProjectResults
}

func (o *GetProjectsByOwnerOK) Error() string {
	return fmt.Sprintf("[GET /projects/{owner}][%d] getProjectsByOwnerOK  %+v", 200, o.Payload)
}
func (o *GetProjectsByOwnerOK) GetPayload() *models.PaginatedProjectResults {
	return o.Payload
}

func (o *GetProjectsByOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedProjectResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsByOwnerBadRequest creates a GetProjectsByOwnerBadRequest with default headers values
func NewGetProjectsByOwnerBadRequest() *GetProjectsByOwnerBadRequest {
	return &GetProjectsByOwnerBadRequest{}
}

/* GetProjectsByOwnerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetProjectsByOwnerBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetProjectsByOwnerBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{owner}][%d] getProjectsByOwnerBadRequest  %+v", 400, o.Payload)
}
func (o *GetProjectsByOwnerBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetProjectsByOwnerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsByOwnerUnauthorized creates a GetProjectsByOwnerUnauthorized with default headers values
func NewGetProjectsByOwnerUnauthorized() *GetProjectsByOwnerUnauthorized {
	return &GetProjectsByOwnerUnauthorized{}
}

/* GetProjectsByOwnerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectsByOwnerUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetProjectsByOwnerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{owner}][%d] getProjectsByOwnerUnauthorized  %+v", 401, o.Payload)
}
func (o *GetProjectsByOwnerUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetProjectsByOwnerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsByOwnerForbidden creates a GetProjectsByOwnerForbidden with default headers values
func NewGetProjectsByOwnerForbidden() *GetProjectsByOwnerForbidden {
	return &GetProjectsByOwnerForbidden{}
}

/* GetProjectsByOwnerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectsByOwnerForbidden struct {
	Payload *models.ErrorMessage
}

func (o *GetProjectsByOwnerForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{owner}][%d] getProjectsByOwnerForbidden  %+v", 403, o.Payload)
}
func (o *GetProjectsByOwnerForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetProjectsByOwnerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsByOwnerNotFound creates a GetProjectsByOwnerNotFound with default headers values
func NewGetProjectsByOwnerNotFound() *GetProjectsByOwnerNotFound {
	return &GetProjectsByOwnerNotFound{}
}

/* GetProjectsByOwnerNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetProjectsByOwnerNotFound struct {
	Payload *models.ErrorMessage
}

func (o *GetProjectsByOwnerNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{owner}][%d] getProjectsByOwnerNotFound  %+v", 404, o.Payload)
}
func (o *GetProjectsByOwnerNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetProjectsByOwnerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsByOwnerInternalServerError creates a GetProjectsByOwnerInternalServerError with default headers values
func NewGetProjectsByOwnerInternalServerError() *GetProjectsByOwnerInternalServerError {
	return &GetProjectsByOwnerInternalServerError{}
}

/* GetProjectsByOwnerInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetProjectsByOwnerInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetProjectsByOwnerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{owner}][%d] getProjectsByOwnerInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProjectsByOwnerInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetProjectsByOwnerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// FetchProjectsReader is a Reader for the FetchProjects structure.
type FetchProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFetchProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFetchProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFetchProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFetchProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFetchProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFetchProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFetchProjectsOK creates a FetchProjectsOK with default headers values
func NewFetchProjectsOK() *FetchProjectsOK {
	return &FetchProjectsOK{}
}

/* FetchProjectsOK describes a response with status code 200, with default header values.

successful operation
*/
type FetchProjectsOK struct {
	Payload *models.PaginatedProjectResults
}

func (o *FetchProjectsOK) Error() string {
	return fmt.Sprintf("[GET /user/projects/own][%d] fetchProjectsOK  %+v", 200, o.Payload)
}
func (o *FetchProjectsOK) GetPayload() *models.PaginatedProjectResults {
	return o.Payload
}

func (o *FetchProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedProjectResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchProjectsBadRequest creates a FetchProjectsBadRequest with default headers values
func NewFetchProjectsBadRequest() *FetchProjectsBadRequest {
	return &FetchProjectsBadRequest{}
}

/* FetchProjectsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type FetchProjectsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *FetchProjectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /user/projects/own][%d] fetchProjectsBadRequest  %+v", 400, o.Payload)
}
func (o *FetchProjectsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchProjectsUnauthorized creates a FetchProjectsUnauthorized with default headers values
func NewFetchProjectsUnauthorized() *FetchProjectsUnauthorized {
	return &FetchProjectsUnauthorized{}
}

/* FetchProjectsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FetchProjectsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *FetchProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/projects/own][%d] fetchProjectsUnauthorized  %+v", 401, o.Payload)
}
func (o *FetchProjectsUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchProjectsForbidden creates a FetchProjectsForbidden with default headers values
func NewFetchProjectsForbidden() *FetchProjectsForbidden {
	return &FetchProjectsForbidden{}
}

/* FetchProjectsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FetchProjectsForbidden struct {
	Payload *models.ErrorMessage
}

func (o *FetchProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /user/projects/own][%d] fetchProjectsForbidden  %+v", 403, o.Payload)
}
func (o *FetchProjectsForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchProjectsNotFound creates a FetchProjectsNotFound with default headers values
func NewFetchProjectsNotFound() *FetchProjectsNotFound {
	return &FetchProjectsNotFound{}
}

/* FetchProjectsNotFound describes a response with status code 404, with default header values.

Not found
*/
type FetchProjectsNotFound struct {
	Payload *models.ErrorMessage
}

func (o *FetchProjectsNotFound) Error() string {
	return fmt.Sprintf("[GET /user/projects/own][%d] fetchProjectsNotFound  %+v", 404, o.Payload)
}
func (o *FetchProjectsNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchProjectsInternalServerError creates a FetchProjectsInternalServerError with default headers values
func NewFetchProjectsInternalServerError() *FetchProjectsInternalServerError {
	return &FetchProjectsInternalServerError{}
}

/* FetchProjectsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type FetchProjectsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *FetchProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/projects/own][%d] fetchProjectsInternalServerError  %+v", 500, o.Payload)
}
func (o *FetchProjectsInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
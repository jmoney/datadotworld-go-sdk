// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// SimpleSearchResourcesReader is a Reader for the SimpleSearchResources structure.
type SimpleSearchResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SimpleSearchResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSimpleSearchResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSimpleSearchResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSimpleSearchResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSimpleSearchResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSimpleSearchResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSimpleSearchResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSimpleSearchResourcesOK creates a SimpleSearchResourcesOK with default headers values
func NewSimpleSearchResourcesOK() *SimpleSearchResourcesOK {
	return &SimpleSearchResourcesOK{}
}

/* SimpleSearchResourcesOK describes a response with status code 200, with default header values.

successful operation
*/
type SimpleSearchResourcesOK struct {
	Payload *models.PaginatedGenericResults
}

func (o *SimpleSearchResourcesOK) Error() string {
	return fmt.Sprintf("[POST /search/resources][%d] simpleSearchResourcesOK  %+v", 200, o.Payload)
}
func (o *SimpleSearchResourcesOK) GetPayload() *models.PaginatedGenericResults {
	return o.Payload
}

func (o *SimpleSearchResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedGenericResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSimpleSearchResourcesBadRequest creates a SimpleSearchResourcesBadRequest with default headers values
func NewSimpleSearchResourcesBadRequest() *SimpleSearchResourcesBadRequest {
	return &SimpleSearchResourcesBadRequest{}
}

/* SimpleSearchResourcesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SimpleSearchResourcesBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *SimpleSearchResourcesBadRequest) Error() string {
	return fmt.Sprintf("[POST /search/resources][%d] simpleSearchResourcesBadRequest  %+v", 400, o.Payload)
}
func (o *SimpleSearchResourcesBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SimpleSearchResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSimpleSearchResourcesUnauthorized creates a SimpleSearchResourcesUnauthorized with default headers values
func NewSimpleSearchResourcesUnauthorized() *SimpleSearchResourcesUnauthorized {
	return &SimpleSearchResourcesUnauthorized{}
}

/* SimpleSearchResourcesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SimpleSearchResourcesUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *SimpleSearchResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /search/resources][%d] simpleSearchResourcesUnauthorized  %+v", 401, o.Payload)
}
func (o *SimpleSearchResourcesUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SimpleSearchResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSimpleSearchResourcesForbidden creates a SimpleSearchResourcesForbidden with default headers values
func NewSimpleSearchResourcesForbidden() *SimpleSearchResourcesForbidden {
	return &SimpleSearchResourcesForbidden{}
}

/* SimpleSearchResourcesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SimpleSearchResourcesForbidden struct {
	Payload *models.ErrorMessage
}

func (o *SimpleSearchResourcesForbidden) Error() string {
	return fmt.Sprintf("[POST /search/resources][%d] simpleSearchResourcesForbidden  %+v", 403, o.Payload)
}
func (o *SimpleSearchResourcesForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SimpleSearchResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSimpleSearchResourcesNotFound creates a SimpleSearchResourcesNotFound with default headers values
func NewSimpleSearchResourcesNotFound() *SimpleSearchResourcesNotFound {
	return &SimpleSearchResourcesNotFound{}
}

/* SimpleSearchResourcesNotFound describes a response with status code 404, with default header values.

Not found
*/
type SimpleSearchResourcesNotFound struct {
	Payload *models.ErrorMessage
}

func (o *SimpleSearchResourcesNotFound) Error() string {
	return fmt.Sprintf("[POST /search/resources][%d] simpleSearchResourcesNotFound  %+v", 404, o.Payload)
}
func (o *SimpleSearchResourcesNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SimpleSearchResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSimpleSearchResourcesInternalServerError creates a SimpleSearchResourcesInternalServerError with default headers values
func NewSimpleSearchResourcesInternalServerError() *SimpleSearchResourcesInternalServerError {
	return &SimpleSearchResourcesInternalServerError{}
}

/* SimpleSearchResourcesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SimpleSearchResourcesInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *SimpleSearchResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /search/resources][%d] simpleSearchResourcesInternalServerError  %+v", 500, o.Payload)
}
func (o *SimpleSearchResourcesInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SimpleSearchResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

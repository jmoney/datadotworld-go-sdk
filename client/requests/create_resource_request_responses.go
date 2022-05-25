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

// CreateResourceRequestReader is a Reader for the CreateResourceRequest structure.
type CreateResourceRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourceRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateResourceRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateResourceRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateResourceRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateResourceRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateResourceRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateResourceRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateResourceRequestOK creates a CreateResourceRequestOK with default headers values
func NewCreateResourceRequestOK() *CreateResourceRequestOK {
	return &CreateResourceRequestOK{}
}

/* CreateResourceRequestOK describes a response with status code 200, with default header values.

Successfully retrieved dataset
*/
type CreateResourceRequestOK struct {
	Payload *models.ResourceRequestDto
}

func (o *CreateResourceRequestOK) Error() string {
	return fmt.Sprintf("[POST /requests/suggest][%d] createResourceRequestOK  %+v", 200, o.Payload)
}
func (o *CreateResourceRequestOK) GetPayload() *models.ResourceRequestDto {
	return o.Payload
}

func (o *CreateResourceRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceRequestDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestBadRequest creates a CreateResourceRequestBadRequest with default headers values
func NewCreateResourceRequestBadRequest() *CreateResourceRequestBadRequest {
	return &CreateResourceRequestBadRequest{}
}

/* CreateResourceRequestBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateResourceRequestBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *CreateResourceRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /requests/suggest][%d] createResourceRequestBadRequest  %+v", 400, o.Payload)
}
func (o *CreateResourceRequestBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateResourceRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestUnauthorized creates a CreateResourceRequestUnauthorized with default headers values
func NewCreateResourceRequestUnauthorized() *CreateResourceRequestUnauthorized {
	return &CreateResourceRequestUnauthorized{}
}

/* CreateResourceRequestUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateResourceRequestUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *CreateResourceRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /requests/suggest][%d] createResourceRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateResourceRequestUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateResourceRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestForbidden creates a CreateResourceRequestForbidden with default headers values
func NewCreateResourceRequestForbidden() *CreateResourceRequestForbidden {
	return &CreateResourceRequestForbidden{}
}

/* CreateResourceRequestForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateResourceRequestForbidden struct {
	Payload *models.ErrorMessage
}

func (o *CreateResourceRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /requests/suggest][%d] createResourceRequestForbidden  %+v", 403, o.Payload)
}
func (o *CreateResourceRequestForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateResourceRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestNotFound creates a CreateResourceRequestNotFound with default headers values
func NewCreateResourceRequestNotFound() *CreateResourceRequestNotFound {
	return &CreateResourceRequestNotFound{}
}

/* CreateResourceRequestNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateResourceRequestNotFound struct {
	Payload *models.ErrorMessage
}

func (o *CreateResourceRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /requests/suggest][%d] createResourceRequestNotFound  %+v", 404, o.Payload)
}
func (o *CreateResourceRequestNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateResourceRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceRequestInternalServerError creates a CreateResourceRequestInternalServerError with default headers values
func NewCreateResourceRequestInternalServerError() *CreateResourceRequestInternalServerError {
	return &CreateResourceRequestInternalServerError{}
}

/* CreateResourceRequestInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateResourceRequestInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *CreateResourceRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /requests/suggest][%d] createResourceRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateResourceRequestInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateResourceRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// DeleteConnectionReader is a Reader for the DeleteConnection structure.
type DeleteConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConnectionOK creates a DeleteConnectionOK with default headers values
func NewDeleteConnectionOK() *DeleteConnectionOK {
	return &DeleteConnectionOK{}
}

/* DeleteConnectionOK describes a response with status code 200, with default header values.

Connection has been successfully deleted.
*/
type DeleteConnectionOK struct {
	Payload *models.SuccessMessage
}

func (o *DeleteConnectionOK) Error() string {
	return fmt.Sprintf("[DELETE /connections/{owner}/{id}][%d] deleteConnectionOK  %+v", 200, o.Payload)
}
func (o *DeleteConnectionOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *DeleteConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionBadRequest creates a DeleteConnectionBadRequest with default headers values
func NewDeleteConnectionBadRequest() *DeleteConnectionBadRequest {
	return &DeleteConnectionBadRequest{}
}

/* DeleteConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteConnectionBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *DeleteConnectionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /connections/{owner}/{id}][%d] deleteConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteConnectionBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionUnauthorized creates a DeleteConnectionUnauthorized with default headers values
func NewDeleteConnectionUnauthorized() *DeleteConnectionUnauthorized {
	return &DeleteConnectionUnauthorized{}
}

/* DeleteConnectionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteConnectionUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *DeleteConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /connections/{owner}/{id}][%d] deleteConnectionUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteConnectionUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionForbidden creates a DeleteConnectionForbidden with default headers values
func NewDeleteConnectionForbidden() *DeleteConnectionForbidden {
	return &DeleteConnectionForbidden{}
}

/* DeleteConnectionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteConnectionForbidden struct {
	Payload *models.ErrorMessage
}

func (o *DeleteConnectionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /connections/{owner}/{id}][%d] deleteConnectionForbidden  %+v", 403, o.Payload)
}
func (o *DeleteConnectionForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionNotFound creates a DeleteConnectionNotFound with default headers values
func NewDeleteConnectionNotFound() *DeleteConnectionNotFound {
	return &DeleteConnectionNotFound{}
}

/* DeleteConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteConnectionNotFound struct {
	Payload *models.ErrorMessage
}

func (o *DeleteConnectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /connections/{owner}/{id}][%d] deleteConnectionNotFound  %+v", 404, o.Payload)
}
func (o *DeleteConnectionNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionInternalServerError creates a DeleteConnectionInternalServerError with default headers values
func NewDeleteConnectionInternalServerError() *DeleteConnectionInternalServerError {
	return &DeleteConnectionInternalServerError{}
}

/* DeleteConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteConnectionInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *DeleteConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /connections/{owner}/{id}][%d] deleteConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteConnectionInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

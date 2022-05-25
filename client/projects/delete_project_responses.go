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

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/* DeleteProjectOK describes a response with status code 200, with default header values.

Project has been successfully deleted.
*/
type DeleteProjectOK struct {
	Payload *models.SuccessMessage
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}][%d] deleteProjectOK  %+v", 200, o.Payload)
}
func (o *DeleteProjectOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectBadRequest creates a DeleteProjectBadRequest with default headers values
func NewDeleteProjectBadRequest() *DeleteProjectBadRequest {
	return &DeleteProjectBadRequest{}
}

/* DeleteProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteProjectBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *DeleteProjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}][%d] deleteProjectBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteProjectBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectUnauthorized creates a DeleteProjectUnauthorized with default headers values
func NewDeleteProjectUnauthorized() *DeleteProjectUnauthorized {
	return &DeleteProjectUnauthorized{}
}

/* DeleteProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteProjectUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *DeleteProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}][%d] deleteProjectUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteProjectUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/* DeleteProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteProjectForbidden struct {
	Payload *models.ErrorMessage
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}][%d] deleteProjectForbidden  %+v", 403, o.Payload)
}
func (o *DeleteProjectForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectNotFound creates a DeleteProjectNotFound with default headers values
func NewDeleteProjectNotFound() *DeleteProjectNotFound {
	return &DeleteProjectNotFound{}
}

/* DeleteProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteProjectNotFound struct {
	Payload *models.ErrorMessage
}

func (o *DeleteProjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}][%d] deleteProjectNotFound  %+v", 404, o.Payload)
}
func (o *DeleteProjectNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProjectInternalServerError creates a DeleteProjectInternalServerError with default headers values
func NewDeleteProjectInternalServerError() *DeleteProjectInternalServerError {
	return &DeleteProjectInternalServerError{}
}

/* DeleteProjectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteProjectInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *DeleteProjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}][%d] deleteProjectInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteProjectInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

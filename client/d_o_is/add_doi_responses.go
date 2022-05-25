// Code generated by go-swagger; DO NOT EDIT.

package d_o_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// AddDoiReader is a Reader for the AddDoi structure.
type AddDoiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDoiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDoiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddDoiBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddDoiUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddDoiForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddDoiNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddDoiInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddDoiOK creates a AddDoiOK with default headers values
func NewAddDoiOK() *AddDoiOK {
	return &AddDoiOK{}
}

/* AddDoiOK describes a response with status code 200, with default header values.

DOI successfully added to dataset.
*/
type AddDoiOK struct {
	Payload *models.SuccessMessage
}

func (o *AddDoiOK) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}/dois/{doi}][%d] addDoiOK  %+v", 200, o.Payload)
}
func (o *AddDoiOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *AddDoiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDoiBadRequest creates a AddDoiBadRequest with default headers values
func NewAddDoiBadRequest() *AddDoiBadRequest {
	return &AddDoiBadRequest{}
}

/* AddDoiBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddDoiBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *AddDoiBadRequest) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}/dois/{doi}][%d] addDoiBadRequest  %+v", 400, o.Payload)
}
func (o *AddDoiBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AddDoiBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDoiUnauthorized creates a AddDoiUnauthorized with default headers values
func NewAddDoiUnauthorized() *AddDoiUnauthorized {
	return &AddDoiUnauthorized{}
}

/* AddDoiUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddDoiUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *AddDoiUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}/dois/{doi}][%d] addDoiUnauthorized  %+v", 401, o.Payload)
}
func (o *AddDoiUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AddDoiUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDoiForbidden creates a AddDoiForbidden with default headers values
func NewAddDoiForbidden() *AddDoiForbidden {
	return &AddDoiForbidden{}
}

/* AddDoiForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddDoiForbidden struct {
	Payload *models.ErrorMessage
}

func (o *AddDoiForbidden) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}/dois/{doi}][%d] addDoiForbidden  %+v", 403, o.Payload)
}
func (o *AddDoiForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AddDoiForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDoiNotFound creates a AddDoiNotFound with default headers values
func NewAddDoiNotFound() *AddDoiNotFound {
	return &AddDoiNotFound{}
}

/* AddDoiNotFound describes a response with status code 404, with default header values.

Not found
*/
type AddDoiNotFound struct {
	Payload *models.ErrorMessage
}

func (o *AddDoiNotFound) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}/dois/{doi}][%d] addDoiNotFound  %+v", 404, o.Payload)
}
func (o *AddDoiNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AddDoiNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDoiInternalServerError creates a AddDoiInternalServerError with default headers values
func NewAddDoiInternalServerError() *AddDoiInternalServerError {
	return &AddDoiInternalServerError{}
}

/* AddDoiInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type AddDoiInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *AddDoiInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}/dois/{doi}][%d] addDoiInternalServerError  %+v", 500, o.Payload)
}
func (o *AddDoiInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AddDoiInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

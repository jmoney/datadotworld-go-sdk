// Code generated by go-swagger; DO NOT EDIT.

package datasets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// CreateDatasetReader is a Reader for the CreateDataset structure.
type CreateDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDatasetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDatasetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDatasetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDatasetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDatasetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateDatasetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDatasetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDatasetOK creates a CreateDatasetOK with default headers values
func NewCreateDatasetOK() *CreateDatasetOK {
	return &CreateDatasetOK{}
}

/* CreateDatasetOK describes a response with status code 200, with default header values.

Dataset created successfully.
*/
type CreateDatasetOK struct {
	Payload *models.CreateDatasetResponse
}

func (o *CreateDatasetOK) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetOK  %+v", 200, o.Payload)
}
func (o *CreateDatasetOK) GetPayload() *models.CreateDatasetResponse {
	return o.Payload
}

func (o *CreateDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDatasetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetBadRequest creates a CreateDatasetBadRequest with default headers values
func NewCreateDatasetBadRequest() *CreateDatasetBadRequest {
	return &CreateDatasetBadRequest{}
}

/* CreateDatasetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDatasetBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *CreateDatasetBadRequest) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDatasetBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateDatasetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetUnauthorized creates a CreateDatasetUnauthorized with default headers values
func NewCreateDatasetUnauthorized() *CreateDatasetUnauthorized {
	return &CreateDatasetUnauthorized{}
}

/* CreateDatasetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateDatasetUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *CreateDatasetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateDatasetUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateDatasetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetForbidden creates a CreateDatasetForbidden with default headers values
func NewCreateDatasetForbidden() *CreateDatasetForbidden {
	return &CreateDatasetForbidden{}
}

/* CreateDatasetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateDatasetForbidden struct {
	Payload *models.ErrorMessage
}

func (o *CreateDatasetForbidden) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetForbidden  %+v", 403, o.Payload)
}
func (o *CreateDatasetForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateDatasetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetNotFound creates a CreateDatasetNotFound with default headers values
func NewCreateDatasetNotFound() *CreateDatasetNotFound {
	return &CreateDatasetNotFound{}
}

/* CreateDatasetNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateDatasetNotFound struct {
	Payload *models.ErrorMessage
}

func (o *CreateDatasetNotFound) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetNotFound  %+v", 404, o.Payload)
}
func (o *CreateDatasetNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateDatasetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetConflict creates a CreateDatasetConflict with default headers values
func NewCreateDatasetConflict() *CreateDatasetConflict {
	return &CreateDatasetConflict{}
}

/* CreateDatasetConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateDatasetConflict struct {
	Payload *models.ErrorMessage
}

func (o *CreateDatasetConflict) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetConflict  %+v", 409, o.Payload)
}
func (o *CreateDatasetConflict) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateDatasetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetUnprocessableEntity creates a CreateDatasetUnprocessableEntity with default headers values
func NewCreateDatasetUnprocessableEntity() *CreateDatasetUnprocessableEntity {
	return &CreateDatasetUnprocessableEntity{}
}

/* CreateDatasetUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type CreateDatasetUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *CreateDatasetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateDatasetUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateDatasetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDatasetInternalServerError creates a CreateDatasetInternalServerError with default headers values
func NewCreateDatasetInternalServerError() *CreateDatasetInternalServerError {
	return &CreateDatasetInternalServerError{}
}

/* CreateDatasetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateDatasetInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *CreateDatasetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}][%d] createDatasetInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDatasetInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateDatasetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

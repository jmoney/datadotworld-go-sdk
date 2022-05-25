// Code generated by go-swagger; DO NOT EDIT.

package streams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// PatchStreamSchemaReader is a Reader for the PatchStreamSchema structure.
type PatchStreamSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStreamSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchStreamSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchStreamSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchStreamSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchStreamSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchStreamSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPatchStreamSchemaUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchStreamSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchStreamSchemaOK creates a PatchStreamSchemaOK with default headers values
func NewPatchStreamSchemaOK() *PatchStreamSchemaOK {
	return &PatchStreamSchemaOK{}
}

/* PatchStreamSchemaOK describes a response with status code 200, with default header values.

Stream updated successfully.
*/
type PatchStreamSchemaOK struct {
	Payload *models.SuccessMessage
}

func (o *PatchStreamSchemaOK) Error() string {
	return fmt.Sprintf("[PATCH /streams/{owner}/{id}/{streamId}/schema][%d] patchStreamSchemaOK  %+v", 200, o.Payload)
}
func (o *PatchStreamSchemaOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *PatchStreamSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStreamSchemaBadRequest creates a PatchStreamSchemaBadRequest with default headers values
func NewPatchStreamSchemaBadRequest() *PatchStreamSchemaBadRequest {
	return &PatchStreamSchemaBadRequest{}
}

/* PatchStreamSchemaBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PatchStreamSchemaBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *PatchStreamSchemaBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /streams/{owner}/{id}/{streamId}/schema][%d] patchStreamSchemaBadRequest  %+v", 400, o.Payload)
}
func (o *PatchStreamSchemaBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PatchStreamSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStreamSchemaUnauthorized creates a PatchStreamSchemaUnauthorized with default headers values
func NewPatchStreamSchemaUnauthorized() *PatchStreamSchemaUnauthorized {
	return &PatchStreamSchemaUnauthorized{}
}

/* PatchStreamSchemaUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PatchStreamSchemaUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *PatchStreamSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /streams/{owner}/{id}/{streamId}/schema][%d] patchStreamSchemaUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchStreamSchemaUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PatchStreamSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStreamSchemaForbidden creates a PatchStreamSchemaForbidden with default headers values
func NewPatchStreamSchemaForbidden() *PatchStreamSchemaForbidden {
	return &PatchStreamSchemaForbidden{}
}

/* PatchStreamSchemaForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchStreamSchemaForbidden struct {
	Payload *models.ErrorMessage
}

func (o *PatchStreamSchemaForbidden) Error() string {
	return fmt.Sprintf("[PATCH /streams/{owner}/{id}/{streamId}/schema][%d] patchStreamSchemaForbidden  %+v", 403, o.Payload)
}
func (o *PatchStreamSchemaForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PatchStreamSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStreamSchemaNotFound creates a PatchStreamSchemaNotFound with default headers values
func NewPatchStreamSchemaNotFound() *PatchStreamSchemaNotFound {
	return &PatchStreamSchemaNotFound{}
}

/* PatchStreamSchemaNotFound describes a response with status code 404, with default header values.

Not found
*/
type PatchStreamSchemaNotFound struct {
	Payload *models.ErrorMessage
}

func (o *PatchStreamSchemaNotFound) Error() string {
	return fmt.Sprintf("[PATCH /streams/{owner}/{id}/{streamId}/schema][%d] patchStreamSchemaNotFound  %+v", 404, o.Payload)
}
func (o *PatchStreamSchemaNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PatchStreamSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStreamSchemaUnprocessableEntity creates a PatchStreamSchemaUnprocessableEntity with default headers values
func NewPatchStreamSchemaUnprocessableEntity() *PatchStreamSchemaUnprocessableEntity {
	return &PatchStreamSchemaUnprocessableEntity{}
}

/* PatchStreamSchemaUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type PatchStreamSchemaUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *PatchStreamSchemaUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /streams/{owner}/{id}/{streamId}/schema][%d] patchStreamSchemaUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PatchStreamSchemaUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PatchStreamSchemaUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStreamSchemaInternalServerError creates a PatchStreamSchemaInternalServerError with default headers values
func NewPatchStreamSchemaInternalServerError() *PatchStreamSchemaInternalServerError {
	return &PatchStreamSchemaInternalServerError{}
}

/* PatchStreamSchemaInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type PatchStreamSchemaInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *PatchStreamSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /streams/{owner}/{id}/{streamId}/schema][%d] patchStreamSchemaInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchStreamSchemaInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PatchStreamSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

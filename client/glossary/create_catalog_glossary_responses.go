// Code generated by go-swagger; DO NOT EDIT.

package glossary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// CreateCatalogGlossaryReader is a Reader for the CreateCatalogGlossary structure.
type CreateCatalogGlossaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCatalogGlossaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCatalogGlossaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCatalogGlossaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCatalogGlossaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCatalogGlossaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCatalogGlossaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateCatalogGlossaryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCatalogGlossaryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCatalogGlossaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCatalogGlossaryOK creates a CreateCatalogGlossaryOK with default headers values
func NewCreateCatalogGlossaryOK() *CreateCatalogGlossaryOK {
	return &CreateCatalogGlossaryOK{}
}

/* CreateCatalogGlossaryOK describes a response with status code 200, with default header values.

Glossary Metadata created successfully.
*/
type CreateCatalogGlossaryOK struct {
	Payload *models.CreateResponse
}

func (o *CreateCatalogGlossaryOK) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryOK  %+v", 200, o.Payload)
}
func (o *CreateCatalogGlossaryOK) GetPayload() *models.CreateResponse {
	return o.Payload
}

func (o *CreateCatalogGlossaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogGlossaryBadRequest creates a CreateCatalogGlossaryBadRequest with default headers values
func NewCreateCatalogGlossaryBadRequest() *CreateCatalogGlossaryBadRequest {
	return &CreateCatalogGlossaryBadRequest{}
}

/* CreateCatalogGlossaryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateCatalogGlossaryBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogGlossaryBadRequest) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryBadRequest  %+v", 400, o.Payload)
}
func (o *CreateCatalogGlossaryBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogGlossaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogGlossaryUnauthorized creates a CreateCatalogGlossaryUnauthorized with default headers values
func NewCreateCatalogGlossaryUnauthorized() *CreateCatalogGlossaryUnauthorized {
	return &CreateCatalogGlossaryUnauthorized{}
}

/* CreateCatalogGlossaryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateCatalogGlossaryUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogGlossaryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateCatalogGlossaryUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogGlossaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogGlossaryForbidden creates a CreateCatalogGlossaryForbidden with default headers values
func NewCreateCatalogGlossaryForbidden() *CreateCatalogGlossaryForbidden {
	return &CreateCatalogGlossaryForbidden{}
}

/* CreateCatalogGlossaryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateCatalogGlossaryForbidden struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogGlossaryForbidden) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryForbidden  %+v", 403, o.Payload)
}
func (o *CreateCatalogGlossaryForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogGlossaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogGlossaryNotFound creates a CreateCatalogGlossaryNotFound with default headers values
func NewCreateCatalogGlossaryNotFound() *CreateCatalogGlossaryNotFound {
	return &CreateCatalogGlossaryNotFound{}
}

/* CreateCatalogGlossaryNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateCatalogGlossaryNotFound struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogGlossaryNotFound) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryNotFound  %+v", 404, o.Payload)
}
func (o *CreateCatalogGlossaryNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogGlossaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogGlossaryConflict creates a CreateCatalogGlossaryConflict with default headers values
func NewCreateCatalogGlossaryConflict() *CreateCatalogGlossaryConflict {
	return &CreateCatalogGlossaryConflict{}
}

/* CreateCatalogGlossaryConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateCatalogGlossaryConflict struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogGlossaryConflict) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryConflict  %+v", 409, o.Payload)
}
func (o *CreateCatalogGlossaryConflict) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogGlossaryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogGlossaryUnprocessableEntity creates a CreateCatalogGlossaryUnprocessableEntity with default headers values
func NewCreateCatalogGlossaryUnprocessableEntity() *CreateCatalogGlossaryUnprocessableEntity {
	return &CreateCatalogGlossaryUnprocessableEntity{}
}

/* CreateCatalogGlossaryUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type CreateCatalogGlossaryUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogGlossaryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateCatalogGlossaryUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogGlossaryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogGlossaryInternalServerError creates a CreateCatalogGlossaryInternalServerError with default headers values
func NewCreateCatalogGlossaryInternalServerError() *CreateCatalogGlossaryInternalServerError {
	return &CreateCatalogGlossaryInternalServerError{}
}

/* CreateCatalogGlossaryInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateCatalogGlossaryInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogGlossaryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /metadata/glossary/{owner}][%d] createCatalogGlossaryInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateCatalogGlossaryInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogGlossaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

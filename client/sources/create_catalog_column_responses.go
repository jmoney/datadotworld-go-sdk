// Code generated by go-swagger; DO NOT EDIT.

package sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// CreateCatalogColumnReader is a Reader for the CreateCatalogColumn structure.
type CreateCatalogColumnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCatalogColumnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCatalogColumnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCatalogColumnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCatalogColumnUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCatalogColumnForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCatalogColumnNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateCatalogColumnConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCatalogColumnUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCatalogColumnInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCatalogColumnOK creates a CreateCatalogColumnOK with default headers values
func NewCreateCatalogColumnOK() *CreateCatalogColumnOK {
	return &CreateCatalogColumnOK{}
}

/* CreateCatalogColumnOK describes a response with status code 200, with default header values.

Column Metadata created successfully.
*/
type CreateCatalogColumnOK struct {
	Payload *models.CreateResponse
}

func (o *CreateCatalogColumnOK) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnOK  %+v", 200, o.Payload)
}
func (o *CreateCatalogColumnOK) GetPayload() *models.CreateResponse {
	return o.Payload
}

func (o *CreateCatalogColumnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogColumnBadRequest creates a CreateCatalogColumnBadRequest with default headers values
func NewCreateCatalogColumnBadRequest() *CreateCatalogColumnBadRequest {
	return &CreateCatalogColumnBadRequest{}
}

/* CreateCatalogColumnBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateCatalogColumnBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogColumnBadRequest) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnBadRequest  %+v", 400, o.Payload)
}
func (o *CreateCatalogColumnBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogColumnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogColumnUnauthorized creates a CreateCatalogColumnUnauthorized with default headers values
func NewCreateCatalogColumnUnauthorized() *CreateCatalogColumnUnauthorized {
	return &CreateCatalogColumnUnauthorized{}
}

/* CreateCatalogColumnUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateCatalogColumnUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogColumnUnauthorized) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateCatalogColumnUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogColumnUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogColumnForbidden creates a CreateCatalogColumnForbidden with default headers values
func NewCreateCatalogColumnForbidden() *CreateCatalogColumnForbidden {
	return &CreateCatalogColumnForbidden{}
}

/* CreateCatalogColumnForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateCatalogColumnForbidden struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogColumnForbidden) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnForbidden  %+v", 403, o.Payload)
}
func (o *CreateCatalogColumnForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogColumnForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogColumnNotFound creates a CreateCatalogColumnNotFound with default headers values
func NewCreateCatalogColumnNotFound() *CreateCatalogColumnNotFound {
	return &CreateCatalogColumnNotFound{}
}

/* CreateCatalogColumnNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateCatalogColumnNotFound struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogColumnNotFound) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnNotFound  %+v", 404, o.Payload)
}
func (o *CreateCatalogColumnNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogColumnNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogColumnConflict creates a CreateCatalogColumnConflict with default headers values
func NewCreateCatalogColumnConflict() *CreateCatalogColumnConflict {
	return &CreateCatalogColumnConflict{}
}

/* CreateCatalogColumnConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateCatalogColumnConflict struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogColumnConflict) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnConflict  %+v", 409, o.Payload)
}
func (o *CreateCatalogColumnConflict) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogColumnConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogColumnUnprocessableEntity creates a CreateCatalogColumnUnprocessableEntity with default headers values
func NewCreateCatalogColumnUnprocessableEntity() *CreateCatalogColumnUnprocessableEntity {
	return &CreateCatalogColumnUnprocessableEntity{}
}

/* CreateCatalogColumnUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type CreateCatalogColumnUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogColumnUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateCatalogColumnUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogColumnUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCatalogColumnInternalServerError creates a CreateCatalogColumnInternalServerError with default headers values
func NewCreateCatalogColumnInternalServerError() *CreateCatalogColumnInternalServerError {
	return &CreateCatalogColumnInternalServerError{}
}

/* CreateCatalogColumnInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateCatalogColumnInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *CreateCatalogColumnInternalServerError) Error() string {
	return fmt.Sprintf("[POST /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}/columns][%d] createCatalogColumnInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateCatalogColumnInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateCatalogColumnInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

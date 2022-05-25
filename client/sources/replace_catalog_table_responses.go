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

// ReplaceCatalogTableReader is a Reader for the ReplaceCatalogTable structure.
type ReplaceCatalogTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCatalogTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCatalogTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceCatalogTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReplaceCatalogTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReplaceCatalogTableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceCatalogTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewReplaceCatalogTableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplaceCatalogTableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCatalogTableOK creates a ReplaceCatalogTableOK with default headers values
func NewReplaceCatalogTableOK() *ReplaceCatalogTableOK {
	return &ReplaceCatalogTableOK{}
}

/* ReplaceCatalogTableOK describes a response with status code 200, with default header values.

Table replaced successfully.
*/
type ReplaceCatalogTableOK struct {
	Payload *models.SuccessMessage
}

func (o *ReplaceCatalogTableOK) Error() string {
	return fmt.Sprintf("[PUT /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] replaceCatalogTableOK  %+v", 200, o.Payload)
}
func (o *ReplaceCatalogTableOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *ReplaceCatalogTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCatalogTableBadRequest creates a ReplaceCatalogTableBadRequest with default headers values
func NewReplaceCatalogTableBadRequest() *ReplaceCatalogTableBadRequest {
	return &ReplaceCatalogTableBadRequest{}
}

/* ReplaceCatalogTableBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceCatalogTableBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceCatalogTableBadRequest) Error() string {
	return fmt.Sprintf("[PUT /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] replaceCatalogTableBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceCatalogTableBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceCatalogTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCatalogTableUnauthorized creates a ReplaceCatalogTableUnauthorized with default headers values
func NewReplaceCatalogTableUnauthorized() *ReplaceCatalogTableUnauthorized {
	return &ReplaceCatalogTableUnauthorized{}
}

/* ReplaceCatalogTableUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ReplaceCatalogTableUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceCatalogTableUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] replaceCatalogTableUnauthorized  %+v", 401, o.Payload)
}
func (o *ReplaceCatalogTableUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceCatalogTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCatalogTableForbidden creates a ReplaceCatalogTableForbidden with default headers values
func NewReplaceCatalogTableForbidden() *ReplaceCatalogTableForbidden {
	return &ReplaceCatalogTableForbidden{}
}

/* ReplaceCatalogTableForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReplaceCatalogTableForbidden struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceCatalogTableForbidden) Error() string {
	return fmt.Sprintf("[PUT /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] replaceCatalogTableForbidden  %+v", 403, o.Payload)
}
func (o *ReplaceCatalogTableForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceCatalogTableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCatalogTableNotFound creates a ReplaceCatalogTableNotFound with default headers values
func NewReplaceCatalogTableNotFound() *ReplaceCatalogTableNotFound {
	return &ReplaceCatalogTableNotFound{}
}

/* ReplaceCatalogTableNotFound describes a response with status code 404, with default header values.

Not found
*/
type ReplaceCatalogTableNotFound struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceCatalogTableNotFound) Error() string {
	return fmt.Sprintf("[PUT /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] replaceCatalogTableNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceCatalogTableNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceCatalogTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCatalogTableUnprocessableEntity creates a ReplaceCatalogTableUnprocessableEntity with default headers values
func NewReplaceCatalogTableUnprocessableEntity() *ReplaceCatalogTableUnprocessableEntity {
	return &ReplaceCatalogTableUnprocessableEntity{}
}

/* ReplaceCatalogTableUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type ReplaceCatalogTableUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceCatalogTableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] replaceCatalogTableUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ReplaceCatalogTableUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceCatalogTableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCatalogTableInternalServerError creates a ReplaceCatalogTableInternalServerError with default headers values
func NewReplaceCatalogTableInternalServerError() *ReplaceCatalogTableInternalServerError {
	return &ReplaceCatalogTableInternalServerError{}
}

/* ReplaceCatalogTableInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ReplaceCatalogTableInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceCatalogTableInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] replaceCatalogTableInternalServerError  %+v", 500, o.Payload)
}
func (o *ReplaceCatalogTableInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceCatalogTableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
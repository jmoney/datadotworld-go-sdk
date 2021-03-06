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

// SuggestCatalogTableReader is a Reader for the SuggestCatalogTable structure.
type SuggestCatalogTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuggestCatalogTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuggestCatalogTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSuggestCatalogTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSuggestCatalogTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSuggestCatalogTableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuggestCatalogTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSuggestCatalogTableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSuggestCatalogTableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuggestCatalogTableOK creates a SuggestCatalogTableOK with default headers values
func NewSuggestCatalogTableOK() *SuggestCatalogTableOK {
	return &SuggestCatalogTableOK{}
}

/* SuggestCatalogTableOK describes a response with status code 200, with default header values.

Suggested changes to Table updated successfully.
*/
type SuggestCatalogTableOK struct {
	Payload *models.SuccessMessage
}

func (o *SuggestCatalogTableOK) Error() string {
	return fmt.Sprintf("[PATCH /metadata/data/sources/suggest/{owner}/{sourceid}/tables/{tableid}][%d] suggestCatalogTableOK  %+v", 200, o.Payload)
}
func (o *SuggestCatalogTableOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *SuggestCatalogTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogTableBadRequest creates a SuggestCatalogTableBadRequest with default headers values
func NewSuggestCatalogTableBadRequest() *SuggestCatalogTableBadRequest {
	return &SuggestCatalogTableBadRequest{}
}

/* SuggestCatalogTableBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SuggestCatalogTableBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogTableBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /metadata/data/sources/suggest/{owner}/{sourceid}/tables/{tableid}][%d] suggestCatalogTableBadRequest  %+v", 400, o.Payload)
}
func (o *SuggestCatalogTableBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogTableUnauthorized creates a SuggestCatalogTableUnauthorized with default headers values
func NewSuggestCatalogTableUnauthorized() *SuggestCatalogTableUnauthorized {
	return &SuggestCatalogTableUnauthorized{}
}

/* SuggestCatalogTableUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SuggestCatalogTableUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogTableUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /metadata/data/sources/suggest/{owner}/{sourceid}/tables/{tableid}][%d] suggestCatalogTableUnauthorized  %+v", 401, o.Payload)
}
func (o *SuggestCatalogTableUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogTableForbidden creates a SuggestCatalogTableForbidden with default headers values
func NewSuggestCatalogTableForbidden() *SuggestCatalogTableForbidden {
	return &SuggestCatalogTableForbidden{}
}

/* SuggestCatalogTableForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SuggestCatalogTableForbidden struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogTableForbidden) Error() string {
	return fmt.Sprintf("[PATCH /metadata/data/sources/suggest/{owner}/{sourceid}/tables/{tableid}][%d] suggestCatalogTableForbidden  %+v", 403, o.Payload)
}
func (o *SuggestCatalogTableForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogTableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogTableNotFound creates a SuggestCatalogTableNotFound with default headers values
func NewSuggestCatalogTableNotFound() *SuggestCatalogTableNotFound {
	return &SuggestCatalogTableNotFound{}
}

/* SuggestCatalogTableNotFound describes a response with status code 404, with default header values.

Not found
*/
type SuggestCatalogTableNotFound struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogTableNotFound) Error() string {
	return fmt.Sprintf("[PATCH /metadata/data/sources/suggest/{owner}/{sourceid}/tables/{tableid}][%d] suggestCatalogTableNotFound  %+v", 404, o.Payload)
}
func (o *SuggestCatalogTableNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogTableUnprocessableEntity creates a SuggestCatalogTableUnprocessableEntity with default headers values
func NewSuggestCatalogTableUnprocessableEntity() *SuggestCatalogTableUnprocessableEntity {
	return &SuggestCatalogTableUnprocessableEntity{}
}

/* SuggestCatalogTableUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type SuggestCatalogTableUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogTableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /metadata/data/sources/suggest/{owner}/{sourceid}/tables/{tableid}][%d] suggestCatalogTableUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *SuggestCatalogTableUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogTableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogTableInternalServerError creates a SuggestCatalogTableInternalServerError with default headers values
func NewSuggestCatalogTableInternalServerError() *SuggestCatalogTableInternalServerError {
	return &SuggestCatalogTableInternalServerError{}
}

/* SuggestCatalogTableInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SuggestCatalogTableInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogTableInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /metadata/data/sources/suggest/{owner}/{sourceid}/tables/{tableid}][%d] suggestCatalogTableInternalServerError  %+v", 500, o.Payload)
}
func (o *SuggestCatalogTableInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogTableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

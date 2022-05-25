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

// DeleteCatalogTableReader is a Reader for the DeleteCatalogTable structure.
type DeleteCatalogTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCatalogTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCatalogTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCatalogTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCatalogTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCatalogTableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCatalogTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCatalogTableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCatalogTableOK creates a DeleteCatalogTableOK with default headers values
func NewDeleteCatalogTableOK() *DeleteCatalogTableOK {
	return &DeleteCatalogTableOK{}
}

/* DeleteCatalogTableOK describes a response with status code 200, with default header values.

Table has been successfully deleted.
*/
type DeleteCatalogTableOK struct {
	Payload *models.SuccessMessage
}

func (o *DeleteCatalogTableOK) Error() string {
	return fmt.Sprintf("[DELETE /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] deleteCatalogTableOK  %+v", 200, o.Payload)
}
func (o *DeleteCatalogTableOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *DeleteCatalogTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCatalogTableBadRequest creates a DeleteCatalogTableBadRequest with default headers values
func NewDeleteCatalogTableBadRequest() *DeleteCatalogTableBadRequest {
	return &DeleteCatalogTableBadRequest{}
}

/* DeleteCatalogTableBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteCatalogTableBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *DeleteCatalogTableBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] deleteCatalogTableBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCatalogTableBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteCatalogTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCatalogTableUnauthorized creates a DeleteCatalogTableUnauthorized with default headers values
func NewDeleteCatalogTableUnauthorized() *DeleteCatalogTableUnauthorized {
	return &DeleteCatalogTableUnauthorized{}
}

/* DeleteCatalogTableUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteCatalogTableUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *DeleteCatalogTableUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] deleteCatalogTableUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCatalogTableUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteCatalogTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCatalogTableForbidden creates a DeleteCatalogTableForbidden with default headers values
func NewDeleteCatalogTableForbidden() *DeleteCatalogTableForbidden {
	return &DeleteCatalogTableForbidden{}
}

/* DeleteCatalogTableForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCatalogTableForbidden struct {
	Payload *models.ErrorMessage
}

func (o *DeleteCatalogTableForbidden) Error() string {
	return fmt.Sprintf("[DELETE /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] deleteCatalogTableForbidden  %+v", 403, o.Payload)
}
func (o *DeleteCatalogTableForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteCatalogTableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCatalogTableNotFound creates a DeleteCatalogTableNotFound with default headers values
func NewDeleteCatalogTableNotFound() *DeleteCatalogTableNotFound {
	return &DeleteCatalogTableNotFound{}
}

/* DeleteCatalogTableNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteCatalogTableNotFound struct {
	Payload *models.ErrorMessage
}

func (o *DeleteCatalogTableNotFound) Error() string {
	return fmt.Sprintf("[DELETE /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] deleteCatalogTableNotFound  %+v", 404, o.Payload)
}
func (o *DeleteCatalogTableNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteCatalogTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCatalogTableInternalServerError creates a DeleteCatalogTableInternalServerError with default headers values
func NewDeleteCatalogTableInternalServerError() *DeleteCatalogTableInternalServerError {
	return &DeleteCatalogTableInternalServerError{}
}

/* DeleteCatalogTableInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteCatalogTableInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *DeleteCatalogTableInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /metadata/data/sources/{owner}/{sourceid}/tables/{tableid}][%d] deleteCatalogTableInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteCatalogTableInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteCatalogTableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package relationships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// GetTableRelationshipsReader is a Reader for the GetTableRelationships structure.
type GetTableRelationshipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTableRelationshipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTableRelationshipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTableRelationshipsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTableRelationshipsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTableRelationshipsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTableRelationshipsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTableRelationshipsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTableRelationshipsOK creates a GetTableRelationshipsOK with default headers values
func NewGetTableRelationshipsOK() *GetTableRelationshipsOK {
	return &GetTableRelationshipsOK{}
}

/* GetTableRelationshipsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTableRelationshipsOK struct {
	Payload *models.PaginatedMetadataResourceResults
}

func (o *GetTableRelationshipsOK) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/table/{sourceId}/{tableId}][%d] getTableRelationshipsOK  %+v", 200, o.Payload)
}
func (o *GetTableRelationshipsOK) GetPayload() *models.PaginatedMetadataResourceResults {
	return o.Payload
}

func (o *GetTableRelationshipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedMetadataResourceResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTableRelationshipsBadRequest creates a GetTableRelationshipsBadRequest with default headers values
func NewGetTableRelationshipsBadRequest() *GetTableRelationshipsBadRequest {
	return &GetTableRelationshipsBadRequest{}
}

/* GetTableRelationshipsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetTableRelationshipsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetTableRelationshipsBadRequest) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/table/{sourceId}/{tableId}][%d] getTableRelationshipsBadRequest  %+v", 400, o.Payload)
}
func (o *GetTableRelationshipsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetTableRelationshipsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTableRelationshipsUnauthorized creates a GetTableRelationshipsUnauthorized with default headers values
func NewGetTableRelationshipsUnauthorized() *GetTableRelationshipsUnauthorized {
	return &GetTableRelationshipsUnauthorized{}
}

/* GetTableRelationshipsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTableRelationshipsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetTableRelationshipsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/table/{sourceId}/{tableId}][%d] getTableRelationshipsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetTableRelationshipsUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetTableRelationshipsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTableRelationshipsForbidden creates a GetTableRelationshipsForbidden with default headers values
func NewGetTableRelationshipsForbidden() *GetTableRelationshipsForbidden {
	return &GetTableRelationshipsForbidden{}
}

/* GetTableRelationshipsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTableRelationshipsForbidden struct {
	Payload *models.ErrorMessage
}

func (o *GetTableRelationshipsForbidden) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/table/{sourceId}/{tableId}][%d] getTableRelationshipsForbidden  %+v", 403, o.Payload)
}
func (o *GetTableRelationshipsForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetTableRelationshipsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTableRelationshipsNotFound creates a GetTableRelationshipsNotFound with default headers values
func NewGetTableRelationshipsNotFound() *GetTableRelationshipsNotFound {
	return &GetTableRelationshipsNotFound{}
}

/* GetTableRelationshipsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetTableRelationshipsNotFound struct {
	Payload *models.ErrorMessage
}

func (o *GetTableRelationshipsNotFound) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/table/{sourceId}/{tableId}][%d] getTableRelationshipsNotFound  %+v", 404, o.Payload)
}
func (o *GetTableRelationshipsNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetTableRelationshipsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTableRelationshipsInternalServerError creates a GetTableRelationshipsInternalServerError with default headers values
func NewGetTableRelationshipsInternalServerError() *GetTableRelationshipsInternalServerError {
	return &GetTableRelationshipsInternalServerError{}
}

/* GetTableRelationshipsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetTableRelationshipsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetTableRelationshipsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/table/{sourceId}/{tableId}][%d] getTableRelationshipsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTableRelationshipsInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetTableRelationshipsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

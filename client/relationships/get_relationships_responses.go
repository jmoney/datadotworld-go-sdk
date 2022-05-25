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

// GetRelationshipsReader is a Reader for the GetRelationships structure.
type GetRelationshipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRelationshipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRelationshipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRelationshipsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRelationshipsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRelationshipsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRelationshipsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRelationshipsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRelationshipsOK creates a GetRelationshipsOK with default headers values
func NewGetRelationshipsOK() *GetRelationshipsOK {
	return &GetRelationshipsOK{}
}

/* GetRelationshipsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRelationshipsOK struct {
	Payload *models.PaginatedMetadataResourceResults
}

func (o *GetRelationshipsOK) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/resource/{resourceId}][%d] getRelationshipsOK  %+v", 200, o.Payload)
}
func (o *GetRelationshipsOK) GetPayload() *models.PaginatedMetadataResourceResults {
	return o.Payload
}

func (o *GetRelationshipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedMetadataResourceResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRelationshipsBadRequest creates a GetRelationshipsBadRequest with default headers values
func NewGetRelationshipsBadRequest() *GetRelationshipsBadRequest {
	return &GetRelationshipsBadRequest{}
}

/* GetRelationshipsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetRelationshipsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetRelationshipsBadRequest) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/resource/{resourceId}][%d] getRelationshipsBadRequest  %+v", 400, o.Payload)
}
func (o *GetRelationshipsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRelationshipsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRelationshipsUnauthorized creates a GetRelationshipsUnauthorized with default headers values
func NewGetRelationshipsUnauthorized() *GetRelationshipsUnauthorized {
	return &GetRelationshipsUnauthorized{}
}

/* GetRelationshipsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRelationshipsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetRelationshipsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/resource/{resourceId}][%d] getRelationshipsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetRelationshipsUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRelationshipsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRelationshipsForbidden creates a GetRelationshipsForbidden with default headers values
func NewGetRelationshipsForbidden() *GetRelationshipsForbidden {
	return &GetRelationshipsForbidden{}
}

/* GetRelationshipsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRelationshipsForbidden struct {
	Payload *models.ErrorMessage
}

func (o *GetRelationshipsForbidden) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/resource/{resourceId}][%d] getRelationshipsForbidden  %+v", 403, o.Payload)
}
func (o *GetRelationshipsForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRelationshipsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRelationshipsNotFound creates a GetRelationshipsNotFound with default headers values
func NewGetRelationshipsNotFound() *GetRelationshipsNotFound {
	return &GetRelationshipsNotFound{}
}

/* GetRelationshipsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetRelationshipsNotFound struct {
	Payload *models.ErrorMessage
}

func (o *GetRelationshipsNotFound) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/resource/{resourceId}][%d] getRelationshipsNotFound  %+v", 404, o.Payload)
}
func (o *GetRelationshipsNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRelationshipsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRelationshipsInternalServerError creates a GetRelationshipsInternalServerError with default headers values
func NewGetRelationshipsInternalServerError() *GetRelationshipsInternalServerError {
	return &GetRelationshipsInternalServerError{}
}

/* GetRelationshipsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetRelationshipsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetRelationshipsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /metadata/relationships/{owner}/resource/{resourceId}][%d] getRelationshipsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRelationshipsInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetRelationshipsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
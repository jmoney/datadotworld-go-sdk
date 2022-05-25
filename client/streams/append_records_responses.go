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

// AppendRecordsReader is a Reader for the AppendRecords structure.
type AppendRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppendRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAppendRecordsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppendRecordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAppendRecordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAppendRecordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppendRecordsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAppendRecordsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppendRecordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAppendRecordsAccepted creates a AppendRecordsAccepted with default headers values
func NewAppendRecordsAccepted() *AppendRecordsAccepted {
	return &AppendRecordsAccepted{}
}

/* AppendRecordsAccepted describes a response with status code 202, with default header values.

Record(s) accepted.
*/
type AppendRecordsAccepted struct {
	Payload *models.SuccessMessage
}

func (o *AppendRecordsAccepted) Error() string {
	return fmt.Sprintf("[POST /streams/{owner}/{id}/{streamId}][%d] appendRecordsAccepted  %+v", 202, o.Payload)
}
func (o *AppendRecordsAccepted) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *AppendRecordsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendRecordsBadRequest creates a AppendRecordsBadRequest with default headers values
func NewAppendRecordsBadRequest() *AppendRecordsBadRequest {
	return &AppendRecordsBadRequest{}
}

/* AppendRecordsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AppendRecordsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *AppendRecordsBadRequest) Error() string {
	return fmt.Sprintf("[POST /streams/{owner}/{id}/{streamId}][%d] appendRecordsBadRequest  %+v", 400, o.Payload)
}
func (o *AppendRecordsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppendRecordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendRecordsUnauthorized creates a AppendRecordsUnauthorized with default headers values
func NewAppendRecordsUnauthorized() *AppendRecordsUnauthorized {
	return &AppendRecordsUnauthorized{}
}

/* AppendRecordsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AppendRecordsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *AppendRecordsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /streams/{owner}/{id}/{streamId}][%d] appendRecordsUnauthorized  %+v", 401, o.Payload)
}
func (o *AppendRecordsUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppendRecordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendRecordsForbidden creates a AppendRecordsForbidden with default headers values
func NewAppendRecordsForbidden() *AppendRecordsForbidden {
	return &AppendRecordsForbidden{}
}

/* AppendRecordsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AppendRecordsForbidden struct {
	Payload *models.ErrorMessage
}

func (o *AppendRecordsForbidden) Error() string {
	return fmt.Sprintf("[POST /streams/{owner}/{id}/{streamId}][%d] appendRecordsForbidden  %+v", 403, o.Payload)
}
func (o *AppendRecordsForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppendRecordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendRecordsNotFound creates a AppendRecordsNotFound with default headers values
func NewAppendRecordsNotFound() *AppendRecordsNotFound {
	return &AppendRecordsNotFound{}
}

/* AppendRecordsNotFound describes a response with status code 404, with default header values.

Not found
*/
type AppendRecordsNotFound struct {
	Payload *models.ErrorMessage
}

func (o *AppendRecordsNotFound) Error() string {
	return fmt.Sprintf("[POST /streams/{owner}/{id}/{streamId}][%d] appendRecordsNotFound  %+v", 404, o.Payload)
}
func (o *AppendRecordsNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppendRecordsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendRecordsUnprocessableEntity creates a AppendRecordsUnprocessableEntity with default headers values
func NewAppendRecordsUnprocessableEntity() *AppendRecordsUnprocessableEntity {
	return &AppendRecordsUnprocessableEntity{}
}

/* AppendRecordsUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type AppendRecordsUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *AppendRecordsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /streams/{owner}/{id}/{streamId}][%d] appendRecordsUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *AppendRecordsUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppendRecordsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendRecordsInternalServerError creates a AppendRecordsInternalServerError with default headers values
func NewAppendRecordsInternalServerError() *AppendRecordsInternalServerError {
	return &AppendRecordsInternalServerError{}
}

/* AppendRecordsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type AppendRecordsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *AppendRecordsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /streams/{owner}/{id}/{streamId}][%d] appendRecordsInternalServerError  %+v", 500, o.Payload)
}
func (o *AppendRecordsInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppendRecordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
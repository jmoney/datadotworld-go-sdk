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

// SyncReader is a Reader for the Sync structure.
type SyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSyncUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSyncTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSyncInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSyncOK creates a SyncOK with default headers values
func NewSyncOK() *SyncOK {
	return &SyncOK{}
}

/* SyncOK describes a response with status code 200, with default header values.

Sync started.
*/
type SyncOK struct {
	Payload *models.SuccessMessage
}

func (o *SyncOK) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}/{id}/sync][%d] syncOK  %+v", 200, o.Payload)
}
func (o *SyncOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *SyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncBadRequest creates a SyncBadRequest with default headers values
func NewSyncBadRequest() *SyncBadRequest {
	return &SyncBadRequest{}
}

/* SyncBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SyncBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *SyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}/{id}/sync][%d] syncBadRequest  %+v", 400, o.Payload)
}
func (o *SyncBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncUnauthorized creates a SyncUnauthorized with default headers values
func NewSyncUnauthorized() *SyncUnauthorized {
	return &SyncUnauthorized{}
}

/* SyncUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SyncUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *SyncUnauthorized) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}/{id}/sync][%d] syncUnauthorized  %+v", 401, o.Payload)
}
func (o *SyncUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncForbidden creates a SyncForbidden with default headers values
func NewSyncForbidden() *SyncForbidden {
	return &SyncForbidden{}
}

/* SyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SyncForbidden struct {
	Payload *models.ErrorMessage
}

func (o *SyncForbidden) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}/{id}/sync][%d] syncForbidden  %+v", 403, o.Payload)
}
func (o *SyncForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncNotFound creates a SyncNotFound with default headers values
func NewSyncNotFound() *SyncNotFound {
	return &SyncNotFound{}
}

/* SyncNotFound describes a response with status code 404, with default header values.

Not found
*/
type SyncNotFound struct {
	Payload *models.ErrorMessage
}

func (o *SyncNotFound) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}/{id}/sync][%d] syncNotFound  %+v", 404, o.Payload)
}
func (o *SyncNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncTooManyRequests creates a SyncTooManyRequests with default headers values
func NewSyncTooManyRequests() *SyncTooManyRequests {
	return &SyncTooManyRequests{}
}

/* SyncTooManyRequests describes a response with status code 429, with default header values.

Too many requests
*/
type SyncTooManyRequests struct {
	Payload *models.ErrorMessage
}

func (o *SyncTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}/{id}/sync][%d] syncTooManyRequests  %+v", 429, o.Payload)
}
func (o *SyncTooManyRequests) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncInternalServerError creates a SyncInternalServerError with default headers values
func NewSyncInternalServerError() *SyncInternalServerError {
	return &SyncInternalServerError{}
}

/* SyncInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SyncInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *SyncInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datasets/{owner}/{id}/sync][%d] syncInternalServerError  %+v", 500, o.Payload)
}
func (o *SyncInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

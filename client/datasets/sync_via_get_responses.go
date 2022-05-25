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

// SyncViaGetReader is a Reader for the SyncViaGet structure.
type SyncViaGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncViaGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncViaGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSyncViaGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSyncViaGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSyncViaGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyncViaGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSyncViaGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSyncViaGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSyncViaGetOK creates a SyncViaGetOK with default headers values
func NewSyncViaGetOK() *SyncViaGetOK {
	return &SyncViaGetOK{}
}

/* SyncViaGetOK describes a response with status code 200, with default header values.

Sync started.
*/
type SyncViaGetOK struct {
	Payload *models.SuccessMessage
}

func (o *SyncViaGetOK) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}/sync][%d] syncViaGetOK  %+v", 200, o.Payload)
}
func (o *SyncViaGetOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *SyncViaGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncViaGetBadRequest creates a SyncViaGetBadRequest with default headers values
func NewSyncViaGetBadRequest() *SyncViaGetBadRequest {
	return &SyncViaGetBadRequest{}
}

/* SyncViaGetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SyncViaGetBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *SyncViaGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}/sync][%d] syncViaGetBadRequest  %+v", 400, o.Payload)
}
func (o *SyncViaGetBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncViaGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncViaGetUnauthorized creates a SyncViaGetUnauthorized with default headers values
func NewSyncViaGetUnauthorized() *SyncViaGetUnauthorized {
	return &SyncViaGetUnauthorized{}
}

/* SyncViaGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SyncViaGetUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *SyncViaGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}/sync][%d] syncViaGetUnauthorized  %+v", 401, o.Payload)
}
func (o *SyncViaGetUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncViaGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncViaGetForbidden creates a SyncViaGetForbidden with default headers values
func NewSyncViaGetForbidden() *SyncViaGetForbidden {
	return &SyncViaGetForbidden{}
}

/* SyncViaGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SyncViaGetForbidden struct {
	Payload *models.ErrorMessage
}

func (o *SyncViaGetForbidden) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}/sync][%d] syncViaGetForbidden  %+v", 403, o.Payload)
}
func (o *SyncViaGetForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncViaGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncViaGetNotFound creates a SyncViaGetNotFound with default headers values
func NewSyncViaGetNotFound() *SyncViaGetNotFound {
	return &SyncViaGetNotFound{}
}

/* SyncViaGetNotFound describes a response with status code 404, with default header values.

Not found
*/
type SyncViaGetNotFound struct {
	Payload *models.ErrorMessage
}

func (o *SyncViaGetNotFound) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}/sync][%d] syncViaGetNotFound  %+v", 404, o.Payload)
}
func (o *SyncViaGetNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncViaGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncViaGetTooManyRequests creates a SyncViaGetTooManyRequests with default headers values
func NewSyncViaGetTooManyRequests() *SyncViaGetTooManyRequests {
	return &SyncViaGetTooManyRequests{}
}

/* SyncViaGetTooManyRequests describes a response with status code 429, with default header values.

Too many requests
*/
type SyncViaGetTooManyRequests struct {
	Payload *models.ErrorMessage
}

func (o *SyncViaGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}/sync][%d] syncViaGetTooManyRequests  %+v", 429, o.Payload)
}
func (o *SyncViaGetTooManyRequests) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncViaGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncViaGetInternalServerError creates a SyncViaGetInternalServerError with default headers values
func NewSyncViaGetInternalServerError() *SyncViaGetInternalServerError {
	return &SyncViaGetInternalServerError{}
}

/* SyncViaGetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SyncViaGetInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *SyncViaGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}/sync][%d] syncViaGetInternalServerError  %+v", 500, o.Payload)
}
func (o *SyncViaGetInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SyncViaGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

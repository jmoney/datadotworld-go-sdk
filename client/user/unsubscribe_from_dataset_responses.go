// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// UnsubscribeFromDatasetReader is a Reader for the UnsubscribeFromDataset structure.
type UnsubscribeFromDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnsubscribeFromDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnsubscribeFromDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnsubscribeFromDatasetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUnsubscribeFromDatasetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUnsubscribeFromDatasetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnsubscribeFromDatasetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnsubscribeFromDatasetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnsubscribeFromDatasetOK creates a UnsubscribeFromDatasetOK with default headers values
func NewUnsubscribeFromDatasetOK() *UnsubscribeFromDatasetOK {
	return &UnsubscribeFromDatasetOK{}
}

/* UnsubscribeFromDatasetOK describes a response with status code 200, with default header values.

Webhook subscription deleted successfully.
*/
type UnsubscribeFromDatasetOK struct {
	Payload *models.SuccessMessage
}

func (o *UnsubscribeFromDatasetOK) Error() string {
	return fmt.Sprintf("[DELETE /user/webhooks/datasets/{owner}/{id}][%d] unsubscribeFromDatasetOK  %+v", 200, o.Payload)
}
func (o *UnsubscribeFromDatasetOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *UnsubscribeFromDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnsubscribeFromDatasetBadRequest creates a UnsubscribeFromDatasetBadRequest with default headers values
func NewUnsubscribeFromDatasetBadRequest() *UnsubscribeFromDatasetBadRequest {
	return &UnsubscribeFromDatasetBadRequest{}
}

/* UnsubscribeFromDatasetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UnsubscribeFromDatasetBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *UnsubscribeFromDatasetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user/webhooks/datasets/{owner}/{id}][%d] unsubscribeFromDatasetBadRequest  %+v", 400, o.Payload)
}
func (o *UnsubscribeFromDatasetBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UnsubscribeFromDatasetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnsubscribeFromDatasetUnauthorized creates a UnsubscribeFromDatasetUnauthorized with default headers values
func NewUnsubscribeFromDatasetUnauthorized() *UnsubscribeFromDatasetUnauthorized {
	return &UnsubscribeFromDatasetUnauthorized{}
}

/* UnsubscribeFromDatasetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UnsubscribeFromDatasetUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *UnsubscribeFromDatasetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /user/webhooks/datasets/{owner}/{id}][%d] unsubscribeFromDatasetUnauthorized  %+v", 401, o.Payload)
}
func (o *UnsubscribeFromDatasetUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UnsubscribeFromDatasetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnsubscribeFromDatasetForbidden creates a UnsubscribeFromDatasetForbidden with default headers values
func NewUnsubscribeFromDatasetForbidden() *UnsubscribeFromDatasetForbidden {
	return &UnsubscribeFromDatasetForbidden{}
}

/* UnsubscribeFromDatasetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UnsubscribeFromDatasetForbidden struct {
	Payload *models.ErrorMessage
}

func (o *UnsubscribeFromDatasetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /user/webhooks/datasets/{owner}/{id}][%d] unsubscribeFromDatasetForbidden  %+v", 403, o.Payload)
}
func (o *UnsubscribeFromDatasetForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UnsubscribeFromDatasetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnsubscribeFromDatasetNotFound creates a UnsubscribeFromDatasetNotFound with default headers values
func NewUnsubscribeFromDatasetNotFound() *UnsubscribeFromDatasetNotFound {
	return &UnsubscribeFromDatasetNotFound{}
}

/* UnsubscribeFromDatasetNotFound describes a response with status code 404, with default header values.

Not found
*/
type UnsubscribeFromDatasetNotFound struct {
	Payload *models.ErrorMessage
}

func (o *UnsubscribeFromDatasetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/webhooks/datasets/{owner}/{id}][%d] unsubscribeFromDatasetNotFound  %+v", 404, o.Payload)
}
func (o *UnsubscribeFromDatasetNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UnsubscribeFromDatasetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnsubscribeFromDatasetInternalServerError creates a UnsubscribeFromDatasetInternalServerError with default headers values
func NewUnsubscribeFromDatasetInternalServerError() *UnsubscribeFromDatasetInternalServerError {
	return &UnsubscribeFromDatasetInternalServerError{}
}

/* UnsubscribeFromDatasetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UnsubscribeFromDatasetInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *UnsubscribeFromDatasetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /user/webhooks/datasets/{owner}/{id}][%d] unsubscribeFromDatasetInternalServerError  %+v", 500, o.Payload)
}
func (o *UnsubscribeFromDatasetInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UnsubscribeFromDatasetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
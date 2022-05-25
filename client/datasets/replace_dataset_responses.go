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

// ReplaceDatasetReader is a Reader for the ReplaceDataset structure.
type ReplaceDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceDatasetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReplaceDatasetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReplaceDatasetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceDatasetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewReplaceDatasetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplaceDatasetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceDatasetOK creates a ReplaceDatasetOK with default headers values
func NewReplaceDatasetOK() *ReplaceDatasetOK {
	return &ReplaceDatasetOK{}
}

/* ReplaceDatasetOK describes a response with status code 200, with default header values.

Dataset replaced successfully.
*/
type ReplaceDatasetOK struct {
	Payload *models.SuccessMessage
}

func (o *ReplaceDatasetOK) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}][%d] replaceDatasetOK  %+v", 200, o.Payload)
}
func (o *ReplaceDatasetOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *ReplaceDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDatasetBadRequest creates a ReplaceDatasetBadRequest with default headers values
func NewReplaceDatasetBadRequest() *ReplaceDatasetBadRequest {
	return &ReplaceDatasetBadRequest{}
}

/* ReplaceDatasetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceDatasetBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceDatasetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}][%d] replaceDatasetBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceDatasetBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceDatasetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDatasetUnauthorized creates a ReplaceDatasetUnauthorized with default headers values
func NewReplaceDatasetUnauthorized() *ReplaceDatasetUnauthorized {
	return &ReplaceDatasetUnauthorized{}
}

/* ReplaceDatasetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ReplaceDatasetUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceDatasetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}][%d] replaceDatasetUnauthorized  %+v", 401, o.Payload)
}
func (o *ReplaceDatasetUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceDatasetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDatasetForbidden creates a ReplaceDatasetForbidden with default headers values
func NewReplaceDatasetForbidden() *ReplaceDatasetForbidden {
	return &ReplaceDatasetForbidden{}
}

/* ReplaceDatasetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReplaceDatasetForbidden struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceDatasetForbidden) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}][%d] replaceDatasetForbidden  %+v", 403, o.Payload)
}
func (o *ReplaceDatasetForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceDatasetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDatasetNotFound creates a ReplaceDatasetNotFound with default headers values
func NewReplaceDatasetNotFound() *ReplaceDatasetNotFound {
	return &ReplaceDatasetNotFound{}
}

/* ReplaceDatasetNotFound describes a response with status code 404, with default header values.

Not found
*/
type ReplaceDatasetNotFound struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceDatasetNotFound) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}][%d] replaceDatasetNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceDatasetNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceDatasetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDatasetUnprocessableEntity creates a ReplaceDatasetUnprocessableEntity with default headers values
func NewReplaceDatasetUnprocessableEntity() *ReplaceDatasetUnprocessableEntity {
	return &ReplaceDatasetUnprocessableEntity{}
}

/* ReplaceDatasetUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type ReplaceDatasetUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceDatasetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}][%d] replaceDatasetUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ReplaceDatasetUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceDatasetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDatasetInternalServerError creates a ReplaceDatasetInternalServerError with default headers values
func NewReplaceDatasetInternalServerError() *ReplaceDatasetInternalServerError {
	return &ReplaceDatasetInternalServerError{}
}

/* ReplaceDatasetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ReplaceDatasetInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *ReplaceDatasetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /datasets/{owner}/{id}][%d] replaceDatasetInternalServerError  %+v", 500, o.Payload)
}
func (o *ReplaceDatasetInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReplaceDatasetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
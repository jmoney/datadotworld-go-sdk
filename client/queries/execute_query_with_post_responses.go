// Code generated by go-swagger; DO NOT EDIT.

package queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// ExecuteQueryWithPostReader is a Reader for the ExecuteQueryWithPost structure.
type ExecuteQueryWithPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteQueryWithPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteQueryWithPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExecuteQueryWithPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExecuteQueryWithPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExecuteQueryWithPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExecuteQueryWithPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewExecuteQueryWithPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecuteQueryWithPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecuteQueryWithPostOK creates a ExecuteQueryWithPostOK with default headers values
func NewExecuteQueryWithPostOK() *ExecuteQueryWithPostOK {
	return &ExecuteQueryWithPostOK{}
}

/* ExecuteQueryWithPostOK describes a response with status code 200, with default header values.

Successfully executed a query.
*/
type ExecuteQueryWithPostOK struct {
	Payload *models.SuccessMessage
}

func (o *ExecuteQueryWithPostOK) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] executeQueryWithPostOK  %+v", 200, o.Payload)
}
func (o *ExecuteQueryWithPostOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *ExecuteQueryWithPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteQueryWithPostBadRequest creates a ExecuteQueryWithPostBadRequest with default headers values
func NewExecuteQueryWithPostBadRequest() *ExecuteQueryWithPostBadRequest {
	return &ExecuteQueryWithPostBadRequest{}
}

/* ExecuteQueryWithPostBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ExecuteQueryWithPostBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *ExecuteQueryWithPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] executeQueryWithPostBadRequest  %+v", 400, o.Payload)
}
func (o *ExecuteQueryWithPostBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ExecuteQueryWithPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteQueryWithPostUnauthorized creates a ExecuteQueryWithPostUnauthorized with default headers values
func NewExecuteQueryWithPostUnauthorized() *ExecuteQueryWithPostUnauthorized {
	return &ExecuteQueryWithPostUnauthorized{}
}

/* ExecuteQueryWithPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ExecuteQueryWithPostUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *ExecuteQueryWithPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] executeQueryWithPostUnauthorized  %+v", 401, o.Payload)
}
func (o *ExecuteQueryWithPostUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ExecuteQueryWithPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteQueryWithPostForbidden creates a ExecuteQueryWithPostForbidden with default headers values
func NewExecuteQueryWithPostForbidden() *ExecuteQueryWithPostForbidden {
	return &ExecuteQueryWithPostForbidden{}
}

/* ExecuteQueryWithPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExecuteQueryWithPostForbidden struct {
	Payload *models.ErrorMessage
}

func (o *ExecuteQueryWithPostForbidden) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] executeQueryWithPostForbidden  %+v", 403, o.Payload)
}
func (o *ExecuteQueryWithPostForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ExecuteQueryWithPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteQueryWithPostNotFound creates a ExecuteQueryWithPostNotFound with default headers values
func NewExecuteQueryWithPostNotFound() *ExecuteQueryWithPostNotFound {
	return &ExecuteQueryWithPostNotFound{}
}

/* ExecuteQueryWithPostNotFound describes a response with status code 404, with default header values.

Not found
*/
type ExecuteQueryWithPostNotFound struct {
	Payload *models.ErrorMessage
}

func (o *ExecuteQueryWithPostNotFound) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] executeQueryWithPostNotFound  %+v", 404, o.Payload)
}
func (o *ExecuteQueryWithPostNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ExecuteQueryWithPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteQueryWithPostUnprocessableEntity creates a ExecuteQueryWithPostUnprocessableEntity with default headers values
func NewExecuteQueryWithPostUnprocessableEntity() *ExecuteQueryWithPostUnprocessableEntity {
	return &ExecuteQueryWithPostUnprocessableEntity{}
}

/* ExecuteQueryWithPostUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type ExecuteQueryWithPostUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *ExecuteQueryWithPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] executeQueryWithPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ExecuteQueryWithPostUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ExecuteQueryWithPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteQueryWithPostInternalServerError creates a ExecuteQueryWithPostInternalServerError with default headers values
func NewExecuteQueryWithPostInternalServerError() *ExecuteQueryWithPostInternalServerError {
	return &ExecuteQueryWithPostInternalServerError{}
}

/* ExecuteQueryWithPostInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ExecuteQueryWithPostInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *ExecuteQueryWithPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /queries/{id}/results][%d] executeQueryWithPostInternalServerError  %+v", 500, o.Payload)
}
func (o *ExecuteQueryWithPostInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ExecuteQueryWithPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

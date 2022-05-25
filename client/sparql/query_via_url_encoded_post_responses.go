// Code generated by go-swagger; DO NOT EDIT.

package sparql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// QueryViaURLEncodedPostReader is a Reader for the QueryViaURLEncodedPost structure.
type QueryViaURLEncodedPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryViaURLEncodedPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryViaURLEncodedPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryViaURLEncodedPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryViaURLEncodedPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryViaURLEncodedPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryViaURLEncodedPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewQueryViaURLEncodedPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryViaURLEncodedPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryViaURLEncodedPostOK creates a QueryViaURLEncodedPostOK with default headers values
func NewQueryViaURLEncodedPostOK() *QueryViaURLEncodedPostOK {
	return &QueryViaURLEncodedPostOK{}
}

/* QueryViaURLEncodedPostOK describes a response with status code 200, with default header values.

The request has succeeded.
*/
type QueryViaURLEncodedPostOK struct {
}

func (o *QueryViaURLEncodedPostOK) Error() string {
	return fmt.Sprintf("[POST /sparql/{owner}/{id}][%d] queryViaUrlEncodedPostOK ", 200)
}

func (o *QueryViaURLEncodedPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryViaURLEncodedPostBadRequest creates a QueryViaURLEncodedPostBadRequest with default headers values
func NewQueryViaURLEncodedPostBadRequest() *QueryViaURLEncodedPostBadRequest {
	return &QueryViaURLEncodedPostBadRequest{}
}

/* QueryViaURLEncodedPostBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type QueryViaURLEncodedPostBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *QueryViaURLEncodedPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /sparql/{owner}/{id}][%d] queryViaUrlEncodedPostBadRequest  %+v", 400, o.Payload)
}
func (o *QueryViaURLEncodedPostBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryViaURLEncodedPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryViaURLEncodedPostUnauthorized creates a QueryViaURLEncodedPostUnauthorized with default headers values
func NewQueryViaURLEncodedPostUnauthorized() *QueryViaURLEncodedPostUnauthorized {
	return &QueryViaURLEncodedPostUnauthorized{}
}

/* QueryViaURLEncodedPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type QueryViaURLEncodedPostUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *QueryViaURLEncodedPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sparql/{owner}/{id}][%d] queryViaUrlEncodedPostUnauthorized  %+v", 401, o.Payload)
}
func (o *QueryViaURLEncodedPostUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryViaURLEncodedPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryViaURLEncodedPostForbidden creates a QueryViaURLEncodedPostForbidden with default headers values
func NewQueryViaURLEncodedPostForbidden() *QueryViaURLEncodedPostForbidden {
	return &QueryViaURLEncodedPostForbidden{}
}

/* QueryViaURLEncodedPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryViaURLEncodedPostForbidden struct {
	Payload *models.ErrorMessage
}

func (o *QueryViaURLEncodedPostForbidden) Error() string {
	return fmt.Sprintf("[POST /sparql/{owner}/{id}][%d] queryViaUrlEncodedPostForbidden  %+v", 403, o.Payload)
}
func (o *QueryViaURLEncodedPostForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryViaURLEncodedPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryViaURLEncodedPostNotFound creates a QueryViaURLEncodedPostNotFound with default headers values
func NewQueryViaURLEncodedPostNotFound() *QueryViaURLEncodedPostNotFound {
	return &QueryViaURLEncodedPostNotFound{}
}

/* QueryViaURLEncodedPostNotFound describes a response with status code 404, with default header values.

Not found
*/
type QueryViaURLEncodedPostNotFound struct {
	Payload *models.ErrorMessage
}

func (o *QueryViaURLEncodedPostNotFound) Error() string {
	return fmt.Sprintf("[POST /sparql/{owner}/{id}][%d] queryViaUrlEncodedPostNotFound  %+v", 404, o.Payload)
}
func (o *QueryViaURLEncodedPostNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryViaURLEncodedPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryViaURLEncodedPostUnprocessableEntity creates a QueryViaURLEncodedPostUnprocessableEntity with default headers values
func NewQueryViaURLEncodedPostUnprocessableEntity() *QueryViaURLEncodedPostUnprocessableEntity {
	return &QueryViaURLEncodedPostUnprocessableEntity{}
}

/* QueryViaURLEncodedPostUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type QueryViaURLEncodedPostUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *QueryViaURLEncodedPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /sparql/{owner}/{id}][%d] queryViaUrlEncodedPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *QueryViaURLEncodedPostUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryViaURLEncodedPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryViaURLEncodedPostInternalServerError creates a QueryViaURLEncodedPostInternalServerError with default headers values
func NewQueryViaURLEncodedPostInternalServerError() *QueryViaURLEncodedPostInternalServerError {
	return &QueryViaURLEncodedPostInternalServerError{}
}

/* QueryViaURLEncodedPostInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type QueryViaURLEncodedPostInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *QueryViaURLEncodedPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sparql/{owner}/{id}][%d] queryViaUrlEncodedPostInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryViaURLEncodedPostInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryViaURLEncodedPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// CreateProjectSavedQueryReader is a Reader for the CreateProjectSavedQuery structure.
type CreateProjectSavedQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectSavedQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProjectSavedQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProjectSavedQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateProjectSavedQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProjectSavedQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProjectSavedQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProjectSavedQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectSavedQueryOK creates a CreateProjectSavedQueryOK with default headers values
func NewCreateProjectSavedQueryOK() *CreateProjectSavedQueryOK {
	return &CreateProjectSavedQueryOK{}
}

/* CreateProjectSavedQueryOK describes a response with status code 200, with default header values.

Successfully created saved query.
*/
type CreateProjectSavedQueryOK struct {
	Payload *models.QuerySummaryResponse
}

func (o *CreateProjectSavedQueryOK) Error() string {
	return fmt.Sprintf("[POST /projects/{owner}/{id}/queries][%d] createProjectSavedQueryOK  %+v", 200, o.Payload)
}
func (o *CreateProjectSavedQueryOK) GetPayload() *models.QuerySummaryResponse {
	return o.Payload
}

func (o *CreateProjectSavedQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuerySummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectSavedQueryBadRequest creates a CreateProjectSavedQueryBadRequest with default headers values
func NewCreateProjectSavedQueryBadRequest() *CreateProjectSavedQueryBadRequest {
	return &CreateProjectSavedQueryBadRequest{}
}

/* CreateProjectSavedQueryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateProjectSavedQueryBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *CreateProjectSavedQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{owner}/{id}/queries][%d] createProjectSavedQueryBadRequest  %+v", 400, o.Payload)
}
func (o *CreateProjectSavedQueryBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateProjectSavedQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectSavedQueryUnauthorized creates a CreateProjectSavedQueryUnauthorized with default headers values
func NewCreateProjectSavedQueryUnauthorized() *CreateProjectSavedQueryUnauthorized {
	return &CreateProjectSavedQueryUnauthorized{}
}

/* CreateProjectSavedQueryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateProjectSavedQueryUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *CreateProjectSavedQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{owner}/{id}/queries][%d] createProjectSavedQueryUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateProjectSavedQueryUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateProjectSavedQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectSavedQueryForbidden creates a CreateProjectSavedQueryForbidden with default headers values
func NewCreateProjectSavedQueryForbidden() *CreateProjectSavedQueryForbidden {
	return &CreateProjectSavedQueryForbidden{}
}

/* CreateProjectSavedQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateProjectSavedQueryForbidden struct {
	Payload *models.ErrorMessage
}

func (o *CreateProjectSavedQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{owner}/{id}/queries][%d] createProjectSavedQueryForbidden  %+v", 403, o.Payload)
}
func (o *CreateProjectSavedQueryForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateProjectSavedQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectSavedQueryNotFound creates a CreateProjectSavedQueryNotFound with default headers values
func NewCreateProjectSavedQueryNotFound() *CreateProjectSavedQueryNotFound {
	return &CreateProjectSavedQueryNotFound{}
}

/* CreateProjectSavedQueryNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateProjectSavedQueryNotFound struct {
	Payload *models.ErrorMessage
}

func (o *CreateProjectSavedQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{owner}/{id}/queries][%d] createProjectSavedQueryNotFound  %+v", 404, o.Payload)
}
func (o *CreateProjectSavedQueryNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateProjectSavedQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectSavedQueryInternalServerError creates a CreateProjectSavedQueryInternalServerError with default headers values
func NewCreateProjectSavedQueryInternalServerError() *CreateProjectSavedQueryInternalServerError {
	return &CreateProjectSavedQueryInternalServerError{}
}

/* CreateProjectSavedQueryInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateProjectSavedQueryInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *CreateProjectSavedQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{owner}/{id}/queries][%d] createProjectSavedQueryInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateProjectSavedQueryInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CreateProjectSavedQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

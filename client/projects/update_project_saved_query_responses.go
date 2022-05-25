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

// UpdateProjectSavedQueryReader is a Reader for the UpdateProjectSavedQuery structure.
type UpdateProjectSavedQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectSavedQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectSavedQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectSavedQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProjectSavedQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectSavedQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectSavedQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectSavedQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectSavedQueryOK creates a UpdateProjectSavedQueryOK with default headers values
func NewUpdateProjectSavedQueryOK() *UpdateProjectSavedQueryOK {
	return &UpdateProjectSavedQueryOK{}
}

/* UpdateProjectSavedQueryOK describes a response with status code 200, with default header values.

Successfully updated saved query.
*/
type UpdateProjectSavedQueryOK struct {
	Payload *models.QuerySummaryResponse
}

func (o *UpdateProjectSavedQueryOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{owner}/{id}/queries/{queryId}][%d] updateProjectSavedQueryOK  %+v", 200, o.Payload)
}
func (o *UpdateProjectSavedQueryOK) GetPayload() *models.QuerySummaryResponse {
	return o.Payload
}

func (o *UpdateProjectSavedQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuerySummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSavedQueryBadRequest creates a UpdateProjectSavedQueryBadRequest with default headers values
func NewUpdateProjectSavedQueryBadRequest() *UpdateProjectSavedQueryBadRequest {
	return &UpdateProjectSavedQueryBadRequest{}
}

/* UpdateProjectSavedQueryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateProjectSavedQueryBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *UpdateProjectSavedQueryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{owner}/{id}/queries/{queryId}][%d] updateProjectSavedQueryBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateProjectSavedQueryBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UpdateProjectSavedQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSavedQueryUnauthorized creates a UpdateProjectSavedQueryUnauthorized with default headers values
func NewUpdateProjectSavedQueryUnauthorized() *UpdateProjectSavedQueryUnauthorized {
	return &UpdateProjectSavedQueryUnauthorized{}
}

/* UpdateProjectSavedQueryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateProjectSavedQueryUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *UpdateProjectSavedQueryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{owner}/{id}/queries/{queryId}][%d] updateProjectSavedQueryUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateProjectSavedQueryUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UpdateProjectSavedQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSavedQueryForbidden creates a UpdateProjectSavedQueryForbidden with default headers values
func NewUpdateProjectSavedQueryForbidden() *UpdateProjectSavedQueryForbidden {
	return &UpdateProjectSavedQueryForbidden{}
}

/* UpdateProjectSavedQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateProjectSavedQueryForbidden struct {
	Payload *models.ErrorMessage
}

func (o *UpdateProjectSavedQueryForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{owner}/{id}/queries/{queryId}][%d] updateProjectSavedQueryForbidden  %+v", 403, o.Payload)
}
func (o *UpdateProjectSavedQueryForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UpdateProjectSavedQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSavedQueryNotFound creates a UpdateProjectSavedQueryNotFound with default headers values
func NewUpdateProjectSavedQueryNotFound() *UpdateProjectSavedQueryNotFound {
	return &UpdateProjectSavedQueryNotFound{}
}

/* UpdateProjectSavedQueryNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateProjectSavedQueryNotFound struct {
	Payload *models.ErrorMessage
}

func (o *UpdateProjectSavedQueryNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{owner}/{id}/queries/{queryId}][%d] updateProjectSavedQueryNotFound  %+v", 404, o.Payload)
}
func (o *UpdateProjectSavedQueryNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UpdateProjectSavedQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSavedQueryInternalServerError creates a UpdateProjectSavedQueryInternalServerError with default headers values
func NewUpdateProjectSavedQueryInternalServerError() *UpdateProjectSavedQueryInternalServerError {
	return &UpdateProjectSavedQueryInternalServerError{}
}

/* UpdateProjectSavedQueryInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateProjectSavedQueryInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *UpdateProjectSavedQueryInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{owner}/{id}/queries/{queryId}][%d] updateProjectSavedQueryInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateProjectSavedQueryInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UpdateProjectSavedQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

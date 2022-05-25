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

// DeleteDatasetSavedQueryReader is a Reader for the DeleteDatasetSavedQuery structure.
type DeleteDatasetSavedQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDatasetSavedQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDatasetSavedQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDatasetSavedQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDatasetSavedQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDatasetSavedQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDatasetSavedQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDatasetSavedQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDatasetSavedQueryOK creates a DeleteDatasetSavedQueryOK with default headers values
func NewDeleteDatasetSavedQueryOK() *DeleteDatasetSavedQueryOK {
	return &DeleteDatasetSavedQueryOK{}
}

/* DeleteDatasetSavedQueryOK describes a response with status code 200, with default header values.

Successfully deleted saved query.
*/
type DeleteDatasetSavedQueryOK struct {
	Payload *models.SuccessMessage
}

func (o *DeleteDatasetSavedQueryOK) Error() string {
	return fmt.Sprintf("[DELETE /datasets/{owner}/{id}/queries/{queryId}][%d] deleteDatasetSavedQueryOK  %+v", 200, o.Payload)
}
func (o *DeleteDatasetSavedQueryOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *DeleteDatasetSavedQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatasetSavedQueryBadRequest creates a DeleteDatasetSavedQueryBadRequest with default headers values
func NewDeleteDatasetSavedQueryBadRequest() *DeleteDatasetSavedQueryBadRequest {
	return &DeleteDatasetSavedQueryBadRequest{}
}

/* DeleteDatasetSavedQueryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteDatasetSavedQueryBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *DeleteDatasetSavedQueryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /datasets/{owner}/{id}/queries/{queryId}][%d] deleteDatasetSavedQueryBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteDatasetSavedQueryBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteDatasetSavedQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatasetSavedQueryUnauthorized creates a DeleteDatasetSavedQueryUnauthorized with default headers values
func NewDeleteDatasetSavedQueryUnauthorized() *DeleteDatasetSavedQueryUnauthorized {
	return &DeleteDatasetSavedQueryUnauthorized{}
}

/* DeleteDatasetSavedQueryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteDatasetSavedQueryUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *DeleteDatasetSavedQueryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /datasets/{owner}/{id}/queries/{queryId}][%d] deleteDatasetSavedQueryUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteDatasetSavedQueryUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteDatasetSavedQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatasetSavedQueryForbidden creates a DeleteDatasetSavedQueryForbidden with default headers values
func NewDeleteDatasetSavedQueryForbidden() *DeleteDatasetSavedQueryForbidden {
	return &DeleteDatasetSavedQueryForbidden{}
}

/* DeleteDatasetSavedQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDatasetSavedQueryForbidden struct {
	Payload *models.ErrorMessage
}

func (o *DeleteDatasetSavedQueryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /datasets/{owner}/{id}/queries/{queryId}][%d] deleteDatasetSavedQueryForbidden  %+v", 403, o.Payload)
}
func (o *DeleteDatasetSavedQueryForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteDatasetSavedQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatasetSavedQueryNotFound creates a DeleteDatasetSavedQueryNotFound with default headers values
func NewDeleteDatasetSavedQueryNotFound() *DeleteDatasetSavedQueryNotFound {
	return &DeleteDatasetSavedQueryNotFound{}
}

/* DeleteDatasetSavedQueryNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteDatasetSavedQueryNotFound struct {
	Payload *models.ErrorMessage
}

func (o *DeleteDatasetSavedQueryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /datasets/{owner}/{id}/queries/{queryId}][%d] deleteDatasetSavedQueryNotFound  %+v", 404, o.Payload)
}
func (o *DeleteDatasetSavedQueryNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteDatasetSavedQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDatasetSavedQueryInternalServerError creates a DeleteDatasetSavedQueryInternalServerError with default headers values
func NewDeleteDatasetSavedQueryInternalServerError() *DeleteDatasetSavedQueryInternalServerError {
	return &DeleteDatasetSavedQueryInternalServerError{}
}

/* DeleteDatasetSavedQueryInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteDatasetSavedQueryInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *DeleteDatasetSavedQueryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /datasets/{owner}/{id}/queries/{queryId}][%d] deleteDatasetSavedQueryInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteDatasetSavedQueryInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DeleteDatasetSavedQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
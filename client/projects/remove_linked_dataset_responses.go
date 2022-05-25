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

// RemoveLinkedDatasetReader is a Reader for the RemoveLinkedDataset structure.
type RemoveLinkedDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveLinkedDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveLinkedDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveLinkedDatasetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveLinkedDatasetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveLinkedDatasetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveLinkedDatasetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveLinkedDatasetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveLinkedDatasetOK creates a RemoveLinkedDatasetOK with default headers values
func NewRemoveLinkedDatasetOK() *RemoveLinkedDatasetOK {
	return &RemoveLinkedDatasetOK{}
}

/* RemoveLinkedDatasetOK describes a response with status code 200, with default header values.

Dataset linked successfully.
*/
type RemoveLinkedDatasetOK struct {
	Payload *models.SuccessMessage
}

func (o *RemoveLinkedDatasetOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}/linkedDatasets/{linkedDatasetOwner}/{linkedDatasetId}][%d] removeLinkedDatasetOK  %+v", 200, o.Payload)
}
func (o *RemoveLinkedDatasetOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *RemoveLinkedDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLinkedDatasetBadRequest creates a RemoveLinkedDatasetBadRequest with default headers values
func NewRemoveLinkedDatasetBadRequest() *RemoveLinkedDatasetBadRequest {
	return &RemoveLinkedDatasetBadRequest{}
}

/* RemoveLinkedDatasetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RemoveLinkedDatasetBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *RemoveLinkedDatasetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}/linkedDatasets/{linkedDatasetOwner}/{linkedDatasetId}][%d] removeLinkedDatasetBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveLinkedDatasetBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *RemoveLinkedDatasetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLinkedDatasetUnauthorized creates a RemoveLinkedDatasetUnauthorized with default headers values
func NewRemoveLinkedDatasetUnauthorized() *RemoveLinkedDatasetUnauthorized {
	return &RemoveLinkedDatasetUnauthorized{}
}

/* RemoveLinkedDatasetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RemoveLinkedDatasetUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *RemoveLinkedDatasetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}/linkedDatasets/{linkedDatasetOwner}/{linkedDatasetId}][%d] removeLinkedDatasetUnauthorized  %+v", 401, o.Payload)
}
func (o *RemoveLinkedDatasetUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *RemoveLinkedDatasetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLinkedDatasetForbidden creates a RemoveLinkedDatasetForbidden with default headers values
func NewRemoveLinkedDatasetForbidden() *RemoveLinkedDatasetForbidden {
	return &RemoveLinkedDatasetForbidden{}
}

/* RemoveLinkedDatasetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RemoveLinkedDatasetForbidden struct {
	Payload *models.ErrorMessage
}

func (o *RemoveLinkedDatasetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}/linkedDatasets/{linkedDatasetOwner}/{linkedDatasetId}][%d] removeLinkedDatasetForbidden  %+v", 403, o.Payload)
}
func (o *RemoveLinkedDatasetForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *RemoveLinkedDatasetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLinkedDatasetNotFound creates a RemoveLinkedDatasetNotFound with default headers values
func NewRemoveLinkedDatasetNotFound() *RemoveLinkedDatasetNotFound {
	return &RemoveLinkedDatasetNotFound{}
}

/* RemoveLinkedDatasetNotFound describes a response with status code 404, with default header values.

Not found
*/
type RemoveLinkedDatasetNotFound struct {
	Payload *models.ErrorMessage
}

func (o *RemoveLinkedDatasetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}/linkedDatasets/{linkedDatasetOwner}/{linkedDatasetId}][%d] removeLinkedDatasetNotFound  %+v", 404, o.Payload)
}
func (o *RemoveLinkedDatasetNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *RemoveLinkedDatasetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveLinkedDatasetInternalServerError creates a RemoveLinkedDatasetInternalServerError with default headers values
func NewRemoveLinkedDatasetInternalServerError() *RemoveLinkedDatasetInternalServerError {
	return &RemoveLinkedDatasetInternalServerError{}
}

/* RemoveLinkedDatasetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RemoveLinkedDatasetInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *RemoveLinkedDatasetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{owner}/{id}/linkedDatasets/{linkedDatasetOwner}/{linkedDatasetId}][%d] removeLinkedDatasetInternalServerError  %+v", 500, o.Payload)
}
func (o *RemoveLinkedDatasetInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *RemoveLinkedDatasetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetDatasetReader is a Reader for the GetDataset structure.
type GetDatasetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatasetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatasetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDatasetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDatasetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDatasetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDatasetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDatasetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatasetOK creates a GetDatasetOK with default headers values
func NewGetDatasetOK() *GetDatasetOK {
	return &GetDatasetOK{}
}

/* GetDatasetOK describes a response with status code 200, with default header values.

Successfully retrieved dataset
*/
type GetDatasetOK struct {
	Payload *models.DatasetSummaryResponse
}

func (o *GetDatasetOK) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}][%d] getDatasetOK  %+v", 200, o.Payload)
}
func (o *GetDatasetOK) GetPayload() *models.DatasetSummaryResponse {
	return o.Payload
}

func (o *GetDatasetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatasetSummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasetBadRequest creates a GetDatasetBadRequest with default headers values
func NewGetDatasetBadRequest() *GetDatasetBadRequest {
	return &GetDatasetBadRequest{}
}

/* GetDatasetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetDatasetBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetDatasetBadRequest) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}][%d] getDatasetBadRequest  %+v", 400, o.Payload)
}
func (o *GetDatasetBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetDatasetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasetUnauthorized creates a GetDatasetUnauthorized with default headers values
func NewGetDatasetUnauthorized() *GetDatasetUnauthorized {
	return &GetDatasetUnauthorized{}
}

/* GetDatasetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDatasetUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetDatasetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}][%d] getDatasetUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDatasetUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetDatasetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasetForbidden creates a GetDatasetForbidden with default headers values
func NewGetDatasetForbidden() *GetDatasetForbidden {
	return &GetDatasetForbidden{}
}

/* GetDatasetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDatasetForbidden struct {
	Payload *models.ErrorMessage
}

func (o *GetDatasetForbidden) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}][%d] getDatasetForbidden  %+v", 403, o.Payload)
}
func (o *GetDatasetForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetDatasetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasetNotFound creates a GetDatasetNotFound with default headers values
func NewGetDatasetNotFound() *GetDatasetNotFound {
	return &GetDatasetNotFound{}
}

/* GetDatasetNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetDatasetNotFound struct {
	Payload *models.ErrorMessage
}

func (o *GetDatasetNotFound) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}][%d] getDatasetNotFound  %+v", 404, o.Payload)
}
func (o *GetDatasetNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetDatasetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatasetInternalServerError creates a GetDatasetInternalServerError with default headers values
func NewGetDatasetInternalServerError() *GetDatasetInternalServerError {
	return &GetDatasetInternalServerError{}
}

/* GetDatasetInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetDatasetInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetDatasetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasets/{owner}/{id}][%d] getDatasetInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDatasetInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetDatasetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
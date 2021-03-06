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

// FetchDatasetsReader is a Reader for the FetchDatasets structure.
type FetchDatasetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchDatasetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFetchDatasetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFetchDatasetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFetchDatasetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFetchDatasetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFetchDatasetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFetchDatasetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFetchDatasetsOK creates a FetchDatasetsOK with default headers values
func NewFetchDatasetsOK() *FetchDatasetsOK {
	return &FetchDatasetsOK{}
}

/* FetchDatasetsOK describes a response with status code 200, with default header values.

successful operation
*/
type FetchDatasetsOK struct {
	Payload *models.PaginatedDatasetResults
}

func (o *FetchDatasetsOK) Error() string {
	return fmt.Sprintf("[GET /user/datasets/own][%d] fetchDatasetsOK  %+v", 200, o.Payload)
}
func (o *FetchDatasetsOK) GetPayload() *models.PaginatedDatasetResults {
	return o.Payload
}

func (o *FetchDatasetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedDatasetResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchDatasetsBadRequest creates a FetchDatasetsBadRequest with default headers values
func NewFetchDatasetsBadRequest() *FetchDatasetsBadRequest {
	return &FetchDatasetsBadRequest{}
}

/* FetchDatasetsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type FetchDatasetsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *FetchDatasetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /user/datasets/own][%d] fetchDatasetsBadRequest  %+v", 400, o.Payload)
}
func (o *FetchDatasetsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchDatasetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchDatasetsUnauthorized creates a FetchDatasetsUnauthorized with default headers values
func NewFetchDatasetsUnauthorized() *FetchDatasetsUnauthorized {
	return &FetchDatasetsUnauthorized{}
}

/* FetchDatasetsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FetchDatasetsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *FetchDatasetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/datasets/own][%d] fetchDatasetsUnauthorized  %+v", 401, o.Payload)
}
func (o *FetchDatasetsUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchDatasetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchDatasetsForbidden creates a FetchDatasetsForbidden with default headers values
func NewFetchDatasetsForbidden() *FetchDatasetsForbidden {
	return &FetchDatasetsForbidden{}
}

/* FetchDatasetsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FetchDatasetsForbidden struct {
	Payload *models.ErrorMessage
}

func (o *FetchDatasetsForbidden) Error() string {
	return fmt.Sprintf("[GET /user/datasets/own][%d] fetchDatasetsForbidden  %+v", 403, o.Payload)
}
func (o *FetchDatasetsForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchDatasetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchDatasetsNotFound creates a FetchDatasetsNotFound with default headers values
func NewFetchDatasetsNotFound() *FetchDatasetsNotFound {
	return &FetchDatasetsNotFound{}
}

/* FetchDatasetsNotFound describes a response with status code 404, with default header values.

Not found
*/
type FetchDatasetsNotFound struct {
	Payload *models.ErrorMessage
}

func (o *FetchDatasetsNotFound) Error() string {
	return fmt.Sprintf("[GET /user/datasets/own][%d] fetchDatasetsNotFound  %+v", 404, o.Payload)
}
func (o *FetchDatasetsNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchDatasetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchDatasetsInternalServerError creates a FetchDatasetsInternalServerError with default headers values
func NewFetchDatasetsInternalServerError() *FetchDatasetsInternalServerError {
	return &FetchDatasetsInternalServerError{}
}

/* FetchDatasetsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type FetchDatasetsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *FetchDatasetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/datasets/own][%d] fetchDatasetsInternalServerError  %+v", 500, o.Payload)
}
func (o *FetchDatasetsInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *FetchDatasetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// GetServiceAccountsReader is a Reader for the GetServiceAccounts structure.
type GetServiceAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServiceAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetServiceAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceAccountsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetServiceAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServiceAccountsOK creates a GetServiceAccountsOK with default headers values
func NewGetServiceAccountsOK() *GetServiceAccountsOK {
	return &GetServiceAccountsOK{}
}

/* GetServiceAccountsOK describes a response with status code 200, with default header values.

Successfully retrieved service accounts
*/
type GetServiceAccountsOK struct {
	Payload *models.PaginatedServiceAccountDetails
}

func (o *GetServiceAccountsOK) Error() string {
	return fmt.Sprintf("[GET /serviceaccount/{owner}][%d] getServiceAccountsOK  %+v", 200, o.Payload)
}
func (o *GetServiceAccountsOK) GetPayload() *models.PaginatedServiceAccountDetails {
	return o.Payload
}

func (o *GetServiceAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedServiceAccountDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountsBadRequest creates a GetServiceAccountsBadRequest with default headers values
func NewGetServiceAccountsBadRequest() *GetServiceAccountsBadRequest {
	return &GetServiceAccountsBadRequest{}
}

/* GetServiceAccountsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetServiceAccountsBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetServiceAccountsBadRequest) Error() string {
	return fmt.Sprintf("[GET /serviceaccount/{owner}][%d] getServiceAccountsBadRequest  %+v", 400, o.Payload)
}
func (o *GetServiceAccountsBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetServiceAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountsUnauthorized creates a GetServiceAccountsUnauthorized with default headers values
func NewGetServiceAccountsUnauthorized() *GetServiceAccountsUnauthorized {
	return &GetServiceAccountsUnauthorized{}
}

/* GetServiceAccountsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetServiceAccountsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetServiceAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /serviceaccount/{owner}][%d] getServiceAccountsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetServiceAccountsUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetServiceAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountsForbidden creates a GetServiceAccountsForbidden with default headers values
func NewGetServiceAccountsForbidden() *GetServiceAccountsForbidden {
	return &GetServiceAccountsForbidden{}
}

/* GetServiceAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetServiceAccountsForbidden struct {
	Payload *models.ErrorMessage
}

func (o *GetServiceAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /serviceaccount/{owner}][%d] getServiceAccountsForbidden  %+v", 403, o.Payload)
}
func (o *GetServiceAccountsForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetServiceAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountsNotFound creates a GetServiceAccountsNotFound with default headers values
func NewGetServiceAccountsNotFound() *GetServiceAccountsNotFound {
	return &GetServiceAccountsNotFound{}
}

/* GetServiceAccountsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetServiceAccountsNotFound struct {
	Payload *models.ErrorMessage
}

func (o *GetServiceAccountsNotFound) Error() string {
	return fmt.Sprintf("[GET /serviceaccount/{owner}][%d] getServiceAccountsNotFound  %+v", 404, o.Payload)
}
func (o *GetServiceAccountsNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetServiceAccountsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceAccountsInternalServerError creates a GetServiceAccountsInternalServerError with default headers values
func NewGetServiceAccountsInternalServerError() *GetServiceAccountsInternalServerError {
	return &GetServiceAccountsInternalServerError{}
}

/* GetServiceAccountsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetServiceAccountsInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *GetServiceAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /serviceaccount/{owner}][%d] getServiceAccountsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetServiceAccountsInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetServiceAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
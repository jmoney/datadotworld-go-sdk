// Code generated by go-swagger; DO NOT EDIT.

package cancel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// CancelReader is a Reader for the Cancel structure.
type CancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewCancelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelBadRequest creates a CancelBadRequest with default headers values
func NewCancelBadRequest() *CancelBadRequest {
	return &CancelBadRequest{}
}

/* CancelBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CancelBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *CancelBadRequest) Error() string {
	return fmt.Sprintf("[POST /cancel][%d] cancelBadRequest  %+v", 400, o.Payload)
}
func (o *CancelBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CancelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelNotFound creates a CancelNotFound with default headers values
func NewCancelNotFound() *CancelNotFound {
	return &CancelNotFound{}
}

/* CancelNotFound describes a response with status code 404, with default header values.

Not found
*/
type CancelNotFound struct {
	Payload *models.ErrorMessage
}

func (o *CancelNotFound) Error() string {
	return fmt.Sprintf("[POST /cancel][%d] cancelNotFound  %+v", 404, o.Payload)
}
func (o *CancelNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CancelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package partners

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// StitchUploadManifestReader is a Reader for the StitchUploadManifest structure.
type StitchUploadManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StitchUploadManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStitchUploadManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStitchUploadManifestOK creates a StitchUploadManifestOK with default headers values
func NewStitchUploadManifestOK() *StitchUploadManifestOK {
	return &StitchUploadManifestOK{}
}

/* StitchUploadManifestOK describes a response with status code 200, with default header values.

successful operation
*/
type StitchUploadManifestOK struct {
	Payload *models.DatasetIdentifierResponse
}

func (o *StitchUploadManifestOK) Error() string {
	return fmt.Sprintf("[POST /partners/stitchdata/uploads/{account}/{integrationName}/manifest][%d] stitchUploadManifestOK  %+v", 200, o.Payload)
}
func (o *StitchUploadManifestOK) GetPayload() *models.DatasetIdentifierResponse {
	return o.Payload
}

func (o *StitchUploadManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatasetIdentifierResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

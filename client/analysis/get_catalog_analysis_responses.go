// Code generated by go-swagger; DO NOT EDIT.

package analysis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"datadotworld-go-sdk/models"
)

// GetCatalogAnalysisReader is a Reader for the GetCatalogAnalysis structure.
type GetCatalogAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCatalogAnalysisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCatalogAnalysisBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCatalogAnalysisNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCatalogAnalysisOK creates a GetCatalogAnalysisOK with default headers values
func NewGetCatalogAnalysisOK() *GetCatalogAnalysisOK {
	return &GetCatalogAnalysisOK{}
}

/* GetCatalogAnalysisOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCatalogAnalysisOK struct {
	Payload *models.MetadataResourceDto
}

func (o *GetCatalogAnalysisOK) Error() string {
	return fmt.Sprintf("[GET /metadata/analysis/{owner}/{id}][%d] getCatalogAnalysisOK  %+v", 200, o.Payload)
}
func (o *GetCatalogAnalysisOK) GetPayload() *models.MetadataResourceDto {
	return o.Payload
}

func (o *GetCatalogAnalysisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetadataResourceDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogAnalysisBadRequest creates a GetCatalogAnalysisBadRequest with default headers values
func NewGetCatalogAnalysisBadRequest() *GetCatalogAnalysisBadRequest {
	return &GetCatalogAnalysisBadRequest{}
}

/* GetCatalogAnalysisBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCatalogAnalysisBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetCatalogAnalysisBadRequest) Error() string {
	return fmt.Sprintf("[GET /metadata/analysis/{owner}/{id}][%d] getCatalogAnalysisBadRequest  %+v", 400, o.Payload)
}
func (o *GetCatalogAnalysisBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetCatalogAnalysisBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogAnalysisNotFound creates a GetCatalogAnalysisNotFound with default headers values
func NewGetCatalogAnalysisNotFound() *GetCatalogAnalysisNotFound {
	return &GetCatalogAnalysisNotFound{}
}

/* GetCatalogAnalysisNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetCatalogAnalysisNotFound struct {
	Payload *models.ErrorMessage
}

func (o *GetCatalogAnalysisNotFound) Error() string {
	return fmt.Sprintf("[GET /metadata/analysis/{owner}/{id}][%d] getCatalogAnalysisNotFound  %+v", 404, o.Payload)
}
func (o *GetCatalogAnalysisNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GetCatalogAnalysisNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// SuggestCatalogAnalysisReader is a Reader for the SuggestCatalogAnalysis structure.
type SuggestCatalogAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuggestCatalogAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuggestCatalogAnalysisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSuggestCatalogAnalysisBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSuggestCatalogAnalysisUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSuggestCatalogAnalysisForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuggestCatalogAnalysisNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSuggestCatalogAnalysisUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSuggestCatalogAnalysisInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuggestCatalogAnalysisOK creates a SuggestCatalogAnalysisOK with default headers values
func NewSuggestCatalogAnalysisOK() *SuggestCatalogAnalysisOK {
	return &SuggestCatalogAnalysisOK{}
}

/* SuggestCatalogAnalysisOK describes a response with status code 200, with default header values.

Suggested changes to Analysis updated successfully.
*/
type SuggestCatalogAnalysisOK struct {
	Payload *models.SuccessMessage
}

func (o *SuggestCatalogAnalysisOK) Error() string {
	return fmt.Sprintf("[PATCH /metadata/analysis/suggest/{owner}/{id}][%d] suggestCatalogAnalysisOK  %+v", 200, o.Payload)
}
func (o *SuggestCatalogAnalysisOK) GetPayload() *models.SuccessMessage {
	return o.Payload
}

func (o *SuggestCatalogAnalysisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogAnalysisBadRequest creates a SuggestCatalogAnalysisBadRequest with default headers values
func NewSuggestCatalogAnalysisBadRequest() *SuggestCatalogAnalysisBadRequest {
	return &SuggestCatalogAnalysisBadRequest{}
}

/* SuggestCatalogAnalysisBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SuggestCatalogAnalysisBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogAnalysisBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /metadata/analysis/suggest/{owner}/{id}][%d] suggestCatalogAnalysisBadRequest  %+v", 400, o.Payload)
}
func (o *SuggestCatalogAnalysisBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogAnalysisBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogAnalysisUnauthorized creates a SuggestCatalogAnalysisUnauthorized with default headers values
func NewSuggestCatalogAnalysisUnauthorized() *SuggestCatalogAnalysisUnauthorized {
	return &SuggestCatalogAnalysisUnauthorized{}
}

/* SuggestCatalogAnalysisUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SuggestCatalogAnalysisUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogAnalysisUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /metadata/analysis/suggest/{owner}/{id}][%d] suggestCatalogAnalysisUnauthorized  %+v", 401, o.Payload)
}
func (o *SuggestCatalogAnalysisUnauthorized) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogAnalysisUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogAnalysisForbidden creates a SuggestCatalogAnalysisForbidden with default headers values
func NewSuggestCatalogAnalysisForbidden() *SuggestCatalogAnalysisForbidden {
	return &SuggestCatalogAnalysisForbidden{}
}

/* SuggestCatalogAnalysisForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SuggestCatalogAnalysisForbidden struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogAnalysisForbidden) Error() string {
	return fmt.Sprintf("[PATCH /metadata/analysis/suggest/{owner}/{id}][%d] suggestCatalogAnalysisForbidden  %+v", 403, o.Payload)
}
func (o *SuggestCatalogAnalysisForbidden) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogAnalysisForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogAnalysisNotFound creates a SuggestCatalogAnalysisNotFound with default headers values
func NewSuggestCatalogAnalysisNotFound() *SuggestCatalogAnalysisNotFound {
	return &SuggestCatalogAnalysisNotFound{}
}

/* SuggestCatalogAnalysisNotFound describes a response with status code 404, with default header values.

Not found
*/
type SuggestCatalogAnalysisNotFound struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogAnalysisNotFound) Error() string {
	return fmt.Sprintf("[PATCH /metadata/analysis/suggest/{owner}/{id}][%d] suggestCatalogAnalysisNotFound  %+v", 404, o.Payload)
}
func (o *SuggestCatalogAnalysisNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogAnalysisNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogAnalysisUnprocessableEntity creates a SuggestCatalogAnalysisUnprocessableEntity with default headers values
func NewSuggestCatalogAnalysisUnprocessableEntity() *SuggestCatalogAnalysisUnprocessableEntity {
	return &SuggestCatalogAnalysisUnprocessableEntity{}
}

/* SuggestCatalogAnalysisUnprocessableEntity describes a response with status code 422, with default header values.

Bad request
*/
type SuggestCatalogAnalysisUnprocessableEntity struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogAnalysisUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /metadata/analysis/suggest/{owner}/{id}][%d] suggestCatalogAnalysisUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *SuggestCatalogAnalysisUnprocessableEntity) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogAnalysisUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuggestCatalogAnalysisInternalServerError creates a SuggestCatalogAnalysisInternalServerError with default headers values
func NewSuggestCatalogAnalysisInternalServerError() *SuggestCatalogAnalysisInternalServerError {
	return &SuggestCatalogAnalysisInternalServerError{}
}

/* SuggestCatalogAnalysisInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SuggestCatalogAnalysisInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *SuggestCatalogAnalysisInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /metadata/analysis/suggest/{owner}/{id}][%d] suggestCatalogAnalysisInternalServerError  %+v", 500, o.Payload)
}
func (o *SuggestCatalogAnalysisInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuggestCatalogAnalysisInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

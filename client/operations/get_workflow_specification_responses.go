// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetWorkflowSpecificationReader is a Reader for the GetWorkflowSpecification structure.
type GetWorkflowSpecificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowSpecificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowSpecificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetWorkflowSpecificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowSpecificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowSpecificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowSpecificationOK creates a GetWorkflowSpecificationOK with default headers values
func NewGetWorkflowSpecificationOK() *GetWorkflowSpecificationOK {
	return &GetWorkflowSpecificationOK{}
}

/* GetWorkflowSpecificationOK describes a response with status code 200, with default header values.

Request succeeded. Workflow specification is returned.
*/
type GetWorkflowSpecificationOK struct {
	Payload interface{}
}

func (o *GetWorkflowSpecificationOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationOK  %+v", 200, o.Payload)
}
func (o *GetWorkflowSpecificationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetWorkflowSpecificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowSpecificationForbidden creates a GetWorkflowSpecificationForbidden with default headers values
func NewGetWorkflowSpecificationForbidden() *GetWorkflowSpecificationForbidden {
	return &GetWorkflowSpecificationForbidden{}
}

/* GetWorkflowSpecificationForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type GetWorkflowSpecificationForbidden struct {
}

func (o *GetWorkflowSpecificationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationForbidden ", 403)
}

func (o *GetWorkflowSpecificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowSpecificationNotFound creates a GetWorkflowSpecificationNotFound with default headers values
func NewGetWorkflowSpecificationNotFound() *GetWorkflowSpecificationNotFound {
	return &GetWorkflowSpecificationNotFound{}
}

/* GetWorkflowSpecificationNotFound describes a response with status code 404, with default header values.

Request failed. User does not exist.
*/
type GetWorkflowSpecificationNotFound struct {
}

func (o *GetWorkflowSpecificationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationNotFound ", 404)
}

func (o *GetWorkflowSpecificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowSpecificationInternalServerError creates a GetWorkflowSpecificationInternalServerError with default headers values
func NewGetWorkflowSpecificationInternalServerError() *GetWorkflowSpecificationInternalServerError {
	return &GetWorkflowSpecificationInternalServerError{}
}

/* GetWorkflowSpecificationInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type GetWorkflowSpecificationInternalServerError struct {
}

func (o *GetWorkflowSpecificationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/specification][%d] getWorkflowSpecificationInternalServerError ", 500)
}

func (o *GetWorkflowSpecificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
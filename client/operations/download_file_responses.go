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

// DownloadFileReader is a Reader for the DownloadFile structure.
type DownloadFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDownloadFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadFileOK creates a DownloadFileOK with default headers values
func NewDownloadFileOK(writer io.Writer) *DownloadFileOK {
	return &DownloadFileOK{

		Payload: writer,
	}
}

/* DownloadFileOK describes a response with status code 200, with default header values.

Requests succeeded. The file has been downloaded.
*/
type DownloadFileOK struct {
	Payload io.Writer
}

func (o *DownloadFileOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileOK  %+v", 200, o.Payload)
}
func (o *DownloadFileOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadFileBadRequest creates a DownloadFileBadRequest with default headers values
func NewDownloadFileBadRequest() *DownloadFileBadRequest {
	return &DownloadFileBadRequest{}
}

/* DownloadFileBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type DownloadFileBadRequest struct {
}

func (o *DownloadFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileBadRequest ", 400)
}

func (o *DownloadFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadFileForbidden creates a DownloadFileForbidden with default headers values
func NewDownloadFileForbidden() *DownloadFileForbidden {
	return &DownloadFileForbidden{}
}

/* DownloadFileForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type DownloadFileForbidden struct {
}

func (o *DownloadFileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileForbidden ", 403)
}

func (o *DownloadFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadFileNotFound creates a DownloadFileNotFound with default headers values
func NewDownloadFileNotFound() *DownloadFileNotFound {
	return &DownloadFileNotFound{}
}

/* DownloadFileNotFound describes a response with status code 404, with default header values.

Request failed. `file_name` does not exist .
*/
type DownloadFileNotFound struct {
}

func (o *DownloadFileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileNotFound ", 404)
}

func (o *DownloadFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadFileInternalServerError creates a DownloadFileInternalServerError with default headers values
func NewDownloadFileInternalServerError() *DownloadFileInternalServerError {
	return &DownloadFileInternalServerError{}
}

/* DownloadFileInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type DownloadFileInternalServerError struct {
}

func (o *DownloadFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/workspace/{file_name}][%d] downloadFileInternalServerError ", 500)
}

func (o *DownloadFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
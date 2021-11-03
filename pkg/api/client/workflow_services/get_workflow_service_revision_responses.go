// Code generated by go-swagger; DO NOT EDIT.

package workflow_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetWorkflowServiceRevisionReader is a Reader for the GetWorkflowServiceRevision structure.
type GetWorkflowServiceRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowServiceRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowServiceRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowServiceRevisionOK creates a GetWorkflowServiceRevisionOK with default headers values
func NewGetWorkflowServiceRevisionOK() *GetWorkflowServiceRevisionOK {
	return &GetWorkflowServiceRevisionOK{}
}

/* GetWorkflowServiceRevisionOK describes a response with status code 200, with default header values.

successfully got service revision details
*/
type GetWorkflowServiceRevisionOK struct {
}

func (o *GetWorkflowServiceRevisionOK) Error() string {
	return fmt.Sprintf("[GET /api/functions/namespaces/{namespace}/tree/{workflow}?op=function-revision][%d] getWorkflowServiceRevisionOK ", 200)
}

func (o *GetWorkflowServiceRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
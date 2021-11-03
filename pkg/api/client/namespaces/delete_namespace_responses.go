// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/direktiv/direktiv/pkg/api/models"
)

// DeleteNamespaceReader is a Reader for the DeleteNamespace structure.
type DeleteNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNamespaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNamespaceOK creates a DeleteNamespaceOK with default headers values
func NewDeleteNamespaceOK() *DeleteNamespaceOK {
	return &DeleteNamespaceOK{}
}

/* DeleteNamespaceOK describes a response with status code 200, with default header values.

namespace has been successfully deleted
*/
type DeleteNamespaceOK struct {
	Payload models.OkBody
}

func (o *DeleteNamespaceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/namespaces/{namespace}][%d] deleteNamespaceOK  %+v", 200, o.Payload)
}
func (o *DeleteNamespaceOK) GetPayload() models.OkBody {
	return o.Payload
}

func (o *DeleteNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNamespaceDefault creates a DeleteNamespaceDefault with default headers values
func NewDeleteNamespaceDefault(code int) *DeleteNamespaceDefault {
	return &DeleteNamespaceDefault{
		_statusCode: code,
	}
}

/* DeleteNamespaceDefault describes a response with status code -1, with default header values.

an error has occurred
*/
type DeleteNamespaceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete namespace default response
func (o *DeleteNamespaceDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNamespaceDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/namespaces/{namespace}][%d] deleteNamespace default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNamespaceDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteNamespaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
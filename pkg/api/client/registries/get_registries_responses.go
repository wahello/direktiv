// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRegistriesReader is a Reader for the GetRegistries structure.
type GetRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistriesOK creates a GetRegistriesOK with default headers values
func NewGetRegistriesOK() *GetRegistriesOK {
	return &GetRegistriesOK{}
}

/* GetRegistriesOK describes a response with status code 200, with default header values.

successfully got namespace registries
*/
type GetRegistriesOK struct {
}

func (o *GetRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /api/registries/namespaces/{namespace}][%d] getRegistriesOK ", 200)
}

func (o *GetRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

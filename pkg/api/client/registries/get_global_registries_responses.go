// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetGlobalRegistriesReader is a Reader for the GetGlobalRegistries structure.
type GetGlobalRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGlobalRegistriesOK creates a GetGlobalRegistriesOK with default headers values
func NewGetGlobalRegistriesOK() *GetGlobalRegistriesOK {
	return &GetGlobalRegistriesOK{}
}

/* GetGlobalRegistriesOK describes a response with status code 200, with default header values.

successfully got global registries
*/
type GetGlobalRegistriesOK struct {
}

func (o *GetGlobalRegistriesOK) Error() string {
	return fmt.Sprintf("[GET /api/functions/registries/global][%d] getGlobalRegistriesOK ", 200)
}

func (o *GetGlobalRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

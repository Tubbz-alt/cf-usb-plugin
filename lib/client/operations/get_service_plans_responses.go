package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

type GetServicePlansReader struct {
	formats strfmt.Registry
}

func (o *GetServicePlansReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServicePlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetServicePlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServicePlansOK creates a GetServicePlansOK with default headers values
func NewGetServicePlansOK() *GetServicePlansOK {
	return &GetServicePlansOK{}
}

/*GetServicePlansOK

Sucessfull response
*/
type GetServicePlansOK struct {
	Payload []*models.Plan
}

func (o *GetServicePlansOK) Error() string {
	return fmt.Sprintf("[GET /plans][%d] getServicePlansOK  %+v", 200, o.Payload)
}

func (o *GetServicePlansOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicePlansInternalServerError creates a GetServicePlansInternalServerError with default headers values
func NewGetServicePlansInternalServerError() *GetServicePlansInternalServerError {
	return &GetServicePlansInternalServerError{}
}

/*GetServicePlansInternalServerError

Unexpected error
*/
type GetServicePlansInternalServerError struct {
	Payload string
}

func (o *GetServicePlansInternalServerError) Error() string {
	return fmt.Sprintf("[GET /plans][%d] getServicePlansInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServicePlansInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
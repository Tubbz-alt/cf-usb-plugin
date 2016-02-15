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

type GetInfoReader struct {
	formats strfmt.Registry
}

func (o *GetInfoReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInfoOK creates a GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {
	return &GetInfoOK{}
}

/*GetInfoOK

Successful response
*/
type GetInfoOK struct {
	Payload *models.Info
}

func (o *GetInfoOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoOK  %+v", 200, o.Payload)
}

func (o *GetInfoOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Info)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfoInternalServerError creates a GetInfoInternalServerError with default headers values
func NewGetInfoInternalServerError() *GetInfoInternalServerError {
	return &GetInfoInternalServerError{}
}

/*GetInfoInternalServerError

Unexpected error
*/
type GetInfoInternalServerError struct {
	Payload string
}

func (o *GetInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInfoInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

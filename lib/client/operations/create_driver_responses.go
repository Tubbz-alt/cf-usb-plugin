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

type CreateDriverReader struct {
	formats strfmt.Registry
}

func (o *CreateDriverReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDriverCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreateDriverConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateDriverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDriverCreated creates a CreateDriverCreated with default headers values
func NewCreateDriverCreated() *CreateDriverCreated {
	return &CreateDriverCreated{}
}

/*CreateDriverCreated

Driver created
*/
type CreateDriverCreated struct {
	Payload *models.Driver
}

func (o *CreateDriverCreated) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] createDriverCreated  %+v", 201, o.Payload)
}

func (o *CreateDriverCreated) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Driver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDriverConflict creates a CreateDriverConflict with default headers values
func NewCreateDriverConflict() *CreateDriverConflict {
	return &CreateDriverConflict{}
}

/*CreateDriverConflict

A driver with the same type already exists
*/
type CreateDriverConflict struct {
}

func (o *CreateDriverConflict) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] createDriverConflict ", 409)
}

func (o *CreateDriverConflict) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDriverInternalServerError creates a CreateDriverInternalServerError with default headers values
func NewCreateDriverInternalServerError() *CreateDriverInternalServerError {
	return &CreateDriverInternalServerError{}
}

/*CreateDriverInternalServerError

Unexpected error
*/
type CreateDriverInternalServerError struct {
	Payload string
}

func (o *CreateDriverInternalServerError) Error() string {
	return fmt.Sprintf("[POST /drivers][%d] createDriverInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDriverInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type DeleteDriverReader struct {
	formats strfmt.Registry
}

func (o *DeleteDriverReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteDriverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteDriverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteDriverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDriverNoContent creates a DeleteDriverNoContent with default headers values
func NewDeleteDriverNoContent() *DeleteDriverNoContent {
	return &DeleteDriverNoContent{}
}

/*DeleteDriverNoContent

Driver deleted
*/
type DeleteDriverNoContent struct {
}

func (o *DeleteDriverNoContent) Error() string {
	return fmt.Sprintf("[DELETE /drivers/{driver_id}][%d] deleteDriverNoContent ", 204)
}

func (o *DeleteDriverNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverNotFound creates a DeleteDriverNotFound with default headers values
func NewDeleteDriverNotFound() *DeleteDriverNotFound {
	return &DeleteDriverNotFound{}
}

/*DeleteDriverNotFound

Not Found
*/
type DeleteDriverNotFound struct {
}

func (o *DeleteDriverNotFound) Error() string {
	return fmt.Sprintf("[DELETE /drivers/{driver_id}][%d] deleteDriverNotFound ", 404)
}

func (o *DeleteDriverNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverInternalServerError creates a DeleteDriverInternalServerError with default headers values
func NewDeleteDriverInternalServerError() *DeleteDriverInternalServerError {
	return &DeleteDriverInternalServerError{}
}

/*DeleteDriverInternalServerError

Unexpected error
*/
type DeleteDriverInternalServerError struct {
	Payload string
}

func (o *DeleteDriverInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /drivers/{driver_id}][%d] deleteDriverInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDriverInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
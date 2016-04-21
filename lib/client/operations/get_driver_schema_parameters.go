package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDriverSchemaParams creates a new GetDriverSchemaParams object
// with the default values initialized.
func NewGetDriverSchemaParams() *GetDriverSchemaParams {
	var ()
	return &GetDriverSchemaParams{}
}

/*GetDriverSchemaParams contains all the parameters to send to the API endpoint
for the get driver schema operation typically these are written to a http.Request
*/
type GetDriverSchemaParams struct {

	/*DriverID
	  Driver ID

	*/
	DriverID string
}

// WithDriverID adds the driverId to the get driver schema params
func (o *GetDriverSchemaParams) WithDriverID(DriverID string) *GetDriverSchemaParams {
	o.DriverID = DriverID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetDriverSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param driver_id
	if err := r.SetPathParam("driver_id", o.DriverID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

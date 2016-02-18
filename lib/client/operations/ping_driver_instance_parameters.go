package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewPingDriverInstanceParams creates a new PingDriverInstanceParams object
// with the default values initialized.
func NewPingDriverInstanceParams() *PingDriverInstanceParams {
	var ()
	return &PingDriverInstanceParams{}
}

/*PingDriverInstanceParams contains all the parameters to send to the API endpoint
for the ping driver instance operation typically these are written to a http.Request
*/
type PingDriverInstanceParams struct {

	/*DriverInstanceID
	  Driver Instance ID

	*/
	DriverInstanceID string
}

// WithDriverInstanceID adds the driverInstanceId to the ping driver instance params
func (o *PingDriverInstanceParams) WithDriverInstanceID(driverInstanceId string) *PingDriverInstanceParams {
	o.DriverInstanceID = driverInstanceId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PingDriverInstanceParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param driver_instance_id
	if err := r.SetPathParam("driver_instance_id", o.DriverInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
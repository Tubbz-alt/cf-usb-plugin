package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// NewUpdateDriverInstanceParams creates a new UpdateDriverInstanceParams object
// with the default values initialized.
func NewUpdateDriverInstanceParams() *UpdateDriverInstanceParams {
	var ()
	return &UpdateDriverInstanceParams{}
}

/*UpdateDriverInstanceParams contains all the parameters to send to the API endpoint
for the update driver instance operation typically these are written to a http.Request
*/
type UpdateDriverInstanceParams struct {

	/*DriverConfig
	  Add driver_config

	*/
	DriverConfig *models.DriverInstance
	/*DriverInstanceID
	  Driver Instance ID

	*/
	DriverInstanceID string
}

// WithDriverConfig adds the driverConfig to the update driver instance params
func (o *UpdateDriverInstanceParams) WithDriverConfig(DriverConfig *models.DriverInstance) *UpdateDriverInstanceParams {
	o.DriverConfig = DriverConfig
	return o
}

// WithDriverInstanceID adds the driverInstanceId to the update driver instance params
func (o *UpdateDriverInstanceParams) WithDriverInstanceID(DriverInstanceID string) *UpdateDriverInstanceParams {
	o.DriverInstanceID = DriverInstanceID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDriverInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.DriverConfig == nil {
		o.DriverConfig = new(models.DriverInstance)
	}

	if err := r.SetBodyParam(o.DriverConfig); err != nil {
		return err
	}

	// path param driver_instance_id
	if err := r.SetPathParam("driver_instance_id", o.DriverInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

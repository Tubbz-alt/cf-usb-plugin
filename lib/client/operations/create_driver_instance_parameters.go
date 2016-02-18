package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// NewCreateDriverInstanceParams creates a new CreateDriverInstanceParams object
// with the default values initialized.
func NewCreateDriverInstanceParams() *CreateDriverInstanceParams {
	var ()
	return &CreateDriverInstanceParams{}
}

/*CreateDriverInstanceParams contains all the parameters to send to the API endpoint
for the create driver instance operation typically these are written to a http.Request
*/
type CreateDriverInstanceParams struct {

	/*DriverInstance
	  driver instance to be created

	*/
	DriverInstance *models.DriverInstance
}

// WithDriverInstance adds the driverInstance to the create driver instance params
func (o *CreateDriverInstanceParams) WithDriverInstance(driverInstance *models.DriverInstance) *CreateDriverInstanceParams {
	o.DriverInstance = driverInstance
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDriverInstanceParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.DriverInstance == nil {
		o.DriverInstance = new(models.DriverInstance)
	}

	if err := r.SetBodyParam(o.DriverInstance); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
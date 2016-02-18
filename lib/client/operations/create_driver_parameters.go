package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// NewCreateDriverParams creates a new CreateDriverParams object
// with the default values initialized.
func NewCreateDriverParams() *CreateDriverParams {
	var ()
	return &CreateDriverParams{}
}

/*CreateDriverParams contains all the parameters to send to the API endpoint
for the create driver operation typically these are written to a http.Request
*/
type CreateDriverParams struct {

	/*Driver
	  Driver to be created

	*/
	Driver *models.Driver
}

// WithDriver adds the driver to the create driver params
func (o *CreateDriverParams) WithDriver(driver *models.Driver) *CreateDriverParams {
	o.Driver = driver
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDriverParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Driver == nil {
		o.Driver = new(models.Driver)
	}

	if err := r.SetBodyParam(o.Driver); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
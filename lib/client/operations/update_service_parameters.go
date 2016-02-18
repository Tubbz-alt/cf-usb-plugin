package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// NewUpdateServiceParams creates a new UpdateServiceParams object
// with the default values initialized.
func NewUpdateServiceParams() *UpdateServiceParams {
	var ()
	return &UpdateServiceParams{}
}

/*UpdateServiceParams contains all the parameters to send to the API endpoint
for the update service operation typically these are written to a http.Request
*/
type UpdateServiceParams struct {

	/*Service
	  Update service

	*/
	Service *models.Service
	/*ServiceID
	  ID of the service

	*/
	ServiceID string
}

// WithService adds the service to the update service params
func (o *UpdateServiceParams) WithService(service *models.Service) *UpdateServiceParams {
	o.Service = service
	return o
}

// WithServiceID adds the serviceId to the update service params
func (o *UpdateServiceParams) WithServiceID(serviceId string) *UpdateServiceParams {
	o.ServiceID = serviceId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Service == nil {
		o.Service = new(models.Service)
	}

	if err := r.SetBodyParam(o.Service); err != nil {
		return err
	}

	// path param service_id
	if err := r.SetPathParam("service_id", o.ServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetServicePlanParams creates a new GetServicePlanParams object
// with the default values initialized.
func NewGetServicePlanParams() *GetServicePlanParams {
	var ()
	return &GetServicePlanParams{}
}

/*GetServicePlanParams contains all the parameters to send to the API endpoint
for the get service plan operation typically these are written to a http.Request
*/
type GetServicePlanParams struct {

	/*PlanID
	  ID of the plan

	*/
	PlanID string
}

// WithPlanID adds the planId to the get service plan params
func (o *GetServicePlanParams) WithPlanID(planId string) *GetServicePlanParams {
	o.PlanID = planId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicePlanParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param plan_id
	if err := r.SetPathParam("plan_id", o.PlanID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Plan plan

swagger:model plan
*/
type Plan struct {

	/* description
	 */
	Description string `json:"description,omitempty"`

	/* dial id

	Required: true
	Max Length: 36
	Min Length: 36
	*/
	DialID *string `json:"dial_id"`

	/* free
	 */
	Free bool `json:"free,omitempty"`

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* name

	Required: true
	Max Length: 50
	Min Length: 3
	*/
	Name *string `json:"name"`
}

// Validate validates this plan
func (m *Plan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDialID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plan) validateDialID(formats strfmt.Registry) error {

	if err := validate.Required("dial_id", "body", m.DialID); err != nil {
		return err
	}

	if err := validate.MinLength("dial_id", "body", string(*m.DialID), 36); err != nil {
		return err
	}

	if err := validate.MaxLength("dial_id", "body", string(*m.DialID), 36); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

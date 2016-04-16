package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*ResourceTypeOut resource type out

swagger:model resource_type_out
*/
type ResourceTypeOut struct {

	/* description
	 */
	Description string `json:"description,omitempty"`

	/* id

	Required: true
	Max Length: 36
	Min Length: 36
	*/
	ID *string `json:"id"`

	/* name

	Required: true
	Min Length: 1
	*/
	Name *string `json:"name"`
}

// Validate validates this resource type out
func (m *ResourceTypeOut) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *ResourceTypeOut) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 36); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 36); err != nil {
		return err
	}

	return nil
}

func (m *ResourceTypeOut) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}
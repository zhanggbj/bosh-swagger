package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// BaremetalServers baremetal servers
// swagger:model BaremetalServers
type BaremetalServers struct {

	// data
	Data *BaremetalServer `json:"data,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this baremetal servers
func (m *BaremetalServers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaremetalServers) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

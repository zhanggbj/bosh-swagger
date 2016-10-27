package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// InfoData info data
// swagger:model InfoData
type InfoData struct {

	// the name of the BMP server
	Name string `json:"name,omitempty"`

	// the version string of the BMP server
	Version string `json:"version,omitempty"`
}

// Validate validates this info data
func (m *InfoData) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// BaremetalsResponse baremetals response
// swagger:model BaremetalsResponse
type BaremetalsResponse struct {

	// data
	Data *BaremetalsResponseData `json:"data,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this baremetals response
func (m *BaremetalsResponse) Validate(formats strfmt.Registry) error {
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

func (m *BaremetalsResponse) validateData(formats strfmt.Registry) error {

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

// BaremetalsResponseData baremetals response data
// swagger:model BaremetalsResponseData
type BaremetalsResponseData struct {

	// task id
	TaskID int32 `json:"task_id,omitempty"`
}

// Validate validates this baremetals response data
func (m *BaremetalsResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

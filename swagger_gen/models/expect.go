// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Expect expect
// swagger:model Expect
type Expect struct {

	// amqp
	Amqp *ExpectAMQP `json:"amqp,omitempty"`

	// a go template that determines if this behavior triggers
	Condition string `json:"condition,omitempty"`

	// http
	HTTP *ExpectHTTP `json:"http,omitempty"`

	// kafka
	Kafka *ExpectKafka `json:"kafka,omitempty"`
}

// Validate validates this expect
func (m *Expect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmqp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKafka(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Expect) validateAmqp(formats strfmt.Registry) error {

	if swag.IsZero(m.Amqp) { // not required
		return nil
	}

	if m.Amqp != nil {
		if err := m.Amqp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amqp")
			}
			return err
		}
	}

	return nil
}

func (m *Expect) validateHTTP(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {
		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

func (m *Expect) validateKafka(formats strfmt.Registry) error {

	if swag.IsZero(m.Kafka) { // not required
		return nil
	}

	if m.Kafka != nil {
		if err := m.Kafka.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kafka")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Expect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Expect) UnmarshalBinary(b []byte) error {
	var res Expect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
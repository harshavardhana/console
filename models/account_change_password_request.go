// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountChangePasswordRequest account change password request
//
// swagger:model accountChangePasswordRequest
type AccountChangePasswordRequest struct {

	// current secret key
	// Required: true
	CurrentSecretKey *string `json:"current_secret_key"`

	// new secret key
	// Required: true
	NewSecretKey *string `json:"new_secret_key"`
}

// Validate validates this account change password request
func (m *AccountChangePasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountChangePasswordRequest) validateCurrentSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("current_secret_key", "body", m.CurrentSecretKey); err != nil {
		return err
	}

	return nil
}

func (m *AccountChangePasswordRequest) validateNewSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("new_secret_key", "body", m.NewSecretKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountChangePasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountChangePasswordRequest) UnmarshalBinary(b []byte) error {
	var res AccountChangePasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
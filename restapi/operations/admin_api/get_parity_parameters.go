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

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetParityParams creates a new GetParityParams object
//
// There are no default values defined in the spec.
func NewGetParityParams() GetParityParams {

	return GetParityParams{}
}

// GetParityParams contains all the bound params for the get parity operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetParity
type GetParityParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  Minimum: 1
	  In: path
	*/
	DisksPerNode int64
	/*
	  Required: true
	  Minimum: 2
	  In: path
	*/
	Nodes int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetParityParams() beforehand.
func (o *GetParityParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDisksPerNode, rhkDisksPerNode, _ := route.Params.GetOK("disksPerNode")
	if err := o.bindDisksPerNode(rDisksPerNode, rhkDisksPerNode, route.Formats); err != nil {
		res = append(res, err)
	}

	rNodes, rhkNodes, _ := route.Params.GetOK("nodes")
	if err := o.bindNodes(rNodes, rhkNodes, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDisksPerNode binds and validates parameter DisksPerNode from path.
func (o *GetParityParams) bindDisksPerNode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("disksPerNode", "path", "int64", raw)
	}
	o.DisksPerNode = value

	if err := o.validateDisksPerNode(formats); err != nil {
		return err
	}

	return nil
}

// validateDisksPerNode carries on validations for parameter DisksPerNode
func (o *GetParityParams) validateDisksPerNode(formats strfmt.Registry) error {

	if err := validate.MinimumInt("disksPerNode", "path", o.DisksPerNode, 1, false); err != nil {
		return err
	}

	return nil
}

// bindNodes binds and validates parameter Nodes from path.
func (o *GetParityParams) bindNodes(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("nodes", "path", "int64", raw)
	}
	o.Nodes = value

	if err := o.validateNodes(formats); err != nil {
		return err
	}

	return nil
}

// validateNodes carries on validations for parameter Nodes
func (o *GetParityParams) validateNodes(formats strfmt.Registry) error {

	if err := validate.MinimumInt("nodes", "path", o.Nodes, 2, false); err != nil {
		return err
	}

	return nil
}

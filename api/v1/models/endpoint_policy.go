// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointPolicy Policy information of an endpoint
//
// +k8s:deepcopy-gen=true
//
// swagger:model EndpointPolicy
type EndpointPolicy struct {

	// List of identities to which this endpoint is allowed to communicate
	//
	AllowedEgressIdentities []int64 `json:"allowed-egress-identities"`

	// List of identities allowed to communicate to this endpoint
	//
	AllowedIngressIdentities []int64 `json:"allowed-ingress-identities"`

	// Build number of calculated policy in use
	Build int64 `json:"build,omitempty"`

	// cidr policy
	CidrPolicy *CIDRPolicy `json:"cidr-policy,omitempty"`

	// Own identity of endpoint
	ID int64 `json:"id,omitempty"`

	// l4
	L4 *L4Policy `json:"l4,omitempty"`

	// Whether policy enforcement is enabled (ingress, egress, both or none)
	PolicyEnabled EndpointPolicyEnabled `json:"policy-enabled,omitempty"`

	// The agent-local policy revision
	PolicyRevision int64 `json:"policy-revision,omitempty"`
}

// Validate validates this endpoint policy
func (m *EndpointPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidrPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointPolicy) validateCidrPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.CidrPolicy) { // not required
		return nil
	}

	if m.CidrPolicy != nil {
		if err := m.CidrPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cidr-policy")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointPolicy) validateL4(formats strfmt.Registry) error {

	if swag.IsZero(m.L4) { // not required
		return nil
	}

	if m.L4 != nil {
		if err := m.L4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l4")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointPolicy) validatePolicyEnabled(formats strfmt.Registry) error {

	if swag.IsZero(m.PolicyEnabled) { // not required
		return nil
	}

	if err := m.PolicyEnabled.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policy-enabled")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointPolicy) UnmarshalBinary(b []byte) error {
	var res EndpointPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

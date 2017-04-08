package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// NodePool node pool
// swagger:model NodePool
type NodePool struct {

	// A discount strategy indicates the type of discount to be associated with the node pool. This might affect the availability of the nodes in the pools in case of preemptible or spot instances.
	// Possible values are "none", "aggressive", "moderate", "reasonable" #TODO naming should be "reasonable" :-D
	//
	// Required: true
	DiscountStrategy *string `json:"discount_strategy"`

	// Type of the instance to use for the nodes in the pool. All the nodes in the pool share the same instance types
	// Required: true
	InstanceType *string `json:"instance_type"`

	// Name of the node pool
	// Required: true
	Name *string `json:"name"`

	// Profile used for the node pool. Possible values are "worker/default", "worker/database", "worker/gpu", "master". The "master" profile identifies the pool containing the cluster master
	// Required: true
	Profile *string `json:"profile"`
}

// Validate validates this node pool
func (m *NodePool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodePool) validateDiscountStrategy(formats strfmt.Registry) error {

	if err := validate.Required("discount_strategy", "body", m.DiscountStrategy); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instance_type", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NodePool) validateProfile(formats strfmt.Registry) error {

	if err := validate.Required("profile", "body", m.Profile); err != nil {
		return err
	}

	return nil
}

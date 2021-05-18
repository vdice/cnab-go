package bundle

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// Parameter defines a single parameter for a CNAB bundle
type Parameter struct {
	Definition  string    `json:"definition" yaml:"definition"`
	ApplyTo     []string  `json:"applyTo,omitempty" yaml:"applyTo,omitempty"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Destination *Location `json:"destination" yaml:"destination"`
	Required    bool      `json:"required,omitempty" yaml:"required,omitempty"`
}

// GetApplyTo returns the list of actions that the Parameter applies to.
func (p *Parameter) GetApplyTo() []string {
	return p.ApplyTo
}

// AppliesTo returns a boolean value specifying whether or not
// the Parameter applies to the provided action
func (p *Parameter) AppliesTo(action string) bool {
	return AppliesTo(p, action)
}

// Validate a Parameter
func (p *Parameter) Validate(name string, bun Bundle) error {
	if p.Definition == "" {
		return errors.New("parameter definition must be provided")
	}

	if p.Destination == nil {
		return errors.New("parameter destination must be provided")
	}
	if err := p.Destination.Validate(); err != nil {
		return errors.Wrap(err, "invalid parameter destination")
	}

	// Validate parameter definition schema
	schema, ok := bun.Definitions[p.Definition]
	if !ok {
		return fmt.Errorf("unable to find definition for %s", name)
	}
	if _, err := schema.ValidateSchema(); err != nil {
		return errors.Wrapf(err, "invalid definition schema for parameter %q", name)
	}

	// Validate default against definition schema, if exists
	var valResult *multierror.Error
	if schema.Default != nil {
		valErrs, err := schema.Validate(schema.Default)
		if err != nil {
			valResult = multierror.Append(valResult, errors.Wrapf(err, "encountered an error validating parameter %s", name))
		}
		for _, valErr := range valErrs {
			valResult = multierror.Append(valResult, fmt.Errorf("encountered an error validating the default value %v for parameter %q: %v", schema.Default, name, valErr.Error))
		}
	}
	return valResult.ErrorOrNil()
}

package bundle

// Credential represents the definition of a CNAB credential
type Credential struct {
	Location    `yaml:",inline"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool   `json:"required,omitempty" yaml:"required,omitempty"`
}

// Validate a Credential
func (c *Credential) Validate() error {
	return c.Location.Validate()
}

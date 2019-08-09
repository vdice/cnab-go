package driver

import (
	"fmt"

	"github.com/docker/go/canonical/json"

	"github.com/deislabs/cnab-go/driver/command"
	"github.com/deislabs/cnab-go/driver/docker"
	"github.com/deislabs/cnab-go/driver/kubernetes"
	"github.com/deislabs/cnab-go/driver/operation"
)

// ResolvedCred is a credential that has been resolved and is ready for injection into the runtime.
type ResolvedCred struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Driver is capable of running a invocation image
type Driver interface {
	// Run executes the operation inside of the invocation image
	Run(*operation.Operation) (operation.OperationResult, error)
	// Handles receives an ImageType* and answers whether this driver supports that type
	Handles(string) bool
}

// Configurable drivers can explain their configuration, and have it explicitly set
type Configurable interface {
	// Config returns a map of configuration names and values that can be set via environment variable
	Config() map[string]string
	// SetConfig allows setting configuration, where name corresponds to the key in Config, and value is
	// the value to be set.
	SetConfig(map[string]string)
}

// DebugDriver prints the information passed to a driver
//
// It does not ever run the image.
type DebugDriver struct {
	config map[string]string
}

// Run executes the operation on the Debug driver
func (d *DebugDriver) Run(op *operation.Operation) (operation.OperationResult, error) {
	data, err := json.MarshalIndent(op, "", "  ")
	if err != nil {
		return operation.OperationResult{}, err
	}
	fmt.Fprintln(op.Out, string(data))
	return operation.OperationResult{}, nil
}

// Handles always returns true, effectively claiming to work for any image type
func (d *DebugDriver) Handles(dt string) bool {
	return true
}

// Config returns the configuration help text
func (d *DebugDriver) Config() map[string]string {
	return map[string]string{
		"VERBOSE": "Increase verbosity. true, false are supported values",
	}
}

// SetConfig sets configuration for this driver
func (d *DebugDriver) SetConfig(settings map[string]string) {
	d.config = settings
}

// Lookup takes a driver name and tries to resolve the most pertinent driver.
func Lookup(name string) (Driver, error) {
	switch name {
	case "docker":
		return &docker.Driver{}, nil
	case "kubernetes", "k8s":
		return &kubernetes.Driver{}, nil
	case "debug":
		return &DebugDriver{}, nil
	default:
		cmddriver := &command.Driver{Name: name}
		if cmddriver.CheckDriverExists() {
			return cmddriver, nil
		}

		return nil, fmt.Errorf("unsupported driver or driver not found in PATH: %s", name)
	}
}

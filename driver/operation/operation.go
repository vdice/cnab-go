package operation

import "io"

// Operation describes the data passed into the driver to run an operation
type Operation struct {
	// Installation is the name of this installation
	Installation string `json:"installation_name"`
	// The revision ID for this installation
	Revision string `json:"revision"`
	// Action is the action to be performed
	Action string `json:"action"`
	// Parameters are the parameters to be injected into the container
	Parameters map[string]interface{} `json:"parameters"`
	// Image is the invocation image
	Image string `json:"image"`
	// ImageType is the type of image.
	ImageType string `json:"image_type"`
	// Environment contains environment variables that should be injected into the invocation image
	Environment map[string]string `json:"environment"`
	// Files contains files that should be injected into the invocation image.
	Files map[string]string `json:"files"`
	// Outputs is a list of paths starting with `/cnab/app/outputs` that the driver should return the contents of in the OperationResult.
	Outputs []string `json:"outputs"`
	// Output stream for log messages from the driver
	Out io.Writer `json:"-"`
}

// OperationResult is the output of the Driver running an Operation.
type OperationResult struct {
	// Outputs is a map from the container path of an output file to its contents (i.e. /cnab/app/outputs/...).
	Outputs map[string]string
}
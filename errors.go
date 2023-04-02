package fluentio

import "errors"

var (
	// ErrNoConfigProvided is returned when no configuration is provided.
	ErrNoConfigProvided = errors.New("no config provided")
)

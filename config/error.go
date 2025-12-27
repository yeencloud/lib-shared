package config

import (
	"fmt"
)

// MARK: - ConfigNotInitializedError
type ConfigurationNotInitializedError struct {
}

func (e ConfigurationNotInitializedError) Error() string {
	return "configuration is not initialized"
}

func (e ConfigurationNotInitializedError) TroubleshootingTip() string {
	return "initialize configuration with config.Init()"
}

// MARK: - MissingConfigValueError
type MissingConfigValueError struct {
	Key string
}

func (e MissingConfigValueError) Error() string {
	return "missing config value for key " + e.Key
}

// MARK: - UnsupportedConfigTypeError
type UnsupportedConfigTypeError struct {
	Type string

	Variable       string
	AvailableTypes []string
}

func (e UnsupportedConfigTypeError) Error() string {
	return "unsupported type " + e.Type + " for variable " + e.Variable
}

func (e UnsupportedConfigTypeError) TroubleshootingTip() string {
	return fmt.Sprintf("Use any of these supported types: %v", e.AvailableTypes)
}

// MARK: - UnsupportedValueForConversionError
type UnsupportedValueForConversionError struct {
	Value string

	FromType string
	ToType   string
}

func (e UnsupportedValueForConversionError) Error() string {
	return fmt.Sprintf("failed to convert %s for conversion from %s to %s", e.Value, e.FromType, e.ToType)
}

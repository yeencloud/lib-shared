package config

import (
	"fmt"
)

type Secret struct {
	Value string
}

func (s Secret) String() string {
	return fmt.Sprintf("***")
}

func NewSecret(value string) Secret {
	return Secret{Value: value}
}

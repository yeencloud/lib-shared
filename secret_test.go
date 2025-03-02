package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecret(t *testing.T) {
	tests := []struct {
		test string

		value string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			secret := NewSecret(tt.value)
			assert.Equal(t, secret.Value, tt.value)
			assert.Equal(t, secret.String(), "***")
		})
	}
}

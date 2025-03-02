package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceName(t *testing.T) {
	tests := []struct {
		test string

		serviceName string
	}{
		{
			test:        "Valid service name",
			serviceName: "test-service",
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			SetServiceName(tt.test)
			assert.Equal(t, tt.test, string(GetServiceName()))
		})
	}
}

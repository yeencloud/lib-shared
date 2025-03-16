package namespace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricKey(t *testing.T) {
	tests := []struct {
		name   string
		path   Namespace
		expect string
	}{
		{
			name:   "no key",
			path:   Namespace{},
			expect: "unknown",
		},
		{
			name: "no parent",
			path: Namespace{
				Identifier: "test",
			},
			expect: "test",
		},
		{
			name: "with parent",
			path: Namespace{
				Parent: &Namespace{
					Identifier: "parent",
				},
				Identifier: "test",
			},
			expect: "parent_test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.path.MetricKey(), tt.expect)
		})
	}
}

func TestRoot(t *testing.T) {
	tests := []struct {
		name   string
		path   Namespace
		expect Namespace
	}{
		{
			name: "no parent",
			path: Namespace{
				Identifier: "test",
			},
			expect: Namespace{
				Identifier: "test",
			},
		},
		{
			name: "with parent",
			path: Namespace{
				Parent: &Namespace{
					Identifier: "parent",
				},
				Identifier: "test",
			},
			expect: Namespace{
				Identifier: "parent",
			},
		},
		{
			name:   "nothing",
			path:   Namespace{},
			expect: Namespace{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.path.Root(), tt.expect)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		path   Namespace
		expect string
	}{
		{
			name: "no parent",
			path: Namespace{
				Identifier: "test",
			},
			expect: "test",
		},
		{
			name: "with parent",
			path: Namespace{
				Parent: &Namespace{
					Identifier: "parent",
				},
				Identifier: "test",
			},
			expect: "parent.test",
		},
		{
			name: "no identifier",
			path: Namespace{
				Parent: &Namespace{
					Identifier: "parent",
				},
			},
			expect: "parent.unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.path.String(), tt.expect)
		})
	}
}

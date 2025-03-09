package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricKey(t *testing.T) {
	tests := []struct {
		name   string
		path   Path
		expect string
	}{
		{
			name:   "no key",
			path:   Path{},
			expect: "unknown",
		},
		{
			name: "no parent",
			path: Path{
				Identifier: "test",
			},
			expect: "test",
		},
		{
			name: "with parent",
			path: Path{
				Parent: &Path{
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
		path   Path
		expect Path
	}{
		{
			name: "no parent",
			path: Path{
				Identifier: "test",
			},
			expect: Path{
				Identifier: "test",
			},
		},
		{
			name: "with parent",
			path: Path{
				Parent: &Path{
					Identifier: "parent",
				},
				Identifier: "test",
			},
			expect: Path{
				Identifier: "parent",
			},
		},
		{
			name:   "nothing",
			path:   Path{},
			expect: Path{},
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
		path   Path
		expect string
	}{
		{
			name: "no parent",
			path: Path{
				Identifier: "test",
			},
			expect: "test",
		},
		{
			name: "with parent",
			path: Path{
				Parent: &Path{
					Identifier: "parent",
				},
				Identifier: "test",
			},
			expect: "parent.test",
		},
		{
			name: "no identifier",
			path: Path{
				Parent: &Path{
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

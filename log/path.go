package log

import (
	"strings"
)

const separator = "."

type Path struct {
	Parent     *Path
	Identifier string

	IsMetricTag bool
}

// MARK: - Functions
func (l Path) String() string {
	parent := ""
	identifier := l.Identifier

	if l.Parent != nil {
		parent = l.Parent.String()
	}

	if identifier == "" {
		identifier = "unknown"
	}

	if parent == "" {
		return identifier
	}

	return parent + separator + identifier
}

func (l Path) Root() Path {
	if l.Parent == nil {
		return l
	}

	return l.Parent.Root()
}

func (l Path) MetricKey() string {
	key := l.String()

	replaced := strings.ReplaceAll(key, ".", "_")

	return replaced
	//	return strings.ReplaceAll(key[len(l.Root().String()):], ".", "_")
}

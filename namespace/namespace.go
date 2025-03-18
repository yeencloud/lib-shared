package namespace

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

const separator = "."

type Namespace struct {
	Parent     *Namespace
	Identifier string

	IsMetricTag bool
}

type NamespaceValue struct {
	Namespace Namespace
	Value     any
}

func (nv NamespaceValue) String() string {
	return fmt.Sprintf("%v", nv.Value)
}

func (nv NamespaceValue) MetricTag() bool {
	return nv.Namespace.IsMetricTag
}

func (nv NamespaceValue) AsField(entry *logrus.Entry) *logrus.Entry {
	return entry.WithField(nv.Namespace.MetricKey(), nv)
}

// MARK: - Functions
func (l Namespace) WithValue(value any) NamespaceValue {
	return NamespaceValue{
		Namespace: l,
		Value:     value,
	}
}

func (l Namespace) String() string {
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

func (l Namespace) Root() Namespace {
	if l.Parent == nil {
		return l
	}

	return l.Parent.Root()
}

func (l Namespace) MetricKey() string {
	key := l.String()

	replaced := strings.ReplaceAll(key, ".", "_")

	return replaced
}

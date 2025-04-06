package metrics

import (
	"github.com/yeencloud/lib-shared/namespace"
)

var MetricsPointKey = "metrics_point"
var MetricsValuesKey = "metrics_values"

var CorrelationIdKey = namespace.Namespace{Identifier: "correlation_id", IsMetricTag: true}

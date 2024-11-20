package metrics

import (
	"context"

	"github.com/costa92/micros-service/internal/pkg/bootstrap"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

var ProviderSet = wire.NewSet(InitMetrics)

type Metrics struct {
	MetricRequests metric.Int64Counter
	MetricSeconds  metric.Float64Histogram
	Meter          metric.Meter
	AppInfo        *bootstrap.AppInfo
}

func NewMetrics(appInfo *bootstrap.AppInfo) *Metrics {
	return &Metrics{
		AppInfo: appInfo,
	}
}

func InitMetrics(appInfo bootstrap.AppInfo) (*Metrics, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	// tel_scope_info{otel_scope_name="order-server",otel_scope_version=""} 1
	otel.SetMeterProvider(provider)
	meter := provider.Meter(appInfo.Name, metric.WithInstrumentationVersion(appInfo.Version))

	metricRequests, err := metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	if err != nil {
		return nil, err
	}

	metricSeconds, err := metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
	if err != nil {
		return nil, err
	}

	return &Metrics{
		MetricRequests: metricRequests,
		MetricSeconds:  metricSeconds,
		Meter:          meter,
	}, nil
}

// IncrementOrderCount increments the order count metric.
func (m *Metrics) IncrementOrderCount(ctx context.Context) {
	counter, _ := m.Meter.Int64Counter("req_order_count", metric.WithUnit("1"), metric.WithDescription("order count"))
	counter.Add(ctx, 1)
}

// IncrementLabelOrderCount increments the order count metric with a label.
func (m *Metrics) IncrementLabelOrderCount(ctx context.Context, key, value string) {
	counter, _ := m.Meter.Int64Counter("req_order_count", metric.WithUnit("1"), metric.WithDescription("order count"))
	counter.Add(ctx, 1, metric.WithAttributes(attribute.String(key, value)))
}

// RecordOrderCount records the order count metric.
func (m *Metrics) RecordOrderCount(ctx context.Context, value int64) {
	counter, _ := m.Meter.Int64Counter("req_order_count", metric.WithUnit("1"), metric.WithDescription("order count"))
	counter.Add(ctx, value)
}

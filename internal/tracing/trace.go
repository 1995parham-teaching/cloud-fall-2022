package tracing

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

// nolint: ireturn
func New(cfg Config) trace.Tracer {
	if !cfg.Enabled {
		return trace.NewNoopTracerProvider().Tracer("")
	}

	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(cfg.Agent.Host),
			jaeger.WithAgentPort(cfg.Agent.Port),
		),
	)
	if err != nil {
		log.Fatalf("failed to initialize export pipeline: %v", err)
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			semconv.ServiceNamespaceKey.String("1995parham-teaching"),
			semconv.ServiceNameKey.String("cloud-fall-2022"),
		),
	)
	if err != nil {
		panic(err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(cfg.Ratio))),
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	// register the TraceContext propagator globally.
	var tc propagation.TraceContext

	otel.SetTextMapPropagator(tc)

	tracer := otel.Tracer("1995parham-teaching/cloud-fall-2022")

	return tracer
}

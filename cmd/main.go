package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"os"

	"github.com/mequq/go-grpc-http-template/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// ...
	cfg := &config.ViperConfig{}
	conf, err := config.NewConfig("config.yaml")
	if err != nil {

		panic(err)
	}
	// load config
	if err := conf.Load(cfg); err != nil {
		panic(err)
	}

	logger := initZapLogger(cfg)
	logger.Debug("config", zap.Any("cfg", cfg))

	shutdown, err := initProvider()
	if err != nil {
		logger.Warn("failed to init tracer  provider", zap.Error(err))

	}
	defer shutdown(context.Background())

	wireApp, err := wireApp(cfg, logger)
	if err != nil {
		logger.Error("failed to init app", zap.Error(err))
		panic(err)
	}

	go func() {
		logger.Info("starting grpc server")
		if err := wireApp.RunGRPC(); err != nil {
			logger.Error("failed to run grpc server", zap.Error(err))
			panic(err)
		}
	}()
	time.Sleep(time.Second * 10)

	if err := wireApp.RunHTTP(); err != nil {
		logger.Error("failed to run http server", zap.Error(err))
		panic(err)
	}

}

func initProvider() (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceName("test-service"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// If the OpenTelemetry Collector is running on a local cluster (minikube or
	// microk8s), it should be accessible through the NodePort service at the
	// `localhost:30080` endpoint. Otherwise, replace `localhost` with the
	// endpoint of your cluster. If you run the app inside k8s, then you can
	// probably connect directly to the service through dns.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "localhost:4317",
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerProvider.Shutdown, nil
}

// func initTracer(cfg *config.ViperConfig) (func(context.Context) error, error) {
// 	// create a new zipkin exporter
// 	exporter, err := zipkin.New(
// 		cfg.Observability.Tracing.Zipkin.Url,
// 		// zipkin.WithLogger(log.New(os.Stdout, "zipkin: ", log.LstdFlags)),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	batcher := sdktrace.NewBatchSpanProcessor(exporter)
// 	// create a new trace provider
// 	p := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))
// 	otel.SetTextMapPropagator(p)

// 	tp := sdktrace.NewTracerProvider(
// 		sdktrace.WithSpanProcessor(batcher),
// 		sdktrace.WithResource(resource.NewWithAttributes(
// 			semconv.SchemaURL,
// 			semconv.ServiceNameKey.String("app"),
// 			semconv.DeploymentEnvironmentKey.String("production"),
// 		),
// 		),
// 	)
// 	// register the trace provider
// 	otel.SetTracerProvider(tp)

// 	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

// 	// register the global propagator
// 	return tp.Shutdown, nil

// }

// init zap logger from config
func initZapLogger(conf *config.ViperConfig) *zap.Logger {
	// writer
	writers := []io.Writer{}
	// add stdout
	writers = append(writers, os.Stdout)

	// add file
	multi := io.MultiWriter(writers...)

	// set log level
	var level zap.AtomicLevel
	switch conf.Observability.Logging.Level {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// init zap logger
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(multi),
		level,
	))

	return logger
}

// get zerolog logger
// func initZerologLogger(conf *config.ViperConfig) zerolog.Logger {
// 	// writer
// 	writers := []io.Writer{}
// 	// add stdout
// 	writers = append(writers, os.Stdout)

// 	switch conf.Observability.Logging.Level {
// 	case "debug":
// 		zerolog.SetGlobalLevel(zerolog.DebugLevel)
// 	case "info":
// 		zerolog.SetGlobalLevel(zerolog.InfoLevel)
// 	case "warn":
// 		zerolog.SetGlobalLevel(zerolog.WarnLevel)
// 	case "error":
// 		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
// 	case "fatal":
// 		zerolog.SetGlobalLevel(zerolog.FatalLevel)
// 	case "panic":
// 		zerolog.SetGlobalLevel(zerolog.PanicLevel)
// 	default:
// 		zerolog.SetGlobalLevel(zerolog.InfoLevel)
// 	}

// 	// add file
// 	multi := io.MultiWriter(writers...)

// 	return zerolog.New(multi).With().Timestamp().Logger()
// 	// ...
// }

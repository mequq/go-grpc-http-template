package main

import (
	"io"

	"os"

	"github.com/mequq/go-grpc-http-template/config"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	wireApp, err := wireApp(cfg, logger)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := wireApp.RunGRPC(); err != nil {
			panic(err)
		}
	}()

	if err := wireApp.RunHTTP(); err != nil {
		panic(err)
	}

}

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
func initZerologLogger(conf *config.ViperConfig) zerolog.Logger {
	// writer
	writers := []io.Writer{}
	// add stdout
	writers = append(writers, os.Stdout)

	switch conf.Observability.Logging.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// add file
	multi := io.MultiWriter(writers...)

	return zerolog.New(multi).With().Timestamp().Logger()
	// ...
}

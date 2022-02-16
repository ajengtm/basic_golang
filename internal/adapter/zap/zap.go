package zap

import (
	"context"
	"runtime"
	"strings"

	"github.com/opentracing/opentracing-go"
	stackdriver "github.com/tommy351/zap-stackdriver"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	return logger
}

func For(ctx context.Context) *zap.Logger {
	thisLogger := logger
	if span := opentracing.SpanFromContext(ctx); span != nil {
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			traceId := sc.TraceID().String()
			thisLogger = thisLogger.
				With(zap.String("trace_id", traceId))
		}
	}
	return thisLogger
}

func InitLogToFile() {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.OutputPaths = []string{"stdout", "log.txt"}
	zapConfig.ErrorOutputPaths = []string{"stdout", "log.txt"}
	logger, _ = zapConfig.Build()
}

func Init() {
	_, file, line, _ := runtime.Caller(1)
	slash := strings.LastIndex(file, "/")
	file = file[slash+1:]

	config := zap.NewProductionConfig()
	config.EncoderConfig = stackdriver.EncoderConfig

	logger, _ = config.Build(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return &stackdriver.Core{
			Core: core,
		}
	}), zap.Fields(
		stackdriver.LogServiceContext(&stackdriver.ServiceContext{
			Service: "basic-golang",
			Version: "1.0.0",
		}),
	),
		zap.Fields(
			stackdriver.LogReportLocation(&stackdriver.ReportLocation{
				FilePath:     file,
				LineNumber:   line,
				FunctionName: "",
			}),
		),
	)

}

package log

import (
	"context"
	"cros-rate-limit-service/configs"
	"cros-rate-limit-service/internal/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log LogWrapper

func init() {
	logger := NewLogger(zap.InfoLevel)
	Log = LogWrapper{logger: logger}
}

type LogWrapper struct {
	logger *zap.Logger
}

func NewLogger(level zapcore.Level) *zap.Logger {
	core := newCore(level)
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.PanicLevel), zap.Development(), zap.Fields())
}

func newCore(level zapcore.Level) zapcore.Core {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	// 自定义文件-行号输出项
	customerCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath())
	}

	// 自定义日志级别显示
	customerLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(level.CapitalString())
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    customerLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   customerCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		atomicLevel,
	)
}

func Debug(tag string, fields ...zap.Field) {
	Log.logger.Debug(tag, fields...)
}

func DebugF(ctx context.Context, tag string, fields ...zap.Field) {
	_trace := ctx.Value(configs.TraceKey).(*trace.Trace)
	Log.logger.Debug(tag,
		append(fields, zap.String("trx", _trace.Trx))...,
	)
}

func Info(tag string, fields ...zap.Field) {
	Log.logger.Info(tag, fields...)
}

func InfoF(ctx context.Context, tag string, fields ...zap.Field) {
	_trace := ctx.Value(configs.TraceKey).(*trace.Trace)
	Log.logger.Info(tag,
		append(fields, zap.String("trx", _trace.Trx))...,
	)
}

func Warn(tag string, fields ...zap.Field) {
	Log.logger.Warn(tag, fields...)
}

func WarnF(ctx context.Context, tag string, fields ...zap.Field) {
	_trace := ctx.Value(configs.TraceKey).(*trace.Trace)
	Log.logger.Warn(tag,
		append(fields, zap.String("trx", _trace.Trx))...,
	)
}

func Error(tag string, fields ...zap.Field) {
	Log.logger.Error(tag, fields...)
}

func ErrorF(ctx context.Context, tag string, fields ...zap.Field) {
	_trace := ctx.Value(configs.TraceKey).(*trace.Trace)
	Log.logger.Error(tag,
		append(fields, zap.String("trx", _trace.Trx))...,
	)
}

func Fatal(tag string, fields ...zap.Field) {
	Log.logger.Fatal(tag, fields...)
}

func FatalF(ctx context.Context, tag string, fields ...zap.Field) {
	_trace := ctx.Value(configs.TraceKey).(*trace.Trace)
	Log.logger.Fatal(tag,
		append(fields, zap.String("trx", _trace.Trx))...,
	)
}

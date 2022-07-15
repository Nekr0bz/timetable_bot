// Package log from https://github.com/bigwhite/experiments/blob/master/uber-zap-advanced-usage/
package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DebugLevel  Level = zap.DebugLevel  // -1 default level
	InfoLevel   Level = zap.InfoLevel   // 0,
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5

	timeLayout = "2006-01-02 15:04:05.000 MST"
)

// function variables for all field types
// in github.com/uber-go/zap/field.go

var (
	FSkip       = zap.Skip
	FBinary     = zap.Binary
	FBool       = zap.Bool
	FByteString = zap.ByteString
	FComplex128 = zap.Complex128
	FComplex64  = zap.Complex64
	FFloat64    = zap.Float64
	FFloat32    = zap.Float32
	FInt        = zap.Int
	FInt64      = zap.Int64
	FInt32      = zap.Int32
	FInt16      = zap.Int16
	FInt8       = zap.Int8
	FString     = zap.String
	FUint       = zap.Uint
	FUint64     = zap.Uint64
	FUint32     = zap.Uint32
	FUint16     = zap.Uint16
	FUint8      = zap.Uint8
	FUintptr    = zap.Uintptr
	FReflect    = zap.Reflect
	FNamespace  = zap.Namespace
	FStringer   = zap.Stringer
	FTime       = zap.Time
	FTimep      = zap.Timep
	FStack      = zap.Stack
	FStackSkip  = zap.StackSkip
	FDuration   = zap.Duration
	FDurationp  = zap.Durationp
	FAny        = zap.Any
	FError      = zap.Error

	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug
)

// ResetDefault not safe for concurrent use
func ResetDefault(l *Logger) {
	std = l
	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
	Debug = std.Debug
}

type Logger struct{ *zap.Logger }

var std = New()

func Default() *Logger {
	return std
}

// New create a new logger (not support log rotating).
func New() *Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeLayout))
	}

	logger, _ := cfg.Build()
	return &Logger{logger}
}

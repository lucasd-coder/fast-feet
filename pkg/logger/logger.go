package logger

import (
	"context"
	"os"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

var (
	l *Log
)

type Option struct {
	AppName string
	Level   string
}

type Log struct {
	opt Option
}

func NewLog(opt Option) *Log {
	if l == nil {
		l = &Log{opt: opt}
	}
	return l
}

func FromContext(ctx context.Context) Logger {
	log := l.GetLogger()

	logger := &logger{
		logger: log.WithContext(ctx),
	}

	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		logger.WithField("traceID", span.TraceID().String())
	}

	return logger
}

func (l *Log) GetLogger() *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logLevel, _ := logrus.ParseLevel(l.opt.Level)
	logrus.SetLevel(logLevel)

	return log.WithFields(logrus.Fields{
		"logName":  l.opt.AppName,
		"logIndex": "message",
	})
}

func (l *Log) GetGRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_logrus.UnaryServerInterceptor(l.GetLogger())
}

func (l *Log) GetGRPCStreamServerInterceptor() grpc.StreamServerInterceptor {
	return grpc_logrus.StreamServerInterceptor(l.GetLogger())
}

func (l *Log) GetGRPCUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return grpc_logrus.UnaryClientInterceptor(l.GetLogger())
}

func (l *Log) GetGRPCStreamClientInterceptor() grpc.StreamClientInterceptor {
	return grpc_logrus.StreamClientInterceptor(l.GetLogger())
}

type logger struct {
	logger *logrus.Entry
}

func (l *logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *logger) WithFields(keyValues map[string]interface{}) Logger {
	newEntry := l.logger.WithFields(convertToLogrusFields(keyValues))

	newLogger := &logger{
		logger: newEntry,
	}

	return newLogger
}

func (l *logger) WithField(key string, value interface{}) Logger {
	newEntry := l.logger.WithField(key, value)
	newLogger := &logger{
		logger: newEntry,
	}

	return newLogger
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}

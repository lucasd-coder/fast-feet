package authservice

import (
	"context"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"

	"github.com/lucasd-coder/fast-feet/business-service/config"
	"github.com/lucasd-coder/fast-feet/business-service/internal/shared"
	"github.com/lucasd-coder/fast-feet/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(_ context.Context, cfg *config.Config) (*grpc.ClientConn, error) {
	url := cfg.AuthService.URL
	maxRetry := cfg.AuthService.MaxRetry
	retryWaitTime := cfg.AuthService.AuthServiceRetryWaitTime
	retryMaxWaitTime := cfg.AuthService.AuthServiceRetryMaxWaitTime
	optlogger := shared.NewOptLogger(cfg)

	logger := logger.NewLogger(optlogger)

	reg := prometheus.NewRegistry()
	clMetrics := grpcprom.NewClientMetrics(
		grpcprom.WithClientHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	reg.MustRegister(clMetrics)
	exemplarFromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}
		return nil
	}

	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(retryMaxWaitTime)),
		grpc_retry.WithMax(maxRetry),
		grpc_retry.WithPerRetryTimeout(retryWaitTime),
		grpc_retry.WithCodes(codes.Unavailable, codes.DeadlineExceeded),
	}

	conn, err := grpc.Dial(url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			clMetrics.UnaryClientInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			grpc_retry.UnaryClientInterceptor(opts...),
			logger.GetLogUnaryClientInterceptor(),
		),
		grpc.WithChainStreamInterceptor(
			clMetrics.StreamClientInterceptor(grpcprom.WithExemplarFromContext(exemplarFromContext)),
			grpc_retry.StreamClientInterceptor(opts...),
			logger.GetLogStreamClientInterceptor(),
		),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

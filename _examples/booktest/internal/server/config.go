// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc).

package server

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

// Config represents the server configuration
type Config struct {
	ServiceName    string
	Port           int
	PrometheusPort int
	EnableCors     bool
	OtlpEndpoint   string

	Middlewares []HttpMiddlewareType
}

// PrometheusEnabled check configuration
func (c Config) PrometheusEnabled() bool {
	return c.PrometheusPort > 0
}

// TracingEnabled check configuration
func (c Config) TracingEnabled() bool {
	return c.OtlpEndpoint != ""
}

func (c Config) grpcInterceptors() []grpc.UnaryServerInterceptor {
	interceptors := make([]grpc.UnaryServerInterceptor, 0)
	interceptors = append(interceptors, logging.UnaryServerInterceptor(interceptorLogger(slog.Default()),
		logging.WithDisableLoggingFields("protocol", "grpc.component", "grpc.method_type")))
	interceptors = append(interceptors, errorMapper)
	interceptors = append(interceptors, recovery.UnaryServerInterceptor())

	return interceptors
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

package interceptors

import "go.uber.org/fx"

var Module = fx.Module(
	"grpc_interceptors",
	fx.Provide(
		NewLoggingInterceptor,
	),
)

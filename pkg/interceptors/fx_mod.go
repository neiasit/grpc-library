package interceptors

import "go.uber.org/fx"

var Module = fx.Module(
	"grpc_interceptors",
	fx.Provide(
		AsUnaryServerInterceptor(NewLoggingInterceptor),
	),
)

var ModuleWithAuth = fx.Provide(
	"grpc_interceptors_with_auth",
	Module,
	fx.Provide(
		AsUnaryServerInterceptor(NewAuthInterceptor),
	),
)

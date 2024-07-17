package core

import "go.uber.org/fx"

var Module = fx.Module(
	"grpc_core",
	fx.Provide(
		LoadConfig,
		NewGrpcServer,
	),
	fx.Invoke(RunGrpcServer),
)

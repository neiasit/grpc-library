package core

import (
	"github.com/neiasit/grpc-library/constants"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"grpc_core",
	fx.Provide(
		LoadConfig,
		fx.Annotate(
			NewGrpcServer,
			fx.ParamTags("", constants.UnaryServerInterceptorGroup),
		),
	),
	fx.Invoke(RunGrpcServer),
)

package grpc_library

import (
	"github.com/neiasit/grpc-library/core"
	"github.com/neiasit/grpc-library/interceptors"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"grpc_infrastructure",
	interceptors.Module,
	core.Module,
)

var ModuleWithAuth = fx.Module(
	"grpc_infrastructure",
	interceptors.ModuleWithAuth,
	core.Module,
)

package pkg

import (
	"github.com/neiasit/grpc-library/pkg/core"
	"github.com/neiasit/grpc-library/pkg/interceptors"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"grpc_infrastructure",
	interceptors.Module,
	core.Module,
)

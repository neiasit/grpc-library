package interceptors

import (
	"github.com/neiasit/grpc-library/pkg/constants"
	"go.uber.org/fx"
)

func AsUnaryServerInterceptor(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(constants.UnaryServerInterceptorGroup),
	)
}

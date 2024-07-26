package interceptors

import (
	"github.com/neiasit/grpc-library/constants"
	"go.uber.org/fx"
)

func AsUnaryServerInterceptor(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(constants.UnaryServerInterceptorGroup),
	)
}

package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	hello "yuumi-movie/api/helloworld/v1"
	userRouter "yuumi-movie/api/user/interface/v1"
	"yuumi-movie/internal/conf"
	"yuumi-movie/internal/pkg/middleware/auth"
	"yuumi-movie/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {
	skipRouters := map[string]struct{}{
		"/api.user.interface.v1.User/Register": {},
		"/api.user.interface.v1.User/Login":    {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, greeter *service.GreeterService, user *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(jwtc.Secret)).Match(NewWhiteListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	hello.RegisterGreeterHTTPServer(srv, greeter)
	userRouter.RegisterUserHTTPServer(srv, user)
	return srv
}

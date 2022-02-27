package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	realworldv1 "realworld/api/realworld/v1"
	"realworld/internal/conf"
	"realworld/internal/service"
	"realworld/pkg/err_encoder"
	"realworld/pkg/middleware"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, realworld *service.RealWorldService, l log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.Auth(),
		),
		//	改为自己的 error encoder
		http.ErrorEncoder(err_encoder.ErrorEncoder),
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
	realworldv1.RegisterRealWorldHTTPServer(srv, realworld)
	return srv
}

package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"os"
	"realworld/pkg/logger"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"realworld/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func initLogger() log.Logger {
	l := logger.New("info")
	logconf := conf.AllConf.LogConf
	rotationtime, err := time.ParseDuration(logconf.Rotationtime)
	if err != nil {
		panic("rotationtime in log config illegal")
	}
	maxage, err := time.ParseDuration(logconf.Maxage)
	if err != nil {
		panic("maxage in log config illegal")
	}
	writer, err := rotatelogs.New(logconf.Path, rotatelogs.WithRotationTime(rotationtime), rotatelogs.WithMaxAge(maxage))
	if err != nil {
		panic("init log fail")
	}
	l.SetOutput(writer)
	return log.With(l,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

func main() {
	flag.Parse()
	// @TODO 改为自定义的log组件
	// log with 中添加的是全局默认字段
	//logger := log.With(log.NewStdLogger(os.Stdout),
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//	"service.id", id,
	//	"service.name", Name,
	//	"service.version", Version,
	//	"trace_id", tracing.TraceID(),
	//	"span_id", tracing.SpanID(),
	//)
	conf.ReadAllConf(flagconf)
	logger := initLogger()
	app, cleanup, err := initApp(conf.AllConf.BootstrapConf.Server, conf.AllConf.BootstrapConf.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

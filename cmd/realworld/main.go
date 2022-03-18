package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
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

// Set global trace provider
func setTracerProvider() error {
	// Create the Jaeger exporter
	//exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	//if err != nil {
	//	return err
	//}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		//tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("env", "dev"),
		)),
		tracesdk.WithIDGenerator(&Generator{}),
	)
	otel.SetTracerProvider(tp)
	return nil
}

type Generator struct{}

func (g *Generator) NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID) {
	traceIDHex := fmt.Sprintf("%032x", "ZqH8H072nbn6km0tKO8wV861tpQNpc")
	traceID, _ := trace.TraceIDFromHex(traceIDHex)
	//gen.traceID++

	spanID := g.NewSpanID(ctx, traceID)
	fmt.Println(traceID, spanID)
	return traceID, spanID
	//var rngSeed uint64
	//_ = binary.Read(crand.Reader, binary.LittleEndian, &rngSeed)
	//randSource := rand.New(rand.NewSource(rngSeed))
	//
	//tid := trace.TraceID{}
	//randSource.Read(tid[:])
	//sid := trace.SpanID{}
	//randSource.Read(sid[:])
	//fmt.Println(tid, sid)
	//return tid, sid
}

// NewSpanID returns a non-zero span ID from a randomly-chosen sequence.
func (gen *Generator) NewSpanID(ctx context.Context, traceID trace.TraceID) trace.SpanID {
	spanIDHex := fmt.Sprintf("%016x", 2)
	spanID, _ := trace.SpanIDFromHex(spanIDHex)
	return spanID
}

func main() {
	flag.Parse()
	// 读取配置文件
	conf.ReadAllConf(flagconf)

	// 初始化tracer
	setTracerProvider()
	// 初始化日志
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

package tracing

import (
	"io"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func init() {
	os.Setenv("JAEGER_AGENT_HOST", "192.168.1.67")
}

func InitTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(err)
	}

	cfg.ServiceName = service
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LogSpans = true

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}

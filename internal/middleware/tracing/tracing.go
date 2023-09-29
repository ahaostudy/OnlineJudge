package tracing

import (
	"fmt"
	"io"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"

	cfg "main/config"
)

func InitTracer(service string) (opentracing.Tracer, io.Closer) {
	os.Setenv("JAEGER_AGENT_HOST", cfg.ConfJaeger.Host)
	os.Setenv("JAEGER_AGENT_PORT", fmt.Sprintf("%d", cfg.ConfJaeger.Port))

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

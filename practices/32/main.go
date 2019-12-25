package main

import (
	"io"
	"log"
	"time"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 32")
}

func main() {
	sConf := &jaegercfg.SamplerConfig{
		Type:  jaeger.SamplerTypeRateLimiting,
		Param: 1.0,
	}
	rConf := &jaegercfg.ReporterConfig{
		QueueSize:           128,
		BufferFlushInterval: 10 * time.Second,
		LocalAgentHostPort:  "localhost:6831",
		LogSpans:            true,
	}
	closer, errJaeger := InitJaeger("32", sConf, rConf)
	if errJaeger != nil {
		panic(errJaeger)
	}
	defer closer.Close()

	var sp opentracing.Span
	testCarrier := map[string]string{}
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.TextMap,
		opentracing.TextMapCarrier(testCarrier))
	if err != nil {
		log.Printf("fuck err: %+v", err)
		sp = opentracing.StartSpan("span a")
	} else {
		sp = opentracing.StartSpan(
			"span a",
			ext.RPCServerOption(wireContext))
	}
	sp.SetTag("yoyo", "haha")
	defer sp.Finish()
}

func InitJaeger(componentName string, samplerConf *jaegercfg.SamplerConfig, reporterConf *jaegercfg.ReporterConfig) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: componentName,
		Sampler:     samplerConf,
		Reporter:    reporterConf,
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return closer, nil
}

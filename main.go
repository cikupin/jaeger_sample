package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cikupin/jaeger_sample/internal"
	"github.com/go-chi/chi"
	"github.com/opentracing/opentracing-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
)

const (
	serviceName = "service testing app"
)

func registerTracer() (opentracing.Tracer, io.Closer) {
	cfg := &jaegerCfg.Configuration{
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	tracer, closer, err := cfg.New(serviceName)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func main() {
	// register jaeger
	tracer, closer := registerTracer()
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	r := chi.NewRouter()
	r.Get("/user/{user_id}", internal.HandlerDummy)

	log.Println("server is running using port 8080 ...")
	http.ListenAndServe(":8080", r)
}

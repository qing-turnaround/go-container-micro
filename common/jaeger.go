package common

import (
	"io"
	"time"

	"github.com/uber/jaeger-client-go"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// 创建链路追踪实例

func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	// io.Closer 用于关闭数据流
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1, // 1, 0表示真假
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  addr,
		},
	}
	return cfg.NewTracer()
}

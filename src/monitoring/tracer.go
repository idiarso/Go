package monitoring

import (
    "context"
    "github.com/opentracing/opentracing-go"
    "github.com/uber/jaeger-client-go"
    "github.com/uber/jaeger-client-go/config"
)

func InitTracer(serviceName string) (opentracing.Tracer, error) {
    cfg := &config.Configuration{
        ServiceName: serviceName,
        Sampler: &config.SamplerConfig{
            Type:  jaeger.SamplerTypeConst,
            Param: 1,
        },
        Reporter: &config.ReporterConfig{
            LogSpans: true,
        },
    }
    
    tracer, _, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
    if err != nil {
        return nil, err
    }
    
    opentracing.SetGlobalTracer(tracer)
    return tracer, nil
}

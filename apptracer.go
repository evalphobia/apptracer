package apptracer

import (
	"golang.org/x/net/context"

	"github.com/evalphobia/apptracer/platform"
)

// AppTracer has multiple tracing platform clients
// to send tracing activity to the platforms.
type AppTracer struct {
	Config
	ClientList []PlatformClient
}

// New returns initialized AppTracer.
func New(conf Config) *AppTracer {
	return &AppTracer{
		Config: conf,
	}
}

// AddClient adds PlatformClient into ClientList.
func (a *AppTracer) AddClient(c PlatformClient) {
	a.ClientList = append(a.ClientList, c)
}

// Trace creates *TraceWrapper from context.
// TraceWrapper contains platform clients and generates span data for the platforms.
// TraceWrapper is cached in the context and you can get same TraceWrapper
// when pass the same context.
func (a *AppTracer) Trace(ctx context.Context) *TraceWrapper {
	switch {
	case ctx == nil,
		a == nil,
		!a.Config.Enable:
		return &TraceWrapper{}
	}

	t, ok := ctx.Value(traceKey{}).(*TraceWrapper)
	if ok {
		return t
	}

	t = &TraceWrapper{
		traceList:     make([]platform.Trace, 0, len(a.ClientList)),
		defaultLabels: a.Config.DefaultLabels,
		serviceName:   a.Config.ServiceName,
		version:       a.Config.Version,
		environment:   a.Config.Environment,
	}
	for _, cli := range a.ClientList {
		tr, err := cli.NewTrace(ctx)
		if err == nil {
			t.traceList = append(t.traceList, tr)
		}
	}
	t.ctx = context.WithValue(ctx, traceKey{}, t)
	return t
}

type traceKey struct{} // used to get TraceWrapper from context.

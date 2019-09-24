package lightstep

import (
	"time"

	"github.com/lightstep/lightstep-tracer-go"
	"google.golang.org/grpc"
)

// Config contains token and settings for LightStep.
type Config struct {
	// required
	AccessToken string

	// optional
	ServiceName        string
	ReportTimeout      time.Duration
	ReportingPeriod    time.Duration
	MinReportingPeriod time.Duration
	ReconnectPeriod    time.Duration

	MaxBufferedSpans          int
	MaxLogKeyLen              int
	MaxLogValueLen            int
	MaxLogsPerSpan            int
	DropSpanLogs              bool
	MetaEventReportingEnabled bool

	UseGRPC                     bool
	DialOptions                 []grpc.DialOption
	GRPCMaxCallSendMsgSizeBytes int

	Collector    lightstep.Endpoint
	LightStepAPI lightstep.Endpoint
	Recorder     lightstep.SpanRecorder
	ConnFactory  lightstep.ConnectorFactory
}

// ToOption converts Config to lightstep-tracer-go.Options.
func (c Config) ToOption() lightstep.Options {
	opt := lightstep.Options{
		AccessToken:               c.AccessToken,
		ReportTimeout:             c.ReportTimeout,
		ReportingPeriod:           c.ReportingPeriod,
		MinReportingPeriod:        c.MinReportingPeriod,
		ReconnectPeriod:           c.ReconnectPeriod,
		MaxBufferedSpans:          c.MaxBufferedSpans,
		MaxLogKeyLen:              c.MaxLogKeyLen,
		MaxLogValueLen:            c.MaxLogValueLen,
		MaxLogsPerSpan:            c.MaxLogsPerSpan,
		DropSpanLogs:              c.DropSpanLogs,
		MetaEventReportingEnabled: c.MetaEventReportingEnabled,
		Collector:                 c.Collector,
		LightStepAPI:              c.LightStepAPI,
		Recorder:                  c.Recorder,
		ConnFactory:               c.ConnFactory,
		Tags:                      make(map[string]interface{}),
	}
	if c.ServiceName != "" {
		opt.Tags[lightstep.ComponentNameKey] = c.ServiceName
	}

	if c.UseGRPC {
		opt.UseGRPC = c.UseGRPC
		opt.DialOptions = c.DialOptions
		opt.GRPCMaxCallSendMsgSizeBytes = c.GRPCMaxCallSendMsgSizeBytes
	}

	return opt
}

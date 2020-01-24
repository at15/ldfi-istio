package ldfiadapter

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"istio.io/api/mixer/adapter/model/v1beta1"
	"istio.io/istio/mixer/template/metric"
)

// Server is basic server interface
type Server interface {
	Addr() string
	Close() error
	Run(shutdown chan error)
}

// LDFIAdapter supports metric template.
type LDFIAdapter struct {
	listener net.Listener
	server   *grpc.Server
}

var _ metric.HandleMetricServiceServer = &LDFIAdapter{}

// HandleMetric records metric entries
func (s *LDFIAdapter) HandleMetric(ctx context.Context, r *metric.HandleMetricRequest) (*v1beta1.ReportResult, error) {
	return nil, nil
}

// Addr returns the listening address of the server
func (s *LDFIAdapter) Addr() string {
	return s.listener.Addr().String()
}

// Run starts the server run
func (s *LDFIAdapter) Run(shutdown chan error) {
	shutdown <- s.server.Serve(s.listener)
}

// Close gracefully shuts down the server; used for testing
func (s *LDFIAdapter) Close() error {
	if s.server != nil {
		s.server.GracefulStop()
	}

	if s.listener != nil {
		_ = s.listener.Close()
	}

	return nil
}

// NewLDFIAdapter creates a new IBP adapter that listens at provided port.
func NewLDFIAdapter(addr string) (Server, error) {
	if addr == "" {
		addr = "0"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		return nil, fmt.Errorf("unable to listen on socket: %v", err)
	}
	s := &LDFIAdapter{
		listener: listener,
	}
	fmt.Printf("listening on \"%v\"\n", s.Addr())
	s.server = grpc.NewServer()
	metric.RegisterHandleMetricServiceServer(s.server, s)
	return s, nil
}

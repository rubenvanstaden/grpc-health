package grpc_test

import (
	"testing"

	rpc "github.com/rubenvanstaden/grpc-health/grpc"
)

type Server struct {
	*rpc.Server
}

func MustOpenServer(tb testing.TB) *Server {
	tb.Helper()

	s := &Server{
		Server: rpc.NewServer(),
	}

	s.Server.Addr = TEST_URL

	err := s.Open()
	if err != nil {
		tb.Fatal(err)
	}

	return s
}

func MustCloseServer(tb testing.TB, s *Server) {
	tb.Helper()
	s.Close()
}

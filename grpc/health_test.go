package grpc_test

import (
	"context"
	"testing"
	"time"

	"github.com/rubenvanstaden/grpc-health/core"
	"github.com/rubenvanstaden/grpc-health/grpc"
	"github.com/rubenvanstaden/grpc-health/test"
)

const (
	TEST_URL = "localhost:8080"
)

func TestIntegration_Health(t *testing.T) {

	m := MustOpenServer(t)
	defer MustCloseServer(t, m)

	flags := core.Flags{
		Addr:        TEST_URL,
		Service:     "",
		ConnTimeout: 5 * time.Second,
		RPCTimeout:  5 * time.Second,
		Verbose:     false,
	}

	rpc := grpc.NewClient(flags)
	defer rpc.Close()

	client := grpc.NewHealthClient(rpc)

	t.Run("GRPC", func(t *testing.T) {

		code := client.Check(context.Background(), flags.Service)
		test.Equals(t, 0, code)

	})
}

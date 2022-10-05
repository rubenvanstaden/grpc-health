package core

import (
	"time"
)

const (
	StatusSuccess = iota + 0
	// StatusInvalidArguments indicates specified invalid arguments.
	StatusInvalidArguments
	// StatusConnectionFailure indicates connection failed.
	StatusConnectionFailure
	// StatusRPCFailure indicates rpc failed.
	StatusRPCFailure
	// StatusUnhealthy indicates rpc succeeded but indicates unhealthy service.
	StatusUnhealthy
)

type Flags struct {
	Addr        string
	Service     string
	ConnTimeout time.Duration
	RPCTimeout  time.Duration
	Verbose     bool
}

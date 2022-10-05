package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/rubenvanstaden/grpc-health/core"
	"github.com/rubenvanstaden/grpc-health/grpc"
)

var (
	flAddr        string
	flService     string
	flConnTimeout time.Duration
	flRPCTimeout  time.Duration
	flVerbose     bool
)

func init() {

	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	log.SetFlags(0)

	flagSet.StringVar(&flAddr, "addr", "", "(required) tcp host:port to connect")
	flagSet.StringVar(&flService, "service", "", "service name to check (default: \"\")")
	flagSet.DurationVar(&flConnTimeout, "connect-timeout", 5*time.Second, "timeout for establishing connection")
	flagSet.DurationVar(&flRPCTimeout, "rpc-timeout", 5*time.Second, "timeout for health check rpc")
	flagSet.BoolVar(&flVerbose, "v", false, "verbose logs")

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		os.Exit(core.StatusInvalidArguments)
	}

	argError := func(s string, v ...interface{}) {
		os.Exit(core.StatusInvalidArguments)
	}

	if flAddr == "" {
		argError("-addr not specified")
	}

	if flConnTimeout <= 0 {
		argError("-connect-timeout must be greater than zero (specified: %v)", flConnTimeout)
	}

	if flRPCTimeout <= 0 {
		argError("-rpc-timeout must be greater than zero (specified: %v)", flRPCTimeout)
	}

	if flVerbose {
		log.Printf("\n")
		log.Printf("parsed options:")
		log.Printf("> addr=%s conn_timeout=%v rpc_timeout=%v", flAddr, flConnTimeout, flRPCTimeout)
		log.Printf("\n")
	}
}

func main() {

	code := core.StatusSuccess
	defer func() {
		os.Exit(code)
	}()

	flags := core.Flags{
		Addr:        flAddr,
		Service:     flService,
		ConnTimeout: flConnTimeout,
		RPCTimeout:  flRPCTimeout,
		Verbose:     flVerbose,
	}

	client := grpc.NewClient(flags)
	defer client.Close()

	health := grpc.NewHealthClient(client)

	rpcStart := time.Now()

	ctx, rpcCancel := context.WithTimeout(context.Background(), flRPCTimeout)
	defer rpcCancel()

	code = health.Check(ctx, flService)

	rpcDuration := time.Since(rpcStart)

	if flags.Verbose {
		log.Printf("time elapsed: connect=%v rpc=%v", client.ConnDuration, rpcDuration)
	}

	log.Printf("status code: %v", code)
}

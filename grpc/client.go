package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rubenvanstaden/grpc-health/core"
)

type Client struct {
	flags        core.Flags
	StatusCode   int
	ctx          context.Context
	conn         *grpc.ClientConn
	cancel       context.CancelFunc
	connTimeout  time.Duration
	ConnDuration time.Duration
}

func NewClient(flags core.Flags) *Client {
	return &Client{
		flags: flags,
	}
}

func (self *Client) Open() *grpc.ClientConn {

	connStart := time.Now()

	self.ctx, self.cancel = context.WithTimeout(context.Background(), self.flags.ConnTimeout)

	conn, err := grpc.DialContext(self.ctx, self.flags.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
		log.Fatalf("ConnTimeout: %s", self.flags.ConnTimeout)
		if err == context.DeadlineExceeded {
			log.Printf("timeout: failed to connect service %q within %v", self.flags.Addr, self.flags.ConnTimeout)
		} else {
			log.Printf("error: failed to connect service at %q: %+v", self.flags.Addr, err)
		}
		self.StatusCode = core.StatusConnectionFailure
	}

	self.conn = conn

	self.ConnDuration = time.Since(connStart)

	if self.flags.Verbose {
		log.Printf("dial connection: established (took %v)", self.ConnDuration)
	}

	return conn
}

func (self *Client) Close() {

	self.conn.Close()
	self.cancel()

	if self.flags.Verbose {
		log.Println("dial connection: closed")
	}
}

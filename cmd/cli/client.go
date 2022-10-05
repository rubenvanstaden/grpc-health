package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	URL          string
	ctx          context.Context
	conn         *grpc.ClientConn
	connStart    time.Time
	cancel       context.CancelFunc
	ConnDuration time.Duration
}

func NewClient(url string) *Client {
	return &Client{URL: url}
}

func (self *Client) Open() *grpc.ClientConn {

	self.connStart = time.Now()

	self.ctx, self.cancel = context.WithTimeout(context.Background(), flConnTimeout)

	conn, err := grpc.DialContext(self.ctx, self.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
		if err == context.DeadlineExceeded {
			log.Printf("timeout: failed to connect service %q within %v", flAddr, flConnTimeout)
		} else {
			log.Printf("error: failed to connect service at %q: %+v", flAddr, err)
		}
		statusCode = StatusConnectionFailure
	}

	self.conn = conn

	self.ConnDuration = time.Since(self.connStart)

	if flVerbose {
		log.Printf("dial connection: established (took %v)", self.ConnDuration)
	}

	return conn
}

func (self *Client) Close() {

	self.conn.Close()
	self.cancel()

	if flVerbose {
		log.Println("dial connection: closed")
	}
}

package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/rubenvanstaden/grpc-health/grpc/pb"
)

type Server struct {
	server *grpc.Server
	Addr   string
}

func NewServer() *Server {
	return &Server{
		server: grpc.NewServer(),
	}
}

func (self *Server) Open() error {

	self.registerServers()

	err := self.serve()
	if err != nil {
		return err
	}

	return nil
}

func (self *Server) Close() {
	self.server.Stop()
}

func (self *Server) registerServers() {
	service := NewHealthService()
	pb.RegisterHealthServer(self.server, service)
	reflection.Register(self.server)
}

func (self *Server) serve() error {

	listener, err := net.Listen("tcp", self.Addr)
	if err != nil {
		return err
	}

	go func() {
		err := self.server.Serve(listener)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	return nil
}

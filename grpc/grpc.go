package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	Listener   net.Listener
	GRPCServer *grpc.Server
}

func NewServer(port string) (*Server, error) {
	list, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return &Server{}, err
	}

	s := grpc.NewServer()

	return &Server{
		Listener:   list,
		GRPCServer: s,
	}, nil
}

func (s *Server) Run() error {
	fmt.Println("Server GRPC running on port: " + s.Listener.Addr().String())
	return s.GRPCServer.Serve(s.Listener)
}

func NewGRPCSocket(address string) (*grpc.ClientConn, error) {
	socket, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	return socket, err
}

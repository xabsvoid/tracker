package grpc

import (
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type (
	Service interface {
		RegisterService(grpc.ServiceRegistrar)
	}

	Server struct {
		network  string
		address  string
		services []Service
	}
)

func NewServer(
	network string,
	address string,
	services ...Service,
) *Server {
	return &Server{
		network:  network,
		address:  address,
		services: services,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	for _, service := range s.services {
		service.RegisterService(server)
	}

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	var shutdown sync.WaitGroup
	shutdown.Add(1)
	go func() {
		<-shutdownCh
		server.GracefulStop()
		shutdown.Done()
	}()

	err = server.Serve(listener)
	if err != nil {
		return err
	}

	shutdown.Wait()

	return nil
}

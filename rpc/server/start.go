package server

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/rs/zerolog/log"
)

type StartResult struct {
	Address string
	Error   error
}

type StopFunc func()

func Start(srv *rpc.Server, port int) (stopFn StopFunc, addr string, err error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, "", fmt.Errorf("failed to listen on port %d: %w", port, err)
	}

	go srv.Accept(listener)

	stopFn = func() {
		if err = listener.Close(); err != nil {
			log.Printf("failed to close listener on port %d", port)
		}
	}

	return stopFn, listener.Addr().String(), nil
}

package server

import (
	"fmt"
	"net"

	"github.com/fiuskyws/thoth/src/manager"
	"github.com/fiuskyws/thoth/src/proto"
	"github.com/fiuskyws/thoth/src/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type (
	// GRPC stores all info/data necessary to establish a
	// gRPC server.
	GRPC struct {
		listener net.Listener
		mgr      *manager.Manager
		repo     *repository.GRPCRepo
		srv      *grpc.Server
		opts     []grpc.ServerOption
		port     uint
	}
)

// NewGRPC returns a new GRPC connector.
func NewGRPC(mgr *manager.Manager, port uint) API {
	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)

	repo := repository.NewGRPCRepo(mgr)

	return &GRPC{
		mgr:  mgr,
		opts: opts,
		repo: repo,
		srv:  srv,
		port: port,
	}
}

// Start - starts the gRPC server.
func (g *GRPC) Start() error {
	var err error
	g.listener, err = net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", g.port))
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}
	proto.RegisterThothServer(g.srv, proto.ThothServer(g.repo))
	if err := g.srv.Serve(g.listener); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}

// Close - closes the gRPC server.
func (g *GRPC) Close() error {
	g.srv.GracefulStop()
	if err := g.listener.Close(); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}

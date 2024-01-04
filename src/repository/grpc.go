package repository

import (
	"context"

	"github.com/fiuskyws/thoth/src/manager"
	"github.com/fiuskyws/thoth/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCRepo struct {
	proto.UnimplementedThothServer
	mgr *manager.Manager
}

func NewGRPCRepo(mgr *manager.Manager) *GRPCRepo {
	return &GRPCRepo{
		mgr: mgr,
	}
}
func (g *GRPCRepo) CreateTable(_ context.Context, req *proto.CreateTableRequest) (*proto.CreateTableResponse, error) {
	if err := g.mgr.CreateTable(req.Name); err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}
	return &proto.CreateTableResponse{}, nil
}

func (g *GRPCRepo) GetTables(_ context.Context, _ *proto.GetTablesRequest) (*proto.GetTablesResponse, error) {
	tables := g.mgr.GetTables()
	return &proto.GetTablesResponse{
		Tables: tables,
	}, nil
}

func (g *GRPCRepo) Get(_ context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	val, err := g.mgr.Get(req.Table, req.Key)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}
	return &proto.GetResponse{
		Value: val,
	}, nil
}

func (g *GRPCRepo) Set(_ context.Context, req *proto.SetRequest) (*proto.SetResponse, error) {
	err := g.mgr.Set(req.Table, req.Key, req.Value)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}
	return &proto.SetResponse{}, nil
}

func (g *GRPCRepo) Delete(_ context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	err := g.mgr.Delete(req.Table, req.Key)
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}
	return &proto.DeleteResponse{}, nil
}

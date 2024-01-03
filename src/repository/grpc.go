package repository

import (
	"context"

	"github.com/fiuskyws/thoth/src/manager"
	"github.com/fiuskyws/thoth/src/proto"
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
func (g *GRPCRepo) CreateTable(context.Context, *proto.CreateTableRequest) (*proto.CreateTableResponse, error) {
	return nil, nil
}

func (g *GRPCRepo) Delete(context.Context, *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return nil, nil
}

func (g *GRPCRepo) Get(context.Context, *proto.GetRequest) (*proto.GetResponse, error) {
	return nil, nil
}

func (g *GRPCRepo) GetTables(context.Context, *proto.GetTablesRequest) (*proto.GetTablesResponse, error) {
	return nil, nil
}

func (g *GRPCRepo) Set(context.Context, *proto.SetRequest) (*proto.SetResponse, error) {
	return nil, nil
}

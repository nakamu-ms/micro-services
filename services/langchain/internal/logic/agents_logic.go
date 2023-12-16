package logic

import (
	"context"

	"github.com/nakamu-ms/micro-services/services/langchain/internal/svc"
	"github.com/nakamu-ms/micro-services/services/langchain/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAgentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgentsLogic {
	return &AgentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AgentsLogic) Agents(in *pb.AgentsRequest) (*pb.AgentsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.AgentsResponse{}, nil
}

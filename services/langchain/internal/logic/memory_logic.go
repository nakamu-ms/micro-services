package logic

import (
	"context"

	"github.com/nakamu-ms/micro-services/services/langchain/internal/svc"
	"github.com/nakamu-ms/micro-services/services/langchain/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemoryLogic {
	return &MemoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemoryLogic) Memory(in *pb.MemoryRequest) (*pb.MemoryResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.MemoryResponse{}, nil
}

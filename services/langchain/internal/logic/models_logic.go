package logic

import (
	"context"

	"github.com/nakamu-ms/micro-services/services/langchain/internal/svc"
	"github.com/nakamu-ms/micro-services/services/langchain/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModelsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModelsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModelsLogic {
	return &ModelsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModelsLogic) Models(in *pb.ModelsRequest) (*pb.ModelsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ModelsResponse{}, nil
}

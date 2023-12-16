package logic

import (
	"context"

	"github.com/nakamu-ms/micro-services/services/langchain/internal/svc"
	"github.com/nakamu-ms/micro-services/services/langchain/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChainsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChainsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChainsLogic {
	return &ChainsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChainsLogic) Chains(in *pb.ChainsRequest) (*pb.ChainsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ChainsResponse{}, nil
}

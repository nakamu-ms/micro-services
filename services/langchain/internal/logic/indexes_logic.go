package logic

import (
	"context"

	"github.com/nakamu-ms/micro-services/services/langchain/internal/svc"
	"github.com/nakamu-ms/micro-services/services/langchain/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIndexesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexesLogic {
	return &IndexesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IndexesLogic) Indexes(in *pb.IndexesRequest) (*pb.IndexesResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.IndexesResponse{}, nil
}

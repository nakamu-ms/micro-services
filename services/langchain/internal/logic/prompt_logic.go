package logic

import (
	"context"

	"github.com/nakamu-ms/micro-services/services/langchain/internal/svc"
	"github.com/nakamu-ms/micro-services/services/langchain/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PromptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPromptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PromptLogic {
	return &PromptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PromptLogic) Prompt(in *pb.PromptRequest) (*pb.PromptResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.PromptResponse{}, nil
}

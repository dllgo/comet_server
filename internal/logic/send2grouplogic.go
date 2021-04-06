package logic

import (
	"context"

	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type Send2groupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSend2groupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Send2groupLogic {
	return &Send2groupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Send2groupLogic) Send2Group(in *comet_server.Input) (*comet_server.Output, error) {
	// todo: add your logic here and delete this line

	return &comet_server.Output{}, nil
}

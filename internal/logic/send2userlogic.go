package logic

import (
	"context"

	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type Send2userLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSend2userLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Send2userLogic {
	return &Send2userLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Send2userLogic) Send2User(in *comet_server.Request) (*comet_server.Response, error) {
	// todo: add your logic here and delete this line

	return &comet_server.Response{}, nil
}

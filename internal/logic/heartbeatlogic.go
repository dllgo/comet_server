package logic

import (
	"context"

	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type HeartbeatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHeartbeatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HeartbeatLogic {
	return &HeartbeatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HeartbeatLogic) Heartbeat(in *comet_server.Input) (*comet_server.Output, error) {
	// todo: add your logic here and delete this line

	return &comet_server.Output{
		Type:      comet_server.PackageType_PT_HEARTBEAT,
		RequestId: in.RequestId,
		}, nil
}

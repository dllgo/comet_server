package logic

import (
	"context"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type SyncLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncLogic {
	return &SyncLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncLogic) Sync(in *comet_server.Input, c comet.IConn) (*comet_server.Output, error) {
	// todo: add your logic here and delete this line

	return &comet_server.Output{
		Type:      comet_server.PackageType_PT_SYNC,
		RequestId: in.RequestId,
		}, nil
}

package logic

import (
	"context"

	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type BusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessLogic {
	return &BusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BusinessLogic) Business(in *comet_server.Input) (*comet_server.Output, error) {
	// todo: add your logic here and delete this line

	return &comet_server.Output{
		Type:      comet_server.PackageType_PT_HANDSHAKE,
		RequestId: in.RequestId,
		}, nil
}

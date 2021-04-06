package logic

import (
	"context"
	"log"
	"strconv"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type HandshakeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHandshakeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandshakeLogic {
	return &HandshakeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HandshakeLogic) Handshake(in *comet_server.Input, c comet.IConn) (*comet_server.Output, error) {
	// todo: add your logic here and delete this line
	var msg = "ok"
    var signinput comet_server.SignInInput
	if err := signinput.XXX_Unmarshal(in.Data); err != nil {
		log.Println("[HandshakeLogic] Handshake 解码错误", err)
		msg = "[HandshakeLogic] Handshake 解码错误"
		return nil,err
	}
	//校验设备

	//保存登录 
	uid := strconv.FormatInt(signinput.UserId,10) 
	ctx := context.WithValue(context.Background(), "uid", uid)
	c.SetContext(ctx)
	comet.ConnectHandlerIns().C(uid,c)
	//
	return &comet_server.Output{
		Type:      comet_server.PackageType_PT_HANDSHAKE,
		RequestId: in.RequestId,
		Code: 200,
		Message: msg,
	}, nil
}

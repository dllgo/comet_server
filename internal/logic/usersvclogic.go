package logic

import (
	"context"
	"log"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UsersvcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsersvcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersvcLogic {
	return &UsersvcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
用户响应
*/
func (l *UsersvcLogic) Userresponse(in *comet_server.Request, code int64, message string, data []byte) *comet_server.Response {
	return &comet_server.Response{
		I:       in.I,
		ID:      in.ID,
		Service: in.Service,
		Path:    in.Path,
		Version: in.Version,
		ST:      in.ST,
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func (l *UsersvcLogic) Usersvc(in *comet_server.Request, c comet.IConn) (*comet_server.Response, error) {
	// todo: add your logic here and delete this line
	log.Println("[UsersvcLogic] 用户服务=====>", in.Path)
	switch in.Path {
	case "login":
		// 用户登录
		return l.Userloignsvc(in, c)
	case "online":
		// 上线
		return l.Useronlinesvc(in, c)
	case "offline":
		//离线
		return l.Userofflinesvc(in, c)
	default:
		log.Println("[UsersvcLogic] 用户服务=====>", "handler path other")
	}
	return l.Userresponse(in, 200, "", nil), nil
}
func (l *UsersvcLogic) Userloignsvc(in *comet_server.Request, c comet.IConn) (*comet_server.Response, error) {
	// todo: add your logic here and delete this line
	log.Println("[UsersvcLogic] 用户服务=====>login")
	// var msg = "ok"
	// var signinput comet_server.SignInInput
	// if err := signinput.XXX_Unmarshal(in.Data); err != nil {
	// 	log.Println("[HandshakeLogic] Handshake 解码错误", err)
	// 	msg = "[HandshakeLogic] Handshake 解码错误"
	// 	return nil,err
	// }
	// //校验设备

	// //保存登录
	// uid := strconv.FormatInt(signinput.UserId,10)
	// ctx := context.WithValue(context.Background(), "uid", uid)
	// c.SetContext(ctx)
	// comet.ConnectHandlerIns().C(uid,c)
	// //

	return l.Userresponse(in, 200, "", nil), nil
}
func (l *UsersvcLogic) Useronlinesvc(in *comet_server.Request, c comet.IConn) (*comet_server.Response, error) {
	// todo: add your logic here and delete this line
	log.Println("[UsersvcLogic] 用户服务=====>online")
	return l.Userresponse(in, 200, "", nil), nil
}
func (l *UsersvcLogic) Userofflinesvc(in *comet_server.Request, c comet.IConn) (*comet_server.Response, error) {
	// todo: add your logic here and delete this line
	log.Println("[UsersvcLogic] 用户服务=====>offline")
	return l.Userresponse(in, 200, "", nil), nil
}

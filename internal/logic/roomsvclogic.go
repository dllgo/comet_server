package logic

import (
	"context"
	"log"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoomsvcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoomsvcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoomsvcLogic {
	return &RoomsvcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
房间响应
*/
func (l *RoomsvcLogic) Roomresponse(in *comet_server.Request, code int64, message string, data []byte) *comet_server.Response {
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
func (l *RoomsvcLogic) Roomsvc(in *comet_server.Request, c comet.IConn) (*comet_server.Response, error) {
	// todo: add your logic here and delete this line
	log.Println("[RoomsvcLogic] 房间服务=====>", in.Path)
	switch in.Path {
	case "":
	default:
		log.Println("[RoomsvcLogic] 房间服务=====>", "handler path other")
	}
	return l.Roomresponse(in, 200, "", nil), nil
}

package logic

import (
	"context"
	"log"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GroupsvcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupsvcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupsvcLogic {
	return &GroupsvcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
群响应
*/
func (l *GroupsvcLogic) Groupresponse(in *comet_server.Request, code int64, message string, data []byte) *comet_server.Response {
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
func (l *GroupsvcLogic) Groupsvc(in *comet_server.Request, c comet.IConn) (*comet_server.Response, error) {
	// todo: add your logic here and delete this line
	log.Println("[GroupsvcLogic] 群服务=====>", in.Path)
	switch in.Path {
	case "":
	default:
		log.Println("[GroupsvcLogic] 群服务=====>", "handler path other")
	}
	return l.Groupresponse(in, 200, "", nil), nil
}

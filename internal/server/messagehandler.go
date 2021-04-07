package server

import (
	"context"
	"log"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/logic"
	"github.com/dllgo/comet_server/internal/svc"
)

func NewMessageHandler(svcCtx *svc.ServiceContext) *MessageHandler {
	return &MessageHandler{
		svcCtx: svcCtx,
	}
}

//tcp event
type MessageHandler struct {
	comet.IEvent
	svcCtx *svc.ServiceContext
}

func (e *MessageHandler) OnClosed(c comet.IConn, err error) {
	log.Println("[EventHandler OnClosed] client: " + c.RemoteAddr().String())
}
func (e *MessageHandler) GetUid(c comet.IConn) string {
	if c != nil && c.Context() != nil {
		ctx := c.Context().(context.Context)
		if ctx != nil {
			cid := ctx.Value("uid").(string)
			return cid
		}
		return ""

	}
	return ""
}
func (e *MessageHandler) OnMessage(frame []byte, c comet.IConn) {

	var input comet_server.Request
	if err := input.XXX_Unmarshal(frame); err != nil {
		log.Println("[MessageHandler] handle 解码错误", err)
		return
	}
	log.Println("[MessageHandler] handle 接收到", e.GetUid(c), input.Service, input.Path, "的消息")
	switch input.Service {
	case "usersvc":
		// 用户服务
		usvc := logic.NewUsersvcLogic(context.Background(), e.svcCtx)
		usvc.Usersvc(&input, c)
	case "groupsvc":
		// 群服务
		gsvc := logic.NewGroupsvcLogic(context.Background(), e.svcCtx)
		gsvc.Groupsvc(&input, c)
	case "roomsvc":
		//房间服务
		rsvc := logic.NewRoomsvcLogic(context.Background(), e.svcCtx)
		rsvc.Roomsvc(&input, c)
	case "pushsvc":
		// 推送服务
		psvc := logic.NewPushsvcLogic(context.Background(), e.svcCtx)
		psvc.Pushsvc(&input, c)
	case "imsvc":
		// im服务
		isvc := logic.NewImsvcLogic(context.Background(), e.svcCtx)
		isvc.Imsvc(&input, c)
	default:
		log.Println("[MessageHandler] handle 接收到", "handler switch other")
	}
}

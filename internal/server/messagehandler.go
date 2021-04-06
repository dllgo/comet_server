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
func (e *MessageHandler)OnClosed(c comet.IConn, err error){
	log.Println("[EventHandler OnClosed] client: " + c.RemoteAddr().String() )
}
func (e *MessageHandler)GetUid(c comet.IConn) string{
	if c!=nil && c.Context() != nil {
		ctx := c.Context().(context.Context)
		if ctx !=nil {
			cid := ctx.Value("uid").(string)
			return cid	
		}
		return ""
		
	}
	return ""
}
func (e *MessageHandler)OnMessage(frame []byte, c comet.IConn){
	
	var input comet_server.Input
	if err := input.XXX_Unmarshal(frame); err != nil {
		log.Println("[MessageHandler] handle 解码错误", err)
		return
	} 
	log.Println("[TcpHandler] handle 接收到", e.GetUid(c),input.Type, "的消息")
	switch input.Type {
	case comet_server.PackageType_PT_HANDSHAKE: 
		// 握手
		e.handshake(c, input)
	case comet_server.PackageType_PT_SYNC: 
		// 同步
		e.sync(c, input)
	case comet_server.PackageType_PT_HEARTBEAT:
		//心跳
		e.heartbeat(c, input)
	case comet_server.PackageType_PT_MESSAGE: 
		// 消息
		e.message(c, input)
	case comet_server.PackageType_PT_ACK:
		// 回执
		e.ack(c, input)
	default:
		log.Println("[TcpHandler] handle 接收到", "handler switch other")
	}
}
// handshake 握手
func (e *MessageHandler) handshake(c comet.IConn, input comet_server.Input) {
	log.Println("[MessageHandler] handle", "handshake 握手")
	l := logic.NewHandshakeLogic(context.Background(), e.svcCtx)
	output ,err := l.Handshake(&input,c)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(output.GetData())	
	}
}

// ack 回执
func (e *MessageHandler) ack(c comet.IConn, input comet_server.Input) {
	log.Println("[MessageHandler] handle", "ack")
	l := logic.NewAckLogic(context.Background(), e.svcCtx)
	output ,err := l.Ack(&input,c)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(output.GetData())	
	}
}

// sync 同步
func (e *MessageHandler) sync(c comet.IConn, input comet_server.Input) {
	log.Println("[MessageHandler] handle", "sync")
	l := logic.NewSyncLogic(context.Background(), e.svcCtx)
	output ,err := l.Sync(&input,c)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(output.GetData())	
	}
}

// Heartbeat 心跳
func (e *MessageHandler) heartbeat(c comet.IConn, input comet_server.Input) {
	log.Println("[MessageHandler] handle", "Heartbeat 心跳")
	l := logic.NewHeartbeatLogic(context.Background(), e.svcCtx)
	output ,err := l.Heartbeat(&input,c)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(output.GetData())	
	}
}

// Message 消息处理
func (e *MessageHandler) message(c comet.IConn, input comet_server.Input) {
	log.Println("[MessageHandler] handle", "Business")
	 
	l := logic.NewBusinessLogic(context.Background(), e.svcCtx)
	output ,err := l.Business(&input,c)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(output.GetData())	
	}
}
 
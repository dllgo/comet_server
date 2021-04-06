package svc

import (
	"context"
	"log"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
)

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{ 
	}
}

//tcp event
type MessageHandler struct {
	comet.IEvent
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
	 
	log.Println("[TcpHandler] handle 接收到", e.GetUid(c), "的消息")
	var input comet_server.Input
	if err := input.XXX_Unmarshal(frame); err != nil {
		log.Println("[MessageHandler] handle 解码错误", err)
		return
	}
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
	log.Println("[TcpHandler] handle", "handshake 握手")
	frame ,err := e.output(input,comet_server.PackageType_PT_HANDSHAKE)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(frame)	
	}
}

// ack 回执
func (e *MessageHandler) ack(c comet.IConn, input comet_server.Input) {
	log.Println("[TcpHandler] handle", "ack")

	frame ,err := e.output(input,comet_server.PackageType_PT_ACK)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(frame)	
	}
	
}

// sync 同步
func (e *MessageHandler) sync(c comet.IConn, input comet_server.Input) {
	log.Println("[TcpHandler] handle", "sync")
	frame ,err := e.output(input,comet_server.PackageType_PT_SYNC)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(frame)	
	}
}

// Heartbeat 心跳
func (e *MessageHandler) heartbeat(c comet.IConn, input comet_server.Input) {
	log.Println("[TcpHandler] handle", "Heartbeat 心跳")
	
	frame ,err := e.output(input,comet_server.PackageType_PT_HEARTBEAT)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(frame)	
	}
}

// Message 消息处理
func (e *MessageHandler) message(c comet.IConn, input comet_server.Input) {
	log.Println("[TcpHandler] handle", "Heartbeat 心跳")
	 
	frame ,err := e.output(input,comet_server.PackageType_PT_HEARTBEAT)
	if err != nil {
		return
	}
	if c!=nil {
		c.AsyncWrite(frame)	
	}
}

//out

func (e *MessageHandler) output(input comet_server.Input,ptype comet_server.PackageType) ([]byte,error){
	var output = comet_server.Output{
		Type:      ptype,
		RequestId: input.RequestId,
	}

	return output.GetData(),nil
}
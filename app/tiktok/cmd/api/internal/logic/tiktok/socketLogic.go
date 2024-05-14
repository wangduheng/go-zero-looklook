package tiktok

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"
	"looklook/app/tiktok/cmd/rpc/tiktok"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type SocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SocketLogic {
	return &SocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var conn *websocket.Conn

func (l *SocketLogic) ResultsHandler(req *types.ReceiveReq) error {
	liveData, err := l.svcCtx.TiktokRpc.FindLiveOne(l.ctx, &tiktok.FindLiveReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	reqBytes, err := json.Marshal(liveData)
	if err != nil {
		return err
	}
	logx.Info("write socket data: ", string(reqBytes))
	if conn == nil {
		logx.Error("没有上线用户")
		return nil
	}
	if err := conn.WriteMessage(websocket.TextMessage, reqBytes); err != nil {
		return errors.Wrap(err, "Error sending data to WebSocket client")
	}
	return nil
}

func (l *SocketLogic) ServeWs(w http.ResponseWriter, r *http.Request) {
	var err error
	conn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Error(err)
		return
	}
	// defer conn.Close()
	go handleConnectionClose(conn)
	go sendHeartbeat(conn, 30*time.Second)

}
func handleConnectionClose(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			conn = nil
			// 在连接关闭时执行相应的清理操作
			return
		}
	}
}
func sendHeartbeat(conn *websocket.Conn, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		// 发送心跳消息
		err := conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			logx.Error("Error sending heartbeat:", err)
			return
		}
	}
}

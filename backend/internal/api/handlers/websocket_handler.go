package handlers

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"

	"master-guide-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketHandler WebSocket 处理器
type WebSocketHandler struct {
	websocketMgr *utils.WebSocketManager
}

// NewWebSocketHandler 创建 WebSocket 处理器
func NewWebSocketHandler(websocketMgr *utils.WebSocketManager) *WebSocketHandler {
	return &WebSocketHandler{
		websocketMgr: websocketMgr,
	}
}

// WebSocket 连接升级器
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源，生产环境应该限制
		return true
	},
}

// HandleWebSocket WebSocket 连接处理
// @Summary WebSocket 连接
// @Description 建立 WebSocket 连接，支持实时消息和在线状态
// @Tags WebSocket
// @Accept json
// @Produce json
// @Success 101 {string} string "Switching Protocols"
// @Router /ws [get]
func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	// 生成客户端ID
	clientID := h.generateClientID()

	// 创建 WebSocket 客户端
	client := &utils.WebSocketClient{
		ID:       clientID,
		UserID:   "", // 将在认证后设置
		Conn:     conn,
		Send:     make(chan []byte, 256),
		Manager:  h.websocketMgr,
		IsOnline: false,
	}

	// 注册客户端
	h.websocketMgr.RegisterClient(client)

	// 启动读写协程
	go client.WritePump()
	go client.ReadPump()
}

// generateClientID 生成客户端ID
func (h *WebSocketHandler) generateClientID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

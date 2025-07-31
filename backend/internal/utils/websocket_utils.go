package utils

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"master-guide-backend/internal/model"

	"github.com/gorilla/websocket"
)

// WebSocketManager WebSocket 连接管理器
type WebSocketManager struct {
	clients    map[string]*WebSocketClient
	broadcast  chan *model.WebSocketEvent
	register   chan *WebSocketClient
	unregister chan *WebSocketClient
	mutex      sync.RWMutex
}

// WebSocketClient WebSocket 客户端连接
type WebSocketClient struct {
	ID       string
	UserID   string
	Conn     *websocket.Conn
	Send     chan []byte
	Manager  *WebSocketManager
	IsOnline bool
	LastSeen time.Time
}

// NewWebSocketManager 创建 WebSocket 管理器
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:    make(map[string]*WebSocketClient),
		broadcast:  make(chan *model.WebSocketEvent),
		register:   make(chan *WebSocketClient),
		unregister: make(chan *WebSocketClient),
	}
}

// Start 启动 WebSocket 管理器
func (manager *WebSocketManager) Start() {
	for {
		select {
		case client := <-manager.register:
			manager.mutex.Lock()
			manager.clients[client.ID] = client
			manager.mutex.Unlock()
			log.Printf("Client registered: %s", client.ID)

		case client := <-manager.unregister:
			manager.mutex.Lock()
			if _, ok := manager.clients[client.ID]; ok {
				delete(manager.clients, client.ID)
				close(client.Send)
			}
			manager.mutex.Unlock()
			log.Printf("Client unregistered: %s", client.ID)

		case event := <-manager.broadcast:
			manager.broadcastEvent(event)
		}
	}
}

// RegisterClient 注册客户端
func (manager *WebSocketManager) RegisterClient(client *WebSocketClient) {
	manager.register <- client
}

// UnregisterClient 注销客户端
func (manager *WebSocketManager) UnregisterClient(client *WebSocketClient) {
	manager.unregister <- client
}

// BroadcastEvent 广播事件
func (manager *WebSocketManager) BroadcastEvent(event *model.WebSocketEvent) {
	manager.broadcast <- event
}

// SendToUser 发送消息给指定用户
func (manager *WebSocketManager) SendToUser(userID string, event *model.WebSocketEvent) {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	for _, client := range manager.clients {
		if client.UserID == userID && client.IsOnline {
			message, err := json.Marshal(event)
			if err != nil {
				log.Printf("Error marshaling event: %v", err)
				continue
			}
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(manager.clients, client.ID)
			}
		}
	}
}

// GetOnlineUsers 获取在线用户列表
func (manager *WebSocketManager) GetOnlineUsers() []*model.OnlineUser {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	var onlineUsers []*model.OnlineUser
	for _, client := range manager.clients {
		if client.IsOnline {
			onlineUsers = append(onlineUsers, &model.OnlineUser{
				UserID:   client.UserID,
				IsOnline: client.IsOnline,
				LastSeen: client.LastSeen,
			})
		}
	}
	return onlineUsers
}

// IsUserOnline 检查用户是否在线
func (manager *WebSocketManager) IsUserOnline(userID string) bool {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	for _, client := range manager.clients {
		if client.UserID == userID && client.IsOnline {
			return true
		}
	}
	return false
}

// broadcastEvent 广播事件到所有客户端
func (manager *WebSocketManager) broadcastEvent(event *model.WebSocketEvent) {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling event: %v", err)
		return
	}

	for _, client := range manager.clients {
		if client.IsOnline {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(manager.clients, client.ID)
			}
		}
	}
}

// ReadPump 读取客户端消息
func (c *WebSocketClient) ReadPump() {
	defer func() {
		c.Manager.UnregisterClient(c)
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512)
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}

		// 处理接收到的消息
		c.handleMessage(message)
	}
}

// WritePump 向客户端发送消息
func (c *WebSocketClient) WritePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// handleMessage 处理接收到的消息
func (c *WebSocketClient) handleMessage(message []byte) {
	var event model.WebSocketEvent
	if err := json.Unmarshal(message, &event); err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return
	}

	switch event.Event {
	case "authenticate":
		c.handleAuthenticate(event.Data)
	case "message":
		c.handleChatMessage(event.Data)
	default:
		log.Printf("Unknown event type: %s", event.Event)
	}
}

// handleAuthenticate 处理认证消息
func (c *WebSocketClient) handleAuthenticate(data interface{}) {
	// 这里应该验证 JWT token
	// 简化处理，实际应该调用认证服务
	c.IsOnline = true
	c.LastSeen = time.Now()

	response := model.WebSocketEvent{
		Event: "authenticate",
		Data: model.AuthenticateResponse{
			Success: true,
			Message: "认证成功",
			UserID:  c.UserID,
		},
	}

	message, _ := json.Marshal(response)
	c.Send <- message
}

// handleChatMessage 处理聊天消息
func (c *WebSocketClient) handleChatMessage(data interface{}) {
	// 这里应该保存消息到数据库并广播给其他用户
	// 简化处理，直接广播
	event := model.WebSocketEvent{
		Event: "message",
		Data:  data,
	}

	c.Manager.BroadcastEvent(&event)
}

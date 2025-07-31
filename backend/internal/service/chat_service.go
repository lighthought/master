package service

import (
	"context"
	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
	"master-guide-backend/internal/utils"
)

// ChatService 聊天服务接口
type ChatService interface {
	GetOnlineUsers(ctx context.Context) (*model.OnlineUsersResponse, error)
	GetChatMessages(ctx context.Context, req *model.ChatMessagesRequest) (*model.ChatMessagesResponse, error)
	SendMessage(ctx context.Context, message *model.ChatMessage) error
}

type chatService struct {
	chatRepo     repository.ChatRepository
	websocketMgr *utils.WebSocketManager
}

func NewChatService(chatRepo repository.ChatRepository, websocketMgr *utils.WebSocketManager) ChatService {
	return &chatService{
		chatRepo:     chatRepo,
		websocketMgr: websocketMgr,
	}
}

func (s *chatService) GetOnlineUsers(ctx context.Context) (*model.OnlineUsersResponse, error) {
	// 从 WebSocket 管理器获取在线用户
	onlineUsers := s.websocketMgr.GetOnlineUsers()

	// 补充用户信息
	for _, user := range onlineUsers {
		if user.Name == "" {
			// 从数据库获取用户信息
			userProfile, err := s.chatRepo.GetUserProfile(ctx, user.UserID)
			if err == nil {
				user.Name = userProfile.Name
			}
		}
	}

	return &model.OnlineUsersResponse{
		OnlineUsers: onlineUsers,
	}, nil
}

func (s *chatService) GetChatMessages(ctx context.Context, req *model.ChatMessagesRequest) (*model.ChatMessagesResponse, error) {
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 50
	}

	// 获取聊天记录
	messages, total, err := s.chatRepo.GetChatMessages(ctx, req.TargetID, req.CircleID, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 计算分页信息
	totalPages := (int(total) + req.PageSize - 1) / req.PageSize

	pagination := &model.PaginationResponse{
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
		Data:       messages,
	}

	return &model.ChatMessagesResponse{
		Messages:   messages,
		Pagination: pagination,
	}, nil
}

func (s *chatService) SendMessage(ctx context.Context, message *model.ChatMessage) error {
	// 保存消息到数据库
	err := s.chatRepo.SaveMessage(ctx, message)
	if err != nil {
		return err
	}

	// 通过 WebSocket 广播消息
	event := &model.WebSocketEvent{
		Event: "message",
		Data:  message,
	}

	// 如果是私聊消息，只发送给目标用户
	if message.ToUserID != "" {
		s.websocketMgr.SendToUser(message.ToUserID, event)
	} else {
		// 如果是圈子消息，广播给所有在线用户
		s.websocketMgr.BroadcastEvent(event)
	}

	return nil
}

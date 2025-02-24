package yuque

import (
	"context"
	"net/http"
	"time"
)

type userService struct {
	client *Client
}

// Hello 心跳
func (s *userService) Hello(ctx context.Context, opts ...RequestOption) (*HelloResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "hello", nil, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(HelloResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

type HelloResponse struct {
	Message string `json:"message,omitempty"`
}

// User 获取当前 Token 的用户详情
func (s *userService) User(ctx context.Context, opts ...RequestOption) (*UserResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "user", nil, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(UserResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

type UserResponse struct {
	ID               int       `json:"id,omitempty"`
	Type             string    `json:"type,omitempty"`
	Login            string    `json:"login,omitempty"`
	Name             string    `json:"name,omitempty"`
	AvatarURL        string    `json:"avatar_url,omitempty"`
	BooksCount       int       `json:"books_count,omitempty"`
	PublicBooksCount int       `json:"public_books_count,omitempty"`
	FollowersCount   int       `json:"followers_count,omitempty"`
	FollowingCount   int       `json:"following_count,omitempty"`
	Public           int       `json:"public,omitempty"`
	Description      string    `json:"description,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	WorkID           string    `json:"work_id,omitempty"`
}

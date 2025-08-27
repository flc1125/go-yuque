package yuque

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type repoService struct {
	client *Client
}

// Book 知识库/Repository
type Book struct {
	ID               int        `json:"id"`
	Type             BookType   `json:"type"`
	Slug             string     `json:"slug"`
	Name             string     `json:"name"`
	UserID           int        `json:"user_id"`
	Description      string     `json:"description"`
	TocYML           string     `json:"toc_yml"`
	CreatorID        int        `json:"creator_id"`
	Public           AccessType `json:"public"`
	ItemsCount       int        `json:"items_count"`
	LikesCount       int        `json:"likes_count"`
	WatchesCount     int        `json:"watches_count"`
	ContentUpdatedAt time.Time  `json:"content_updated_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	Namespace        string     `json:"namespace"`
	User             *User      `json:"user"`
}

// GetRepos 获取知识库列表
func (s *repoService) GetRepos(ctx context.Context, ownerLogin string, request *GetReposRequest, opts ...RequestOption) ([]*Book, *Response, error) {
	var path string
	if request != nil && request.IsGroup {
		path = fmt.Sprintf("groups/%s/repos", ownerLogin)
	} else {
		path = fmt.Sprintf("users/%s/repos", ownerLogin)
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, request, opts)
	if err != nil {
		return nil, nil, err
	}

	var repos []*Book
	resp, err := s.client.Do(req, &repos)
	if err != nil {
		return nil, resp, err
	}

	return repos, resp, nil
}

// GetReposRequest 获取知识库列表请求
type GetReposRequest struct {
	Offset  *int      `url:"offset,omitempty"` // 偏移量 [分页参数]
	Limit   *int      `url:"limit,omitempty"`  // 每页数量 [分页参数]
	Type    *BookType `url:"type,omitempty"`   // 类型 [筛选条件]
	IsGroup bool      `url:"-"`                // 是否为团队 (内部使用)
}

// CreateRepo 创建知识库
func (s *repoService) CreateRepo(ctx context.Context, ownerLogin string, request *CreateRepoRequest, opts ...RequestOption) (*Book, *Response, error) {
	var path string
	if request != nil && request.IsGroup {
		path = fmt.Sprintf("groups/%s/repos", ownerLogin)
	} else {
		path = fmt.Sprintf("users/%s/repos", ownerLogin)
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, request, opts)
	if err != nil {
		return nil, nil, err
	}

	var repo Book
	resp, err := s.client.Do(req, &repo)
	if err != nil {
		return nil, resp, err
	}

	return &repo, resp, nil
}

// CreateRepoRequest 创建知识库请求
type CreateRepoRequest struct {
	Name            string     `json:"name"`                      // 名称 (必填)
	Slug            string     `json:"slug"`                      // 路径 (必填)
	Description     *string    `json:"description,omitempty"`     // 简介
	Public          AccessType `json:"public"`                    // 公开性
	EnhancedPrivacy *bool      `json:"enhancedPrivacy,omitempty"` // 增强私密性
	IsGroup         bool       `json:"-"`                         // 是否为团队 (内部使用)
}

// GetRepo 获取知识库详情
func (s *repoService) GetRepo(ctx context.Context, ownerLogin, repoSlug string, opts ...RequestOption) (*Book, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("repos/%s/%s", ownerLogin, repoSlug), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var repo Book
	resp, err := s.client.Do(req, &repo)
	if err != nil {
		return nil, resp, err
	}

	return &repo, resp, nil
}

// GetRepoByID 根据ID获取知识库详情
func (s *repoService) GetRepoByID(ctx context.Context, repoID int, opts ...RequestOption) (*Book, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("repos/%d", repoID), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var repo Book
	resp, err := s.client.Do(req, &repo)
	if err != nil {
		return nil, resp, err
	}

	return &repo, resp, nil
}

// UpdateRepo 更新知识库
func (s *repoService) UpdateRepo(ctx context.Context, ownerLogin, repoSlug string, request *UpdateRepoRequest, opts ...RequestOption) (*Book, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("repos/%s/%s", ownerLogin, repoSlug), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var repo Book
	resp, err := s.client.Do(req, &repo)
	if err != nil {
		return nil, resp, err
	}

	return &repo, resp, nil
}

// UpdateRepoByID 根据ID更新知识库
func (s *repoService) UpdateRepoByID(ctx context.Context, repoID int, request *UpdateRepoRequest, opts ...RequestOption) (*Book, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("repos/%d", repoID), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var repo Book
	resp, err := s.client.Do(req, &repo)
	if err != nil {
		return nil, resp, err
	}

	return &repo, resp, nil
}

// UpdateRepoRequest 更新知识库请求
type UpdateRepoRequest struct {
	Name        *string    `json:"name,omitempty"`        // 名称
	Slug        *string    `json:"slug,omitempty"`        // 路径
	Description *string    `json:"description,omitempty"` // 简介
	Public      AccessType `json:"public"`                // 公开性
	TOC         *string    `json:"toc,omitempty"`         // 目录
}

// DeleteRepo 删除知识库
func (s *repoService) DeleteRepo(ctx context.Context, ownerLogin, repoSlug string, opts ...RequestOption) (*Book, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("repos/%s/%s", ownerLogin, repoSlug), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var repo Book
	resp, err := s.client.Do(req, &repo)
	if err != nil {
		return nil, resp, err
	}

	return &repo, resp, nil
}

// DeleteRepoByID 根据ID删除知识库
func (s *repoService) DeleteRepoByID(ctx context.Context, repoID int, opts ...RequestOption) (*Book, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("repos/%d", repoID), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var repo Book
	resp, err := s.client.Do(req, &repo)
	if err != nil {
		return nil, resp, err
	}

	return &repo, resp, nil
}

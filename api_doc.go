package yuque

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type docService struct {
	client *Client
}

type Doc struct {
	ID               int       `json:"id,omitempty"`
	Type             DocType   `json:"type,omitempty"`
	Slug             string    `json:"slug,omitempty"`
	Title            string    `json:"title,omitempty"`
	Description      string    `json:"description,omitempty"`
	Cover            string    `json:"cover,omitempty"`
	UserID           int       `json:"user_id,omitempty"`
	BookID           int       `json:"book_id,omitempty"`
	LastEditorID     int       `json:"last_editor_id,omitempty"`
	Public           int       `json:"public,omitempty"`
	Status           int       `json:"status,omitempty"`
	LikesCount       int       `json:"likes_count,omitempty"`
	ReadCount        int       `json:"read_count,omitempty"`
	CommentsCount    int       `json:"comments_count,omitempty"`
	WordCount        int       `json:"word_count,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	ContentUpdatedAt time.Time `json:"content_updated_at,omitempty"`
	PublishedAt      time.Time `json:"published_at,omitempty"`
	FirstPublishedAt time.Time `json:"first_published_at,omitempty"`
	User             *User     `json:"user,omitempty"`
	LastEditor       *User     `json:"last_editor,omitempty"`
	Hits             int       `json:"hits,omitempty"`
}

func (s *docService) GetDocs(ctx context.Context, bookID any, request *GetDocsRequest, opts ...RequestOption) (*GetDocsResponse, *Response, error) {
	bid, err := parseID(bookID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("repos/%s/docs", bid), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var docs []*Doc
	resp, err := s.client.Do(req, &docs)
	if err != nil {
		return nil, resp, err
	}

	var total int
	if meta := resp.meta(); meta != nil {
		total = meta.Total
	}

	return &GetDocsResponse{
		Total: total,
		Docs:  docs,
	}, resp, nil
}

type GetDocsRequest struct {
	// 偏移量 [分页参数]
	Offset *int `url:"offset,omitempty"`

	// 每页数量 [分页参数]
	Limit *int `url:"limit,omitempty"`

	// 获取的额外字段, 多个字段以逗号分隔
	//
	// 注意: 每页数量超过 100 本字段会失效
	//
	// 支持的字段有:
	//
	// hits: 文档阅读数
	// tags: 标签
	// latest_version_id: 最新已发版本 ID
	OptionalProperties *string `url:"optional_properties,omitempty"`
}

type GetDocsResponse struct {
	Total int    `json:"total,omitempty"`
	Docs  []*Doc `json:"docs,omitempty"`
}

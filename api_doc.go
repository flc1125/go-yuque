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
	ID               int        `json:"id,omitempty"`
	Type             DocType    `json:"type,omitempty"`
	Slug             string     `json:"slug,omitempty"`
	Title            string     `json:"title,omitempty"`
	Description      string     `json:"description,omitempty"`
	Cover            string     `json:"cover,omitempty"`
	UserID           int        `json:"user_id,omitempty"`
	BookID           int        `json:"book_id,omitempty"`
	LastEditorID     int        `json:"last_editor_id,omitempty"`
	Public           AccessType `json:"public,omitempty"`
	Status           int        `json:"status,omitempty"`
	LikesCount       *int       `json:"likes_count,omitempty"`
	ReadCount        int        `json:"read_count,omitempty"`
	CommentsCount    *int       `json:"comments_count,omitempty"`
	WordCount        int        `json:"word_count,omitempty"`
	CreatedAt        time.Time  `json:"created_at,omitempty"`
	UpdatedAt        time.Time  `json:"updated_at,omitempty"`
	ContentUpdatedAt time.Time  `json:"content_updated_at,omitempty"`
	PublishedAt      time.Time  `json:"published_at,omitempty"`
	FirstPublishedAt time.Time  `json:"first_published_at,omitempty"`
	User             *User      `json:"user,omitempty"`
	LastEditor       *User      `json:"last_editor,omitempty"`
	Hits             int        `json:"hits,omitempty"`

	// 以下字段是创建文档、获取文档详情等时才有的字段
	Format          *DocFormat `json:"format,omitempty"`
	BodyDraft       *string    `json:"body_draft,omitempty"`
	Body            *string    `json:"body,omitempty"`
	BodySheet       *string    `json:"body_sheet,omitempty"`
	BodyTable       *string    `json:"body_table,omitempty"`
	BodyHTML        *string    `json:"body_html,omitempty"`
	BodyLake        *string    `json:"body_lake,omitempty"`
	Book            *Book      `json:"book,omitempty"`
	Creator         *User      `json:"creator,omitempty"`
	Tags            []DocTag   `json:"tags,omitempty"`
	LatestVersionID int        `json:"latest_version_id,omitempty"`
}

type DocTag struct {
	ID        int       `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	DocID     int       `json:"doc_id,omitempty"`
	BookID    int       `json:"book_id,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetDocs 获取知识库下的文档列表
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

// CreateDoc 创建文档
//
// 注意: 创建文档后不会自动添加到目录，需要调用"知识库目录更新接口"更新到目录中
func (s *docService) CreateDoc(ctx context.Context, bookID any, request *CreateDocRequest, opts ...RequestOption) (*Doc, *Response, error) {
	bid, err := parseID(bookID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("repos/%s/docs", bid), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var doc Doc
	resp, err := s.client.Do(req, &doc)
	if err != nil {
		return nil, resp, err
	}

	return &doc, resp, nil
}

// GetDoc 获取文档详情
//
// bookID: 知识库 ID 或 命名空间(group_login/book_slug)
// docID: 文档 ID 或 slug
func (s *docService) GetDoc(ctx context.Context, bookID any, docID any, opts ...RequestOption) (*Doc, *Response, error) {
	bid, err := parseID(bookID)
	if err != nil {
		return nil, nil, err
	}

	did, err := parseID(docID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("repos/%s/docs/%s", bid, did), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var doc Doc
	resp, err := s.client.Do(req, &doc)
	if err != nil {
		return nil, resp, err
	}

	return &doc, resp, nil
}

type CreateDocRequest struct {
	Slug   *string     `json:"slug,omitempty"`   // 路径
	Title  *string     `json:"title,omitempty"`  // 标题，默认：无标题
	Public *AccessType `json:"public,omitempty"` // 公开性 (0:私密, 1:公开, 2:企业内公开) 不填则继承知识库的公开性
	Format *DocFormat  `json:"format,omitempty"` // 内容格式 (markdown:Markdown 格式, html:HTML 标准格式, lake:语雀 Lake 格式) 不填则默认为 markdown
	Body   *string     `json:"body,omitempty"`   // 正文内容
}

// 获取文档详情
// 更新文档
// 删除文档
// 获取文档历史版本列表
// 获取文档历史版本详情

// GetTOCs 获取目录
func (s *docService) GetTOCs(ctx context.Context, bookID any, opts ...RequestOption) ([]*TOC, *Response, error) {
	bid, err := parseID(bookID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("repos/%s/toc", bid), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var rawTocs []*rawTOC
	resp, err := s.client.Do(req, &rawTocs)
	if err != nil {
		return nil, resp, err
	}

	tocs := make([]*TOC, len(rawTocs))
	for i, rawTOC := range rawTocs {
		toc := rawTOC.TOC

		// 兼容 id/doc_id 为 string 的情况，其实都是空值
		toc.ID = rawTocParseID(rawTOC.ID)
		toc.DocID = rawTocParseID(rawTOC.DocID)

		tocs[i] = &toc
	}

	return tocs, resp, nil
}

type rawTOC struct {
	TOC
	ID    any `json:"id,omitempty"`     // Deprecated 文档 ID
	DocID any `json:"doc_id,omitempty"` // 文档 ID
}

func rawTocParseID(id any) int {
	switch v := id.(type) {
	case int, int8, int16, int32, int64:
		return int(v.(int64))
	case float64, float32:
		return int(v.(float64))
	case string:
		return 0
	default:
		return 0
	}
}

type TOCType string

const (
	TOCTypeDoc   TOCType = "DOC"
	TOCTypeLink  TOCType = "LINK"
	TOCTypeTitle TOCType = "TITLE"
)

type TOC struct {
	UUID        string  `json:"uuid,omitempty"`         // 节点唯一 ID
	Type        TOCType `json:"type,omitempty"`         // Enum: "DOC" "LINK" "TITLE" 节点类型 (DOC:文档, LINK:外链, TITLE:分组)
	Title       string  `json:"title,omitempty"`        // 节点名称
	URL         string  `json:"url,omitempty"`          // 节点 URL
	Slug        string  `json:"slug,omitempty"`         // Deprecated 节点 URL
	ID          int     `json:"id,omitempty"`           // Deprecated 文档 ID
	DocID       int     `json:"doc_id,omitempty"`       // 文档 ID
	Level       int     `json:"level,omitempty"`        // 节点层级
	Depth       int     `json:"depth,omitempty"`        // Deprecated 节点层级
	OpenWindow  int     `json:"open_window,omitempty"`  // 是否在新窗口打开 (0:当前页打开, 1:新窗口打开)
	Visible     int     `json:"visible,omitempty"`      // 是否可见 (0:不可见, 1:可见)
	PrevUUID    string  `json:"prev_uuid,omitempty"`    // 同级前一个节点 uuid
	SiblingUUID string  `json:"sibling_uuid,omitempty"` // 同级后一个节点 uuid
	ChildUUID   string  `json:"child_uuid,omitempty"`   // 子级第一个节点 uuid
	ParentUUID  string  `json:"parent_uuid,omitempty"`  // 父级节点 uuid
}

// 更新目录

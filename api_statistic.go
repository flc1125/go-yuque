package yuque

import (
	"context"
	"fmt"
	"net/http"
)

type statisticService struct {
	client *Client
}

// GetGroupStatistics 团队.汇总统计数据
func (s *statisticService) GetGroupStatistics(ctx context.Context, login any, opts ...RequestOption) (*GetGroupStatisticsResponse, *Response, error) {
	lid, err := parseID(login)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("groups/%s/statistics", lid), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var response GetGroupStatisticsResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return &response, resp, nil
}

type GetGroupStatisticsResponse struct {
	BizDate                string  `json:"bizdate,omitempty"`
	UserID                 int     `json:"user_id,omitempty"`
	OrganizationID         int     `json:"organization_id,omitempty"`
	MemberCount            int     `json:"member_count,omitempty"`             // 成员数
	CollaboratorCount      int     `json:"collaborator_count,omitempty"`       // 协作者数
	DayReadCount           int     `json:"day_read_count,omitempty"`           // 当日阅读量
	DayWriteCount          int     `json:"day_write_count,omitempty"`          // 当日编辑次数
	WriteCount             int     `json:"write_count,omitempty"`              // 编辑次数
	ReadCount              int     `json:"read_count,omitempty"`               // 阅读量
	ReadCount30            int     `json:"read_count_30,omitempty"`            // 阅读量 (30天)
	ReadCount365           int     `json:"read_count_365,omitempty"`           // 阅读量 (一年)
	CommentCount           int     `json:"comment_count,omitempty"`            // 评论量
	CommentCount30         int     `json:"comment_count_30,omitempty"`         // 评论量 (30天)
	CommentCount365        int     `json:"comment_count_365,omitempty"`        // 评论量 (一年)
	LikeCount              int     `json:"like_count,omitempty"`               // 点赞量
	LikeCount30            int     `json:"like_count_30,omitempty"`            // 点赞量 (30天)
	LikeCount365           int     `json:"like_count_365,omitempty"`           // 点赞量 (一年)
	FollowCount            int     `json:"follow_count,omitempty"`             // 关注量
	CollectCount           int     `json:"collect_count,omitempty"`            // 收藏量
	DocCount               int     `json:"doc_count,omitempty"`                // 文档数量
	SheetCount             int     `json:"sheet_count,omitempty"`              // 表格数量
	BoardCount             int     `json:"board_count,omitempty"`              // 画板数量
	ShowCount              int     `json:"show_count,omitempty"`               // 演示文稿数量
	ResourceCount          int     `json:"resource_count,omitempty"`           // 资源数量
	ArtBoardCount          int     `json:"artboard_count,omitempty"`           // 图集数量
	AttachmentCount        int     `json:"attachment_count,omitempty"`         // 附件数量
	BookCount              int     `json:"book_count,omitempty"`               // 知识库总数量
	PublicBookCount        int     `json:"public_book_count,omitempty"`        // 公开知识库数量
	PrivateBookCount       int     `json:"private_book_count,omitempty"`       // 私密知识库数量
	BookBookCount          int     `json:"book_book_count,omitempty"`          // 文档知识库数量
	BookResourceCount      int     `json:"book_resource_count,omitempty"`      // 资源知识库数量
	BookDesignCount        int     `json:"book_design_count,omitempty"`        // 图片知识库数量
	BookThreadCount        int     `json:"book_thread_count,omitempty"`        // 话题知识库数量
	DataUsage              int     `json:"data_usage,omitempty"`               // 流量使用量
	GrainsCount            int     `json:"grains_count,omitempty"`             // 当前稻谷数
	GrainsCountSum         int     `json:"grains_count_sum,omitempty"`         // 累计获得稻谷数
	GrainsCountConsume     int     `json:"grains_count_consume,omitempty"`     // 已消耗稻谷数
	InteractionPeopleCount int     `json:"interaction_people_count,omitempty"` // 知识交流人数
	ContentCount           int     `json:"content_count,omitempty"`            // 知识财富数
	CollaborationCount     int     `json:"collaboration_count,omitempty"`      // 知识协同次数
	WorkingHours           int     `json:"working_hours,omitempty"`            // 协同提效时长 (小时)
	Baike                  float64 `json:"baike,omitempty"`                    // 百科全书卷数
	TableCount             int     `json:"table_count,omitempty"`
}

// GetGroupMemberStatistics 团队.成员统计数据
func (s *statisticService) GetGroupMemberStatistics(ctx context.Context, login string, request *GetGroupMemberStatisticsRequest, opts ...RequestOption) (*GetGroupMemberStatisticsResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("groups/%s/statistics/members", login), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response GetGroupMemberStatisticsResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return &response, resp, nil
}

// GetGroupMemberStatisticsRequest 团队成员统计请求
type GetGroupMemberStatisticsRequest struct {
	Name      *string    `url:"name,omitempty"`      // 成员名 [过滤条件]
	Range     *TimeRange `url:"range,omitempty"`     // 时间范围 [过滤条件]
	Page      *int       `url:"page,omitempty"`      // 页码
	Limit     *int       `url:"limit,omitempty"`     // 分页数量
	SortField *string    `url:"sortField,omitempty"` // 排序字段
	SortOrder *SortOrder `url:"sortOrder,omitempty"` // 排序方向
}

// GetGroupMemberStatisticsResponse 团队成员统计响应
type GetGroupMemberStatisticsResponse struct {
	Members []*MemberStatistics `json:"members,omitempty"` // 成员统计数据
	Total   int                 `json:"total,omitempty"`   // 总数量
}

// MemberStatistics 成员统计数据
type MemberStatistics struct {
	BizDate          string `json:"bizdate,omitempty"`             // 统计日期 (YYYYMMDD)
	UserID           string `json:"user_id,omitempty"`             // 成员 ID
	GroupID          string `json:"group_id,omitempty"`            // 团队 ID
	OrganizationID   string `json:"organization_id,omitempty"`     // 空间 ID
	WriteCount       string `json:"write_count,omitempty"`         // 编辑次数
	WriteCount30     string `json:"write_count_30,omitempty"`      // 编辑次数 (30天)
	WriteCount365    string `json:"write_count_365,omitempty"`     // 编辑次数 (一年)
	WriteDocCount    string `json:"write_doc_count,omitempty"`     // 编辑文档数
	WriteDocCount30  string `json:"write_doc_count_30,omitempty"`  // 编辑文档数 (30天)
	WriteDocCount365 string `json:"write_doc_count_365,omitempty"` // 编辑文档数 (一年)
	ReadCount        string `json:"read_count,omitempty"`          // 阅读量
	ReadCount30      string `json:"read_count_30,omitempty"`       // 阅读量 (30天)
	ReadCount365     string `json:"read_count_365,omitempty"`      // 阅读量 (一年)
	LikeCount        string `json:"like_count,omitempty"`          // 点赞量
	LikeCount30      string `json:"like_count_30,omitempty"`       // 点赞量 (30天)
	LikeCount365     string `json:"like_count_365,omitempty"`      // 点赞量 (一年)
	User             string `json:"user,omitempty"`                // 用户信息
}

// GetGroupBookStatistics 团队.知识库统计数据
func (s *statisticService) GetGroupBookStatistics(ctx context.Context, login string, request *GetGroupBookStatisticsRequest, opts ...RequestOption) (*GetGroupBookStatisticsResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("groups/%s/statistics/books", login), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response GetGroupBookStatisticsResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return &response, resp, nil
}

// GetGroupBookStatisticsRequest 团队知识库统计请求
type GetGroupBookStatisticsRequest struct {
	Name      *string    `url:"name,omitempty"`      // 知识库名 [过滤条件]
	Range     *TimeRange `url:"range,omitempty"`     // 时间范围 [过滤条件]
	Page      *int       `url:"page,omitempty"`      // 页码
	Limit     *int       `url:"limit,omitempty"`     // 分页数量
	SortField *string    `url:"sortField,omitempty"` // 排序字段
	SortOrder *SortOrder `url:"sortOrder,omitempty"` // 排序方向
}

// GetGroupBookStatisticsResponse 团队知识库统计响应
type GetGroupBookStatisticsResponse struct {
	Books []*BookStatistics `json:"books,omitempty"` // 知识库统计数据
	Total int               `json:"total,omitempty"` // 总数量
}

// BookStatistics 知识库统计数据
type BookStatistics struct {
	BizDate                string `json:"bizdate,omitempty"`                  // 统计日期 (YYYYMMDD)
	BookID                 string `json:"book_id,omitempty"`                  // 知识库 ID
	Slug                   string `json:"slug,omitempty"`                     // 知识库 slug
	Name                   string `json:"name,omitempty"`                     // 知识库名称
	Type                   string `json:"type,omitempty"`                     // 知识库类型
	IsPublic               string `json:"is_public,omitempty"`                // 是否公开
	ContentUpdatedAtMs     string `json:"content_updated_at_ms,omitempty"`    // 最近更新时间
	UserID                 string `json:"user_id,omitempty"`                  // 知识库归属 ID
	OrganizationID         string `json:"organization_id,omitempty"`          // 知识库归属空间 ID
	DayReadCount           string `json:"day_read_count,omitempty"`           // 当日阅读量
	DayWriteCount          string `json:"day_write_count,omitempty"`          // 当日编辑次数
	DayLikeCount           string `json:"day_like_count,omitempty"`           // 当日点赞量
	PostCount              string `json:"post_count,omitempty"`               // 文档数
	WordCount              string `json:"word_count,omitempty"`               // 字数
	WriteCount             string `json:"write_count,omitempty"`              // 编辑次数
	WriteCount30           string `json:"write_count_30,omitempty"`           // 编辑次数 (30天)
	ReadCount              string `json:"read_count,omitempty"`               // 阅读量
	ReadCount30            string `json:"read_count_30,omitempty"`            // 阅读量 (30天)
	ReadCount365           string `json:"read_count_365,omitempty"`           // 阅读量 (一年)
	LikeCount              string `json:"like_count,omitempty"`               // 点赞量
	LikeCount7             string `json:"like_count_7,omitempty"`             // 点赞量 (7天)
	LikeCount30            string `json:"like_count_30,omitempty"`            // 点赞量 (30天)
	LikeCount365           string `json:"like_count_365,omitempty"`           // 点赞量 (一年)
	WatchCount             string `json:"watch_count,omitempty"`              // 关注量
	WatchCount7            string `json:"watch_count_7,omitempty"`            // 关注量 (7天)
	WatchCount30           string `json:"watch_count_30,omitempty"`           // 关注量 (30天)
	WatchCount365          string `json:"watch_count_365,omitempty"`          // 关注量 (一年)
	CommentCount           string `json:"comment_count,omitempty"`            // 评论量
	CommentCount30         string `json:"comment_count_30,omitempty"`         // 评论量 (30天)
	CommentCount365        string `json:"comment_count_365,omitempty"`        // 评论量 (一年)
	LikeRankRate           string `json:"like_rank_rate,omitempty"`           // 知识库点赞数排名
	Popularity30           string `json:"popularity_30,omitempty"`            // 30 天热度
	DocCount               string `json:"doc_count,omitempty"`                // 文档数量
	SheetCount             string `json:"sheet_count,omitempty"`              // 表格数量
	BoardCount             string `json:"board_count,omitempty"`              // 画板数量
	ShowCount              string `json:"show_count,omitempty"`               // 演示文稿数量
	ResourceCount          string `json:"resource_count,omitempty"`           // 资源数量
	ArtBoardCount          string `json:"artboard_count,omitempty"`           // 图集数量
	AttachmentCount        string `json:"attachment_count,omitempty"`         // 附件数量
	InteractionPeopleCount string `json:"interaction_people_count,omitempty"` // 知识交流人数
	ContentCount           string `json:"content_count,omitempty"`            // 知识财富数
	CollaborationCount     string `json:"collaboration_count,omitempty"`      // 知识协同次数
	WorkingHours           string `json:"working_hours,omitempty"`            // 协同提效时长 (小时)
	Baike                  string `json:"baike,omitempty"`                    // 百科全书卷数
	TableCount             string `json:"table_count,omitempty"`              // 表格数量
}

// GetGroupDocStatistics 团队.文档统计数据
func (s *statisticService) GetGroupDocStatistics(ctx context.Context, login string, request *GetGroupDocStatisticsRequest, opts ...RequestOption) (*GetGroupDocStatisticsResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("groups/%s/statistics/docs", login), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response GetGroupDocStatisticsResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return &response, resp, nil
}

// GetGroupDocStatisticsRequest 团队文档统计请求
type GetGroupDocStatisticsRequest struct {
	BookID    *int       `url:"bookId,omitempty"`    // 指定知识库 [过滤条件]
	Name      *string    `url:"name,omitempty"`      // 文档名 [过滤条件]
	Range     *TimeRange `url:"range,omitempty"`     // 时间范围 [过滤条件]
	Page      *int       `url:"page,omitempty"`      // 页码
	Limit     *int       `url:"limit,omitempty"`     // 分页数量
	SortField *string    `url:"sortField,omitempty"` // 排序字段
	SortOrder *SortOrder `url:"sortOrder,omitempty"` // 排序方向
}

// GetGroupDocStatisticsResponse 团队文档统计响应
type GetGroupDocStatisticsResponse struct {
	Docs  []*DocStatistics `json:"docs,omitempty"`  // 文档统计数据
	Total int              `json:"total,omitempty"` // 总数量
}

// DocStatistics 文档统计数据
type DocStatistics struct {
	BizDate          string `json:"bizdate,omitempty"`            // 统计日期 (YYYYMMDD)
	BookID           string `json:"book_id,omitempty"`            // 知识库 ID
	DocID            string `json:"doc_id,omitempty"`             // 文档 ID
	Slug             string `json:"slug,omitempty"`               // 文档 slug
	Title            string `json:"title,omitempty"`              // 文档标题
	Type             string `json:"type,omitempty"`               // 知识库类型
	IsPublic         string `json:"is_public,omitempty"`          // 是否公开
	CreatedAt        string `json:"created_at,omitempty"`         // 创建时间
	ContentUpdatedAt string `json:"content_updated_at,omitempty"` // 最近更新时间
	UserID           string `json:"user_id,omitempty"`            // 文档归属 ID
	OrganizationID   string `json:"organization_id,omitempty"`    // 文档归属空间 ID
	DayReadCount     string `json:"day_read_count,omitempty"`     // 当日阅读量
	DayWriteCount    string `json:"day_write_count,omitempty"`    // 当日编辑次数
	DayLikeCount     string `json:"day_like_count,omitempty"`     // 当日点赞量
	WordCount        string `json:"word_count,omitempty"`         // 字数
	WriteCount       string `json:"write_count,omitempty"`        // 编辑次数
	ReadCount        string `json:"read_count,omitempty"`         // 阅读量
	ReadCount7       string `json:"read_count_7,omitempty"`       // 阅读量 (7天)
	ReadCount30      string `json:"read_count_30,omitempty"`      // 阅读量 (30天)
	ReadCount365     string `json:"read_count_365,omitempty"`     // 阅读量 (一年)
	LikeCount        string `json:"like_count,omitempty"`         // 点赞量
	LikeCount7       string `json:"like_count_7,omitempty"`       // 点赞量 (7天)
	LikeCount30      string `json:"like_count_30,omitempty"`      // 点赞量 (30天)
	LikeCount365     string `json:"like_count_365,omitempty"`     // 点赞量 (一年)
	CommentCount     string `json:"comment_count,omitempty"`      // 评论量
	CommentCount30   string `json:"comment_count_30,omitempty"`   // 评论量 (30天)
	CommentCount365  string `json:"comment_count_365,omitempty"`  // 评论量 (一年)
	Popularity30     string `json:"popularity_30,omitempty"`      // 30 天热度
	AttachmentCount  string `json:"attachment_count,omitempty"`   // 附件数量
	User             string `json:"user,omitempty"`               // 用户信息
}

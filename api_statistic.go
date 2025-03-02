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

// 团队.成员统计数据
// 团队.知识库统计数据
// 团队.文档统计数据

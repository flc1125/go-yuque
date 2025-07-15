package yuque

import (
	"context"
	"net/http"
)

type searchService struct {
	client *Client
}

// Search 通用搜索
func (s *searchService) Search(ctx context.Context, request *SearchRequest, opts ...RequestOption) (*SearchResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "search", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var results []*SearchResult
	resp, err := s.client.Do(req, &results)
	if err != nil {
		return nil, resp, err
	}

	var meta *MetaInfo
	if metaData := resp.meta(); metaData != nil {
		meta = &MetaInfo{
			Total: metaData.Total,
		}
	}

	return &SearchResponse{
		Meta: meta,
		Data: results,
	}, resp, nil
}

// SearchRequest 搜索请求
type SearchRequest struct {
	Q         string     `url:"q"`                   // 搜索关键词 (必填)
	Type      SearchType `url:"type"`                // 搜索类型 (必填)
	Scope     *string    `url:"scope,omitempty"`     // 搜索范围，不填默认为搜索当前用户/团队
	Page      *int       `url:"page,omitempty"`      // 页码 [分页参数]
	Offset    *int       `url:"offset,omitempty"`    // 页码，非偏移量 [分页参数] (已废弃)
	Creator   *string    `url:"creator,omitempty"`   // 仅搜索指定作者 login [筛选条件]
	CreatorID *int       `url:"creatorId,omitempty"` // 仅搜索指定作者 ID [筛选条件] (已废弃)
}

// SearchResponse 搜索响应
type SearchResponse struct {
	Meta *MetaInfo       `json:"meta,omitempty"`
	Data []*SearchResult `json:"data,omitempty"`
}

// MetaInfo 元信息
type MetaInfo struct {
	Total int `json:"total,omitempty"` // 结果总量
}
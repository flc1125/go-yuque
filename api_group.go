package yuque

import (
	"context"
	"fmt"
	"net/http"
)

type groupService struct {
	client *Client
}

// GetUserGroups 获取用户的团队
func (s *groupService) GetUserGroups(ctx context.Context, userID string, request *GetUserGroupsRequest, opts ...RequestOption) ([]*Group, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("users/%s/groups", userID), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var groups []*Group
	resp, err := s.client.Do(req, &groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}

// GetUserGroupsRequest 获取用户团队请求
type GetUserGroupsRequest struct {
	Role   *GroupMemberRole `url:"role,omitempty"`   // 角色 [过滤条件]
	Offset *int             `url:"offset,omitempty"` // 偏移量 [分页条件]
}

// GetGroupMembers 获取团队的成员
func (s *groupService) GetGroupMembers(ctx context.Context, groupLogin string, request *GetGroupMembersRequest, opts ...RequestOption) ([]*GroupUser, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("groups/%s/users", groupLogin), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var members []*GroupUser
	resp, err := s.client.Do(req, &members)
	if err != nil {
		return nil, resp, err
	}

	return members, resp, nil
}

// GetGroupMembersRequest 获取团队成员请求
type GetGroupMembersRequest struct {
	Role   *GroupMemberRole `url:"role,omitempty"`   // 角色 [筛选条件]
	Offset *int             `url:"offset,omitempty"` // 偏移量 [分页条件]
}

// UpdateGroupMember 变更成员
func (s *groupService) UpdateGroupMember(ctx context.Context, groupLogin string, userID string, request *UpdateGroupMemberRequest, opts ...RequestOption) (*GroupUser, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("groups/%s/users/%s", groupLogin, userID), request, opts)
	if err != nil {
		return nil, nil, err
	}

	var member *GroupUser
	resp, err := s.client.Do(req, &member)
	if err != nil {
		return nil, resp, err
	}

	return member, resp, nil
}

// UpdateGroupMemberRequest 更新团队成员请求
type UpdateGroupMemberRequest struct {
	Role GroupMemberRole `json:"role"` // 角色 (0:管理员, 1:成员, 2:只读成员)
}

// DeleteGroupMember 删除成员
func (s *groupService) DeleteGroupMember(ctx context.Context, groupLogin string, userID string, opts ...RequestOption) (*DeleteGroupMemberResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("groups/%s/users/%s", groupLogin, userID), nil, opts)
	if err != nil {
		return nil, nil, err
	}

	var result *DeleteGroupMemberResponse
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// DeleteGroupMemberResponse 删除团队成员响应
type DeleteGroupMemberResponse struct {
	UserID string `json:"user_id,omitempty"` // 删除用户 ID
}

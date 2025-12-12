package yuque

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatisticService_GetGroupStatistics(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/groups/group_name/statistics", r.URL.Path)

		_, _ = w.Write(loadData(t, "internal/testdata/api/statistic/get_group_statistics.json"))
	}))

	resp, _, err := client.StatisticService.GetGroupStatistics(ctx, "group_name")
	require.NoError(t, err)

	assert.Equal(t, &GetGroupStatisticsResponse{
		BizDate:                "20250301",
		UserID:                 1781111,
		OrganizationID:         35111,
		MemberCount:            26,
		CollaboratorCount:      58,
		DayReadCount:           0,
		DayWriteCount:          0,
		WriteCount:             75556,
		ReadCount:              20822,
		ReadCount30:            756,
		ReadCount365:           8917,
		CommentCount:           507,
		CommentCount30:         11,
		CommentCount365:        249,
		LikeCount:              84,
		LikeCount30:            11,
		LikeCount365:           39,
		FollowCount:            17,
		CollectCount:           27,
		DocCount:               1623,
		SheetCount:             151,
		BoardCount:             29,
		ShowCount:              0,
		ResourceCount:          0,
		ArtBoardCount:          0,
		AttachmentCount:        0,
		BookCount:              28,
		PublicBookCount:        1,
		PrivateBookCount:       27,
		BookBookCount:          25,
		BookResourceCount:      0,
		BookDesignCount:        0,
		BookThreadCount:        3,
		DataUsage:              0,
		GrainsCount:            0,
		GrainsCountSum:         0,
		GrainsCountConsume:     0,
		InteractionPeopleCount: 34,
		ContentCount:           1806,
		CollaborationCount:     97013,
		WorkingHours:           38870,
		Baike:                  1.7,
		TableCount:             3,
	}, resp)
}

func TestStatisticService_GetMemberStatistics(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/groups/group_name/statistics/members", r.URL.Path)
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "write_doc_count", r.URL.Query().Get("sortField"))
		assert.Equal(t, "desc", r.URL.Query().Get("sortOrder"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/statistic/get_member_statistics.json"))
	}))

	resp, _, err := client.StatisticService.GetMemberStatistics(ctx, "group_name", &GetMemberStatisticsRequest{
		Page:      Ptr(1),
		Limit:     Ptr(10),
		SortField: Ptr("write_doc_count"),
		SortOrder: Ptr("desc"),
	})
	require.NoError(t, err)

	assert.Equal(t, 25, resp.Total)
	require.Len(t, resp.Members, 2)

	// 验证第一个成员
	member1 := resp.Members[0]
	assert.Equal(t, 95500, member1.UserID)
	assert.Equal(t, "flc1125", member1.Login)
	assert.Equal(t, "张三", member1.Name)
	assert.Equal(t, 128, member1.WriteDocCount)
	assert.Equal(t, 456, member1.WriteCount)
	assert.Equal(t, 1024, member1.ReadCount)
	assert.Equal(t, 32, member1.LikeCount)

	// 验证第二个成员
	member2 := resp.Members[1]
	assert.Equal(t, 181111, member2.UserID)
	assert.Equal(t, "lisi", member2.Login)
	assert.Equal(t, "李四", member2.Name)
}

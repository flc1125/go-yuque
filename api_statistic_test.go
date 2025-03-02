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

package yuque

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_Hello(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/hello", r.URL.Path)

		_, _ = w.Write(loadData(t, "internal/testdata/api/user/hello.json"))
	}))

	resp, _, err := client.UserService.Hello(ctx)
	require.NoError(t, err)
	assert.Equal(t, &HelloResponse{
		Message: "Hello Yuque",
	}, resp)
}

func TestUserService_GetUser(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/user", r.URL.Path)

		_, _ = w.Write(loadData(t, "internal/testdata/api/user/user.json"))
	}))

	resp, _, err := client.UserService.GetUser(ctx)
	require.NoError(t, err)
	assert.Equal(t, &User{
		ID:               1781111,
		Type:             "Group",
		Login:            "yuque",
		Name:             "yuque group",
		AvatarURL:        "https://cdn.nlark.com/yuque/0/2022/png/95500/1111-avatar/81111.png",
		BooksCount:       28,
		PublicBooksCount: 1,
		FollowersCount:   0,
		FollowingCount:   0,
		Public:           0,
		Description:      "沉淀新人手册",
		CreatedAt:        time.Date(2020, 7, 16, 3, 1, 36, 0, time.UTC),
		UpdatedAt:        time.Date(2025, 2, 10, 8, 19, 55, 0, time.UTC),
		WorkID:           "",
	}, resp)
}

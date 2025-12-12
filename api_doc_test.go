package yuque

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocService_GetDocs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/repos/org/book/docs", r.URL.Path)
		assert.Equal(t, "0", r.URL.Query().Get("offset"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "hits,tags", r.URL.Query().Get("optional_properties"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/doc/get_docs.json"))
	}))

	resp, _, err := client.DocService.GetDocs(ctx, "org/book", &GetDocsRequest{
		Offset:             Ptr(0),
		Limit:              Ptr(10),
		OptionalProperties: Ptr("hits,tags"),
	})
	require.NoError(t, err)

	assert.Equal(t, 53, resp.Total)
	assert.Len(t, resp.Docs, 2)

	doc := resp.Docs[0]
	assert.Equal(t, 200952222, doc.ID)
	assert.Equal(t, "会议室演示", doc.Title)
	assert.Equal(t, DocTypeDoc, doc.Type)
}

func TestDocService_CreateDocs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var req CreateDocRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, "test", *req.Slug)
		assert.Equal(t, "无标题", *req.Title)
		assert.Equal(t, AccessTypePrivate, *req.Public)
		assert.Equal(t, DocFormatMarkdown, *req.Format)
		assert.Equal(t, "string", *req.Body)

		_, _ = w.Write(loadData(t, "internal/testdata/api/doc/create_doc.json"))
	}))

	doc, _, err := client.DocService.CreateDoc(ctx, "org/book", &CreateDocRequest{
		Slug:   Ptr("test"),
		Title:  Ptr("无标题"),
		Public: Ptr(AccessTypePrivate),
		Format: Ptr(DocFormatMarkdown),
		Body:   Ptr("string"),
	})
	require.NoError(t, err)

	assert.Equal(t, &Doc{
		ID:               20751111,
		Type:             DocTypeDoc,
		Slug:             "string",
		Title:            "无标题",
		Description:      "",
		Cover:            "",
		UserID:           12222,
		BookID:           1292222,
		LastEditorID:     12222,
		Format:           Ptr(DocFormatMarkdown),
		BodyDraft:        Ptr(""),
		Body:             Ptr("string"),
		BodyHTML:         Ptr("<p>string</p>\n"),
		Public:           AccessTypePrivate,
		Status:           1,
		LikesCount:       nil,
		ReadCount:        0,
		CommentsCount:    nil,
		WordCount:        1,
		CreatedAt:        mustParseTime(t, "2025-02-25T13:40:06.701Z"),
		UpdatedAt:        mustParseTime(t, "2025-02-25T13:40:06.701Z"),
		ContentUpdatedAt: mustParseTime(t, "2025-02-25T13:40:07.000Z"),
		PublishedAt:      mustParseTime(t, "2025-02-25T13:40:06.662Z"),
		FirstPublishedAt: mustParseTime(t, "2025-02-25T13:40:06.662Z"),
		Hits:             0,
	}, doc)
}

func TestDocService_CreateTOCs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/repos/org/book/toc", r.URL.Path)

		_, _ = w.Write(loadData(t, "internal/testdata/api/doc/get_tocs.json"))
	}))

	tocs, _, err := client.DocService.GetTOCs(ctx, "org/book")
	require.NoError(t, err)
	require.Len(t, tocs, 2)

	assert.Contains(t, tocs, &TOC{
		UUID:        "5wNEEZX3KK_hwec1",
		Type:        TOCTypeDoc,
		Title:       "团队建设",
		URL:         "akc74cm47w1kpi9g",
		Slug:        "akc74cm47w1kpi9g",
		ID:          163724494,
		DocID:       163724494,
		Level:       0,
		Depth:       1,
		OpenWindow:  1,
		Visible:     0,
		PrevUUID:    "",
		SiblingUUID: "XmljcRk3475oflnH",
		ChildUUID:   "ZNUPE7oxf-FP_WbQ",
		ParentUUID:  "",
	}, tocs)

	assert.Contains(t, tocs, &TOC{
		UUID:        "XmljcRk3475oflnH",
		Type:        TOCTypeDoc,
		Title:       "团队文化",
		URL:         "tti0lofsdvypikku",
		Slug:        "tti0lofsdvypikku",
		ID:          0,
		DocID:       0,
		Level:       0,
		Depth:       1,
		OpenWindow:  1,
		Visible:     1,
		PrevUUID:    "5wNEEZX3KK_hwec1",
		SiblingUUID: "Rj13VAHZwhw00nCn",
		ChildUUID:   "",
		ParentUUID:  "",
	}, tocs)
}

func TestDocService_GetDoc(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/repos/org/book/docs/200952222", r.URL.Path)

		_, _ = w.Write(loadData(t, "internal/testdata/api/doc/get_doc.json"))
	}))

	doc, _, err := client.DocService.GetDoc(ctx, "org/book", 200952222)
	require.NoError(t, err)

	// 验证基本字段
	assert.Equal(t, 200952222, doc.ID)
	assert.Equal(t, DocTypeDoc, doc.Type)
	assert.Equal(t, "gvbblbqbgmzmew75", doc.Slug)
	assert.Equal(t, "会议室演示", doc.Title)
	assert.Equal(t, AccessTypePrivate, doc.Public)
	assert.Equal(t, 757, doc.WordCount)
	assert.Equal(t, 128, doc.ReadCount)

	// 验证正文内容
	require.NotNil(t, doc.Body)
	assert.Equal(t, "# 会议室演示\n\n这是文档正文内容", *doc.Body)
	require.NotNil(t, doc.BodyHTML)
	assert.Contains(t, *doc.BodyHTML, "<h1>会议室演示</h1>")
	require.NotNil(t, doc.Format)
	assert.Equal(t, DocFormatMarkdown, *doc.Format)

	// 验证用户信息
	require.NotNil(t, doc.User)
	assert.Equal(t, 181111, doc.User.ID)
	assert.Equal(t, "张三", doc.User.Name)

	// 验证知识库信息
	require.NotNil(t, doc.Book)
	assert.Equal(t, 1292222, doc.Book.ID)
	assert.Equal(t, "新人指南", doc.Book.Name)

	// 验证标签
	require.Len(t, doc.Tags, 2)
	assert.Equal(t, "入门", doc.Tags[0].Title)
	assert.Equal(t, "教程", doc.Tags[1].Title)

	// 验证时间字段
	assert.Equal(t, mustParseTime(t, "2025-01-02T01:29:30.000Z"), doc.CreatedAt)
	assert.Equal(t, mustParseTime(t, "2025-02-08T03:26:48.000Z"), doc.UpdatedAt)
}

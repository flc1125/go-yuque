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
		BodyHtml:         Ptr("<p>string</p>\n"),
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

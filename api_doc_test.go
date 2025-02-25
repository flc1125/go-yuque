package yuque

import (
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

package yuque

type DocType string

// 文档类型 (Doc:普通文档, Sheet:表格, Thread:话题, Board:图集, Table:数据表)
const (
	DocTypeDoc    DocType = "Doc"
	DocTypeSheet  DocType = "Sheet"
	DocTypeThread DocType = "Thread"
	DocTypeBoard  DocType = "Board"
	DocTypeTable  DocType = "Table"
)

package yuque

// DocType 文档类型 (Doc:普通文档, Sheet:表格, Thread:话题, Board:图集, Table:数据表)
type DocType string

const (
	DocTypeDoc    DocType = "Doc"    // 普通文档
	DocTypeSheet  DocType = "Sheet"  // 表格
	DocTypeThread DocType = "Thread" // 话题
	DocTypeBoard  DocType = "Board"  // 图集
	DocTypeTable  DocType = "Table"  // 数据表
)

// AccessType 文档/知识库公开性:公开性 (0:私密, 1:公开, 2:企业内公开)
type AccessType int

const (
	AccessTypePrivate AccessType = 0 // 私密
	AccessTypePublic  AccessType = 1 // 公开
	AccessTypeInner   AccessType = 2 // 企业内公开
)

// DocFormat 内容格式 (markdown:Markdown 格式, html:HTML 标准格式, lake:语雀 Lake 格式，lakesheet:语雀表格)
type DocFormat string

const (
	DocFormatMarkdown  DocFormat = "markdown"  // Markdown 格式
	DocFormatHTML      DocFormat = "html"      // HTML 标准格式
	DocFormatLake      DocFormat = "lake"      // 语雀 Lake 格式
	DocFormatLakeSheet DocFormat = "lakesheet" // 语雀表格
)

// BookType 类型 (Book:文档, Design:图集, Sheet:表格, Resource:资源)
type BookType string

const (
	BookTypeBook     BookType = "Book"     // 文档
	BookTypeDesign   BookType = "Design"   // 图集
	BookTypeSheet    BookType = "Sheet"    // 表格
	BookTypeResource BookType = "Resource" // 资源
)

package yuque

import "time"

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

// SearchType 搜索类型 (doc:文档, repo:知识库)
type SearchType string

const (
	SearchTypeDoc  SearchType = "doc"  // 文档
	SearchTypeRepo SearchType = "repo" // 知识库
)

// GroupMemberRole 成员角色 (0:管理员, 1:成员, 2:只读成员)
type GroupMemberRole int

const (
	GroupMemberRoleAdmin    GroupMemberRole = 0 // 管理员
	GroupMemberRoleMember   GroupMemberRole = 1 // 成员
	GroupMemberRoleReadOnly GroupMemberRole = 2 // 只读成员
)

// TOCAction 目录操作 (appendNode:尾插, prependNode:头插, editNode:编辑节点, removeNode:删除节点)
type TOCAction string

const (
	TOCActionAppendNode  TOCAction = "appendNode"  // 尾插
	TOCActionPrependNode TOCAction = "prependNode" // 头插
	TOCActionEditNode    TOCAction = "editNode"    // 编辑节点
	TOCActionRemoveNode  TOCAction = "removeNode"  // 删除节点
)

// TOCActionMode 操作模式 (sibling:同级, child:子级)
type TOCActionMode string

const (
	TOCActionModeSibling TOCActionMode = "sibling" // 同级
	TOCActionModeChild   TOCActionMode = "child"   // 子级
)

// SortOrder 排序方向
type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"  // 升序
	SortOrderDesc SortOrder = "desc" // 降序
)

// TimeRange 时间范围
type TimeRange int

const (
	TimeRangeAll  TimeRange = 0   // 全部
	TimeRange30D  TimeRange = 30  // 近30天
	TimeRange365D TimeRange = 365 // 近一年
	TimeRange7D   TimeRange = 7   // 近7天
)

// Tag 标签
type Tag struct {
	ID        int       `json:"id,omitempty"`         // TAG ID
	Title     string    `json:"title,omitempty"`      // TAG NAME
	DocID     int       `json:"doc_id,omitempty"`     // 文档 ID
	BookID    int       `json:"book_id,omitempty"`    // 知识库 ID
	UserID    int       `json:"user_id,omitempty"`    // 创建者 ID
	CreatedAt time.Time `json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `json:"updated_at,omitempty"` // 更新时间
}

// GroupUser 团队成员
type GroupUser struct {
	ID        int             `json:"id,omitempty"`         // ID
	GroupID   int             `json:"group_id,omitempty"`   // 团队 ID
	UserID    int             `json:"user_id,omitempty"`    // 成员 ID
	Role      GroupMemberRole `json:"role,omitempty"`       // 成员角色
	CreatedAt time.Time       `json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time       `json:"updated_at,omitempty"` // 更新时间
	Group     *Group          `json:"group,omitempty"`      // 团队信息
	User      *User           `json:"user,omitempty"`       // 用户信息
}

// Group 团队
type Group struct {
	ID               int       `json:"id,omitempty"`                 // 团队 ID
	Type             string    `json:"type,omitempty"`               // 类型 Always 'Group'
	Login            string    `json:"login,omitempty"`              // 路径
	Name             string    `json:"name,omitempty"`               // 名称
	AvatarURL        string    `json:"avatar_url,omitempty"`         // 头像
	BooksCount       int       `json:"books_count,omitempty"`        // 知识库数量
	PublicBooksCount int       `json:"public_books_count,omitempty"` // 公开的知识库数量
	MembersCount     int       `json:"members_count,omitempty"`      // 成员人数
	Public           int       `json:"public,omitempty"`             // 公开性
	Description      string    `json:"description,omitempty"`        // 介绍
	CreatedAt        time.Time `json:"created_at,omitempty"`         // 创建时间
	UpdatedAt        time.Time `json:"updated_at,omitempty"`         // 更新时间
}

// SearchResult 搜索结果
type SearchResult struct {
	ID      int        `json:"id,omitempty"`      // ID
	Type    SearchType `json:"type,omitempty"`    // 类型
	Title   string     `json:"title,omitempty"`   // 标题
	Summary string     `json:"summary,omitempty"` // 摘要
	URL     string     `json:"url,omitempty"`     // 访问路径
	Info    string     `json:"info,omitempty"`    // 归属信息
	Target  any        `json:"target,omitempty"`  // 目标对象 (Doc 或 Book)
}

// DocVersion 文档版本
type DocVersion struct {
	ID        int       `json:"id,omitempty"`         // 版本 ID
	DocID     int       `json:"doc_id,omitempty"`     // 文档 ID
	Slug      string    `json:"slug,omitempty"`       // 文档路径
	Title     string    `json:"title,omitempty"`      // 文档标题
	UserID    int       `json:"user_id,omitempty"`    // 发版人 ID
	CreatedAt time.Time `json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `json:"updated_at,omitempty"` // 更新时间
	User      *User     `json:"user,omitempty"`       // 用户信息
}

// DocVersionDetail 文档版本详情
type DocVersionDetail struct {
	ID        int       `json:"id,omitempty"`         // 版本 ID
	DocID     int       `json:"doc_id,omitempty"`     // 文档 ID
	Slug      string    `json:"slug,omitempty"`       // 文档路径
	Title     string    `json:"title,omitempty"`      // 文档标题
	UserID    int       `json:"user_id,omitempty"`    // 发版人 ID
	Format    DocFormat `json:"format,omitempty"`     // 内容格式
	Body      string    `json:"body,omitempty"`       // 正文原始内容
	BodyHTML  string    `json:"body_html,omitempty"`  // 正文 HTML 标准格式内容
	BodyASL   string    `json:"body_asl,omitempty"`   // 正文语雀 Lake 格式内容
	Diff      string    `json:"diff,omitempty"`       // 版本 DIFF
	CreatedAt time.Time `json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `json:"updated_at,omitempty"` // 更新时间
	User      *User     `json:"user,omitempty"`       // 用户信息
}

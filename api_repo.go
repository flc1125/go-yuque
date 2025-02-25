package yuque

import "time"

type Book struct {
	ID               int        `json:"id"`
	Type             BookType   `json:"type"`
	Slug             string     `json:"slug"`
	Name             string     `json:"name"`
	UserID           int        `json:"user_id"`
	Description      string     `json:"description"`
	TocYML           string     `json:"toc_yml"`
	CreatorID        int        `json:"creator_id"`
	Public           AccessType `json:"public"`
	ItemsCount       int        `json:"items_count"`
	LikesCount       int        `json:"likes_count"`
	WatchesCount     int        `json:"watches_count"`
	ContentUpdatedAt time.Time  `json:"content_updated_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	Namespace        string     `json:"namespace"`
	User             *User      `json:"user"`
}

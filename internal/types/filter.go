package types

type Filter struct {
	Keyword    string `json:"keyword,omitempty" form:"keyword"`
	UserID     int64  `json:"user_id,omitempty" form:"user_id"`
	CategoryID int64  `json:"category_id,omitempty" form:"category_id"`
}

package types

type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

type DataListRes struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

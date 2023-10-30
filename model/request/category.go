package request

import (
	"new-mall/model/common"
)

type CategoryReq struct {
	CategoryId    int    `json:"categoryId"`
	CategoryLevel int    `json:"categoryLevel" `
	ParentId      int    `json:"parentId"`
	CategoryName  string `json:"categoryName" `
	CategoryRank  string `json:"categoryRank" `
}

type SearchCategoryParams struct {
	CategoryLevel int `json:"categoryLevel" form:"categoryLevel"`
	ParentId      int `json:"parentId" form:"parentId"`
	common.PageInfo
}

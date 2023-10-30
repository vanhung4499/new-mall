package common

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

// GetById Find by id structure
type GetById struct {
	ID float64 `json:"id" form:"id"`
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // Role ID
}

type Empty struct{}

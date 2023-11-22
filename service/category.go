package service

import (
	"context"
	"new-mall/pkg/utils"
	"new-mall/repository/db/dao"
	"new-mall/types"
	"sync"
)

var CategorySrvIns *CategorySrv
var CategorySrvOnce sync.Once

type CategorySrv struct {
}

func GetCategorySrv() *CategorySrv {
	CategorySrvOnce.Do(func() {
		CategorySrvIns = &CategorySrv{}
	})
	return CategorySrvIns
}

func (s *CategorySrv) CategoryList(ctx context.Context, req *types.ListCategoryReq) (resp interface{}, err error) {
	categories, err := dao.NewCategoryDao(ctx).ListCategory()
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	cRes := make([]*types.ListCategoryResp, 0)
	for _, v := range categories {
		cRes = append(cRes, &types.ListCategoryResp{
			ID:           v.ID,
			CategoryName: v.CategoryName,
			CreatedAt:    v.CreatedAt.Unix(),
		})
	}

	resp = &types.DataListRes{
		Item:  cRes,
		Total: int64(len(cRes)),
	}

	return
}

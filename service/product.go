package service

import (
	"context"
	"mime/multipart"
	"new-mall/config"
	"new-mall/constant"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
	"new-mall/pkg/utils/upload"
	"new-mall/repository/db/dao"
	"new-mall/repository/db/model"
	"new-mall/types"
	"sync"
)

var ProductSrvIns *ProductSrv
var ProductSrvOnce sync.Once

type ProductSrv struct {
}

func GetProductSrv() *ProductSrv {
	ProductSrvOnce.Do(func() {
		ProductSrvIns = &ProductSrv{}
	})
	return ProductSrvIns
}

func (s *ProductSrv) ProductShow(ctx context.Context, req *types.ProductShowReq) (res interface{}, err error) {
	p, err := dao.NewProductDao(ctx).ShowProductById(req.ID)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	pRes := &types.ProductRes{
		ID:            p.ID,
		Name:          p.Name,
		CategoryID:    p.CategoryID,
		Title:         p.Title,
		Info:          p.Info,
		ImagePath:     p.ImagePath,
		Price:         p.Price,
		DiscountPrice: p.DiscountPrice,
		View:          p.View(),
		CreatedAt:     p.CreatedAt.Unix(),
		Num:           p.Num,
		OnSale:        p.OnSale,
		BossID:        p.BossID,
		BossName:      p.BossName,
		BossAvatar:    p.BossAvatar,
	}

	res = pRes

	return
}

func (s *ProductSrv) ProductCreate(ctx context.Context, files []*multipart.FileHeader, req *types.ProductCreateReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	uId := u.Id
	boss, _ := dao.NewUserDao(ctx).GetUserById(uId)
	// Use the first one as the cover image
	oss := upload.NewOss()
	path, _, err := oss.UploadFile(files[0])

	if err != nil {
		utils.Logger.Error(err)
		return
	}
	product := &model.Product{
		Name:          req.Name,
		CategoryID:    req.CategoryID,
		Title:         req.Title,
		Info:          req.Info,
		ImagePath:     path,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Num:           req.Num,
		OnSale:        true,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for _, file := range files {
		path, _, err = oss.UploadFile(file)
		if err != nil {
			utils.Logger.Error(err)
			return
		}
		productImage := &model.ProductImage{
			ProductID: product.ID,
			ImagePath: path,
		}
		err = dao.NewProductImageDaoByDB(productDao.DB).CreateProductImage(productImage)
		if err != nil {
			utils.Logger.Error(err)
			return
		}
		wg.Done()
	}

	wg.Wait()

	return
}

func (s *ProductSrv) ProductList(ctx context.Context, req *types.ProductListReq) (res interface{}, err error) {
	var total int64
	condition := make(map[string]interface{})
	if req.CategoryID != 0 {
		condition["category_id"] = req.CategoryID
	}
	productDao := dao.NewProductDao(ctx)
	products, _ := productDao.ListProductByCondition(condition, req.BasePage)
	total, err = productDao.CountProductByCondition(condition)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	pResList := make([]*types.ProductRes, 0)
	for _, p := range products {
		pRes := &types.ProductRes{
			ID:            p.ID,
			Name:          p.Name,
			CategoryID:    p.CategoryID,
			Title:         p.Title,
			Info:          p.Info,
			ImagePath:     p.ImagePath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			View:          p.View(),
			CreatedAt:     p.CreatedAt.Unix(),
			Num:           p.Num,
			OnSale:        p.OnSale,
			BossID:        p.BossID,
			BossName:      p.BossName,
			BossAvatar:    p.BossAvatar,
		}
		pResList = append(pResList, pRes)
	}

	res = &types.DataListRes{
		Item:  pResList,
		Total: total,
	}

	return
}

func (s *ProductSrv) ProductDelete(ctx context.Context, req *types.ProductDeleteReq) (res interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	err = dao.NewProductDao(ctx).DeleteProduct(req.ID, u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	return
}

func (s *ProductSrv) ProductUpdate(ctx context.Context, req *types.ProductUpdateReq) (res interface{}, err error) {
	product := &model.Product{
		Name:       req.Name,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Info:       req.Info,
		// ImagePath:       service.ImagePath,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		OnSale:        req.OnSale,
	}
	err = dao.NewProductDao(ctx).UpdateProduct(req.ID, product)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}

func (s *ProductSrv) ProductSearch(ctx context.Context, req *types.ProductSearchReq) (res interface{}, err error) {
	products, count, err := dao.NewProductDao(ctx).SearchProduct(req.Info, req.BasePage)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	pResList := make([]*types.ProductRes, 0)
	for _, p := range products {
		pRes := &types.ProductRes{
			ID:            p.ID,
			Name:          p.Name,
			CategoryID:    p.CategoryID,
			Title:         p.Title,
			Info:          p.Info,
			ImagePath:     p.ImagePath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			View:          p.View(),
			CreatedAt:     p.CreatedAt.Unix(),
			Num:           p.Num,
			OnSale:        p.OnSale,
			BossID:        p.BossID,
			BossName:      p.BossName,
			BossAvatar:    p.BossAvatar,
		}
		pResList = append(pResList, pRes)
	}

	res = &types.DataListRes{
		Item:  pResList,
		Total: count,
	}

	return
}

func (s *ProductSrv) ProductImageList(ctx context.Context, req *types.ListProductImageReq) (res interface{}, err error) {
	productImages, _ := dao.NewProductImageDao(ctx).ListProductImageByProductId(req.ID)
	for i := range productImages {
		if config.Config.System.UploadModel == constant.UploadModelLocal {
			local := config.Config.Local
			productImages[i].ImagePath = local.Path + config.Config.System.HttpPort + local.StorePath + productImages[i].ImagePath
		}
	}

	res = &types.DataListRes{
		Item:  productImages,
		Total: int64(len(productImages)),
	}

	return
}

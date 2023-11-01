package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"new-mall/business"
	"new-mall/global"
	"new-mall/model"
	"new-mall/model/common"
	"new-mall/repository"
	"strconv"
)

func CreateAddress(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.AddressCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewAddressRepo(global.DB)
		biz := business.NewAddressBiz(repo)

		if err := biz.CreateAddress(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		common.Ok(c)
	}
}

func UpdateAddress(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.AddressUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewAddressRepo(global.DB)
		biz := business.NewAddressBiz(repo)

		if err := biz.UpdateAddress(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		common.Ok(c)
	}
}

func DeleteAddress(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewAddressRepo(global.DB)
		biz := business.NewAddressBiz(repo)

		if err := biz.DeleteAddress(c.Request.Context(), id); err != nil {
			panic(err)
		}

		common.Ok(c)
	}
}

func ListAddressByUser(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var filter common.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewAddressRepo(global.DB)
		biz := business.NewAddressBiz(repo)

		addresses, err := biz.ListAddress(c.Request.Context(), &filter, nil)
		if err != nil {
			panic(err)
		}

		common.OkWithData(addresses, c)
	}
}

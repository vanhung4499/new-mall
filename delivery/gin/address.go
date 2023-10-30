package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/business"
	"new-mall/global"
	"new-mall/model"
	"new-mall/model/common"
	"new-mall/repository"
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

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

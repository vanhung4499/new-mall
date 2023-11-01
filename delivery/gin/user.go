package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"new-mall/business"
	"new-mall/global"
	"new-mall/model/common"
	"new-mall/model/request"
	"new-mall/repository"
	"strconv"
)

func RegisterUser(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req request.RegisterUserRequest

		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewUserRepo(global.DB)
		biz := business.NewUserBiz(repo)

		if err := biz.RegisterUser(c.Request.Context(), &req); err != nil {
			panic(err)
		}

		common.Ok(c)
	}
}

func LoginUser(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req request.LoginUserRequest

		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewUserRepo(global.DB)
		biz := business.NewUserBiz(repo)

		token, err := biz.LoginUser(c.Request.Context(), &req)
		if err != nil {
			panic(err)
		}

		common.OkWithData(token, c)
	}
}

func GetUser(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewUserRepo(global.DB)
		biz := business.NewUserBiz(repo)

		user, err := biz.FindUser(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		common.OkWithData(user, c)
	}
}

func UpdateUser(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data request.UpdateUserRequest
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewUserRepo(global.DB)
		biz := business.NewUserBiz(repo)

		if err := biz.UpdateUser(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		common.Ok(c)
	}
}

func DeleteUser(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewUserRepo(global.DB)
		biz := business.NewUserBiz(repo)

		if err := biz.DeleteUser(c.Request.Context(), id); err != nil {
			panic(err)
		}

		common.Ok(c)
	}
}

func ListUser(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		repo := repository.NewUserRepo(global.DB)
		biz := business.NewUserBiz(repo)

		addresses, err := biz.ListUser(c.Request.Context(), nil, &pagingData)
		if err != nil {
			panic(err)
		}

		common.OkWithData(common.NewSuccessResponse(addresses, pagingData, nil), c)
	}
}

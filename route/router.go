package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	api "new-mall/api/v1"
	"new-mall/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("my-session", store))
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		// Auth operations
		v1.POST("auth/register", api.UserRegisterHandler())
		v1.POST("auth/login", api.UserLoginHandler())

		// Product operation
		v1.GET("product/list", api.ListProductsHandler())
		v1.GET("product/show", api.ShowProductHandler())
		v1.POST("product/search", api.SearchProductsHandler())
		v1.GET("product/image/list", api.ListProductImageHandler())
		v1.GET("category/list", api.ListCategoryHandler())
		v1.GET("carousels", api.ListCarouselsHandler())

		authed := v1.Group("/")
		authed.Use(middleware.AuthMiddleware())
		{
			// User operations
			authed.POST("user/update", api.UserUpdateHandler())
			authed.GET("user/show_info", api.ShowUserInfoHandler())
			authed.POST("user/send_email", api.SendEmailHandler())
			authed.GET("user/valid_email", api.ValidEmailHandler())
			authed.POST("user/following", api.UserFollowingHandler())
			authed.POST("user/unfollowing", api.UserUnFollowingHandler())
			authed.POST("user/avatar", api.UploadAvatarHandler())

			// product operations
			authed.POST("product/create", api.CreateProductHandler())
			authed.PUT("product/update", api.UpdateProductHandler())
			authed.DELETE("product/delete", api.DeleteProductHandler())

			// favorite operations
			authed.GET("favorite/list", api.ListFavoritesHandler())
			authed.POST("favorite/create", api.CreateFavoriteHandler())
			authed.POST("favorite/delete", api.DeleteFavoriteHandler())

			// order operations
			authed.POST("order/create", api.CreateOrderHandler())
			authed.GET("order/list", api.ListOrdersHandler())
			authed.GET("order/show", api.ShowOrderHandler())
			authed.DELETE("order/delete", api.DeleteOrderHandler())

			// cart operations
			authed.POST("carts/create", api.CreateCartHandler())
			authed.GET("carts/list", api.ListCartHandler())
			authed.POST("carts/update", api.UpdateCartHandler())
			authed.DELETE("carts/delete", api.DeleteCartHandler())

			// address operations
			authed.POST("addresses/create", api.CreateAddressHandler())
			authed.GET("addresses/show", api.ShowAddressHandler())
			authed.GET("addresses/list", api.ListAddressHandler())
			authed.PUT("addresses/update", api.UpdateAddressHandler())
			authed.DELETE("addresses/delete", api.DeleteAddressHandler())

			// payment operations
			//authed.POST("pay", api.OrderPaymentHandler())
		}
	}
	return r
}

package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/internal/controllers"
	"new-mall/internal/global"
	"new-mall/internal/middleware"
	"new-mall/internal/repositories"
	"new-mall/internal/services"
)

// SetupRoutes configures the routes for the application.
func SetupRoutes() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(middleware.Recover())
	r.Use(sessions.Sessions("my-session", store))
	r.StaticFS("/static", http.Dir("./static"))

	db := global.DB

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	imageRepo := repositories.NewImageRepository(db)
	favoriteRepo := repositories.NewFavoriteRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	addressRepo := repositories.NewAddressRepository(db)
	carouselRepo := repositories.NewCarouselRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo, imageRepo)
	favoriteService := services.NewFavoriteService(favoriteRepo, userRepo, productRepo)
	orderService := services.NewOrderService(orderRepo, addressRepo, cartRepo, userRepo)
	cartService := services.NewCartService(cartRepo, productRepo)
	addressService := services.NewAddressService(addressRepo)
	carouselService := services.NewCarouselService(carouselRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	favoriteController := controllers.NewFavoriteController(favoriteService)
	orderController := controllers.NewOrderController(orderService)
	cartController := controllers.NewCartController(cartService)
	addressController := controllers.NewAddressController(addressService)
	carouselController := controllers.NewCarouselController(carouselService)
	categoryController := controllers.NewCategoryController(categoryService)

	v1 := r.Group("api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		authRoutes := v1.Group("/auth")
		{
			// Auth operations
			authRoutes.POST("/register", userController.RegisterUser())
			authRoutes.POST("/login", userController.LoginUser())
		}

		productRoutes := v1.Group("/products")
		{
			productRoutes.GET("/", productController.ListProduct())
			productRoutes.GET("/:id", productController.GetProduct())
		}

		{
			v1.GET("categories", categoryController.ListCategory())
			v1.GET("carousels", carouselController.ListCarousel())
		}

		authedRoutes := v1.Group("/")
		authedRoutes.Use(middleware.RequiredAuth())
		{
			// User operations
			authedRoutes.PATCH("users", userController.UpdateUser())
			authedRoutes.GET("users/profile", userController.GetProfile())
			authedRoutes.GET("users", userController.GetProfile())
			authedRoutes.GET("users/:id", userController.GetUser())
			authedRoutes.POST("users/avatar", userController.UploadAvatar())

			// product operations
			productRoutes.POST("products", productController.CreateProduct())
			productRoutes.PATCH("products/:id", productController.UpdateProduct())
			productRoutes.DELETE("products/:id", productController.DeleteProduct())

			// favorite operations
			authedRoutes.GET("favorites", favoriteController.ListFavorite())
			authedRoutes.POST("favorites", favoriteController.CreateFavorite())
			authedRoutes.DELETE("favorites", favoriteController.DeleteFavorite())

			// order operations
			authedRoutes.POST("orders", orderController.CreateOrder())
			authedRoutes.GET("orders", orderController.ListOrder())
			authedRoutes.GET("orders/:id", orderController.GetOrder())
			authedRoutes.DELETE("orders/:id", orderController.DeleteOrder())

			// cart operations
			authedRoutes.POST("carts/add", cartController.Add())
			authedRoutes.PATCH("carts", cartController.UpdateCart())
			authedRoutes.DELETE("carts", cartController.DeleteCart())
			authedRoutes.DELETE("carts/items/:productId", cartController.DeleteCartItem())

			// address operations
			authedRoutes.POST("addresses", addressController.CreateAddress())
			authedRoutes.GET("addresses/:id", addressController.GetAddress())
			authedRoutes.GET("addresses", addressController.ListAddress())
			authedRoutes.PATCH("addresses/:id", addressController.UpdateAddress())
			authedRoutes.DELETE("addresses/:id", addressController.DeleteAddress())

		}
	}

	return r
}

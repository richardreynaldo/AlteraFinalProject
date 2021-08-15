package routes

import (
	"finalProject/controllers/articles"
	"finalProject/controllers/coffees"
	"finalProject/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
	ArticleController articles.ArticleController
	CoffeesController coffees.CoffeeController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)

	article := e.Group("article")
	article.GET("/list", cl.ArticleController.GetAll)
	article.POST("/create", cl.ArticleController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	article.PUT("/update", cl.ArticleController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	coffee := e.Group("coffee")
	coffee.GET("/list", cl.CoffeesController.GetAll)
	coffee.POST("/create", cl.CoffeesController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	coffee.PUT("/update", cl.CoffeesController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	// category := e.Group("category")
	// category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))

	// news := e.Group("news")
	// news.POST("/store", cl.NewsController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	// news.PUT("/update", cl.NewsController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
}

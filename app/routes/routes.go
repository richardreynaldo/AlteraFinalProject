package routes

import (
	"finalProject/controllers/articles"
	"finalProject/controllers/coffees"
	"finalProject/controllers/transaction_detail"
	"finalProject/controllers/transaction_header"
	"finalProject/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware               middleware.JWTConfig
	UserController              users.UserController
	ArticleController           articles.ArticleController
	CoffeesController           coffees.CoffeeController
	TransactionHeaderController transaction_header.TransactionHeaderController
	TransactionDetailController transaction_detail.TransactionDetailController
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
	coffee.GET("find/:id", cl.CoffeesController.FindById, middleware.JWTWithConfig(cl.JWTMiddleware))

	transaction := e.Group("transaction")
	transaction.GET("/list", cl.TransactionHeaderController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	transaction.POST("/create", cl.TransactionHeaderController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	transaction.PUT("/update", cl.TransactionHeaderController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
}

package routes

import (
	"finalProject/controllers/articles"
	"finalProject/controllers/coffees"
	"finalProject/controllers/reviews"
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
	ReviewController            reviews.ReviewController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.GET("/token", cl.UserController.CreateToken)

	article := e.Group("article")
	article.GET("/list", cl.ArticleController.GetAll)
	article.POST("/create", cl.ArticleController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	article.PUT("/update", cl.ArticleController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	article.GET("/find/:id", cl.ArticleController.FindById, middleware.JWTWithConfig(cl.JWTMiddleware))

	coffee := e.Group("coffee")
	coffee.GET("/list", cl.CoffeesController.GetAll)
	coffee.POST("/create", cl.CoffeesController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	coffee.PUT("/update", cl.CoffeesController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	coffee.GET("/find/:id", cl.CoffeesController.FindById, middleware.JWTWithConfig(cl.JWTMiddleware))

	transaction := e.Group("transaction")
	transaction.GET("/list", cl.TransactionDetailController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	transaction.POST("/create", cl.TransactionDetailController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	transaction.PUT("/update", cl.TransactionDetailController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	reviews := e.Group("review")
	reviews.POST("/create", cl.ReviewController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
}

package main

import (
	_userUsecase "finalProject/business/users"
	_userController "finalProject/controllers/users"
	_userRepo "finalProject/drivers/databases/users"

	_articleUsecase "finalProject/business/articles"
	_articleController "finalProject/controllers/articles"
	_articleRepo "finalProject/drivers/databases/articles"

	_coffeeUsecase "finalProject/business/coffees"
	_coffeeController "finalProject/controllers/coffees"
	_coffeeRepo "finalProject/drivers/databases/coffees"

	_transactionHeaderUsecase "finalProject/business/transaction_header"
	_transactionHeaderController "finalProject/controllers/transaction_header"
	_transactionHeaderRepo "finalProject/drivers/databases/transaction_header"

	_transactionDetailUsecase "finalProject/business/transaction_detail"
	_transactionDetailController "finalProject/controllers/transaction_detail"
	_transactionDetailRepo "finalProject/drivers/databases/transaction_detail"

	_reviewUsecase "finalProject/business/reviews"
	_reviewController "finalProject/controllers/reviews"
	_reviewRepo "finalProject/drivers/databases/reviews"

	_dbDriver "finalProject/drivers/mysql"

	// _ipLocatorDriver "finalProject/drivers/thirdparties/iplocator"

	_middleware "finalProject/app/middleware"
	_routes "finalProject/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.User{},
		&_articleRepo.Articles{},
		&_coffeeRepo.Coffees{},
		&_transactionHeaderRepo.TransactionHeader{},
		&_transactionDetailRepo.TransactionDetail{},
		&_reviewRepo.Reviews{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	// iplocatorRepo := _ipLocatorDriver.NewIPLocator()

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	articleRepo := _articleRepo.NewMySQLArticleRepository(db)
	articleUsecase := _articleUsecase.NewArticleUsecase(articleRepo, userUsecase, timeoutContext)
	articleCtrl := _articleController.NewArticleController(articleUsecase)

	coffeeRepo := _coffeeRepo.NewMySQLCoffeesRepository(db)
	coffeeUsecase := _coffeeUsecase.NewCoffeesUsecase(coffeeRepo, timeoutContext)
	coffeeCtrl := _coffeeController.NewCoffeesController(coffeeUsecase)

	transactionDetailRepo := _transactionDetailRepo.NewMySQLTransactionDetailRepository(db)
	transactionDetailUsecase := _transactionDetailUsecase.NewTransactionDetailUsecase(transactionDetailRepo, timeoutContext)

	transactionRepo := _transactionHeaderRepo.NewMySQLTransactionHeaderRepository(db)
	transactionUsecase := _transactionHeaderUsecase.NewTransactionHeaderUsecase(transactionRepo, transactionDetailUsecase, timeoutContext)
	transactionCtrl := _transactionHeaderController.NewTransactionHeaderController(transactionUsecase)
	transactionDetailCtrl := _transactionDetailController.NewTransactionDetailController(transactionDetailUsecase, transactionUsecase)

	reviewRepo := _reviewRepo.NewMySQLReviewRepository(db)
	reviewUsecase := _reviewUsecase.NewArticleUsecase(reviewRepo, userUsecase, timeoutContext)
	reviewCtrl := _reviewController.NewReviewController(reviewUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:               configJWT.Init(),
		UserController:              *userCtrl,
		ArticleController:           *articleCtrl,
		CoffeesController:           *coffeeCtrl,
		TransactionHeaderController: *transactionCtrl,
		TransactionDetailController: *transactionDetailCtrl,
		ReviewController:            *reviewCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

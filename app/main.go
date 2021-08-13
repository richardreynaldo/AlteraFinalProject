package main

import (
	_userUsecase "finalProject/business/users"
	_userController "finalProject/controllers/users"
	_userRepo "finalProject/drivers/databases/users"

	_articleUsecase "finalProject/business/articles"
	_articleController "finalProject/controllers/articles"
	_articleRepo "finalProject/drivers/databases/articles"

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

	routesInit := _routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		UserController:    *userCtrl,
		ArticleController: *articleCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}

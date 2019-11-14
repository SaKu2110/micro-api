package main

import(
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/micro-api/sign/config"
	"github.com/SaKu2110/micro-api/sign/controller"
	"github.com/SaKu2110/micro-api/sign/model"
)

func main() {
	var ctrl controller.IsController
	db, err := initializeDataBase()
	if err != nil {
		// TODO: Faild connecting mysql
	}
	ctrl = initializeController(db)
	router := setupRouter(ctrl)
	err = router.Run(":9000")
	if err != nil {
		// TODO: Faild launching router
	}
}

func initializeDataBase() (*gorm.DB, error){
	var db *gorm.DB
	var err error
	var count time.Duration
	token := config.GetConnectionToken()

	count = 1
	for {
		if count > 5 {
			return nil, fmt.Errorf("")
		}
		db, err = gorm.Open("mysql", token)
		if err == nil {
			db.AutoMigrate(&model.User{})
			return db, nil
		}
		time.Sleep(3 * time.Second)

		count++
	}

	return nil, err
}

func initializeController(db *gorm.DB) (controller.IsController){
	return controller.IsController{DB: db}
}

func setupRouter(ctrl controller.IsController) *gin.Engine {
	router := gin.Default()
	router.GET("/", ctrl.GetFunctionMap)
	router.GET("/users", ctrl.GetUsersHandler)
	router.GET("/user/:id", ctrl.GetUserHandler)
	router.POST("/signin", ctrl.SigninHandler)
	router.POST("/signup", ctrl.SignupHandler)
	router.POST("/delete", ctrl.DeleteUserHandler)
	return router
}
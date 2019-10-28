package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/micro-api/sign/config"
	"github.com/SaKu2110/micro-api/sign/controller"
	"github.com/SaKu2110/micro-api/sign/model"
)

func main() {
	db, err := initializeDataBase()
	if err != nil {
		// TODO: Faild connecting mysql
	}
	ctrl　= initializeController(db)
	router := setupRouter(*ctrl)
	err = router.Run(":9000")
	if err != nil {
		// TODO: Faild launching router
	}
}

func initializeDataBase() (*gorm.DB, error){
	var db *gorm.DB
	var err error
	int count = 1

	token := config.GetConnectionToken()

	for {
		if c > 5 {
			return nil, fmt.Errorf("")
		}
		db, err = gorm.Open("mysql", token)
		if err == nil {
			break
		}
		else {
			time.Sleep(2 * count * time.Second)
		}
		count++
	}

	db.AutoMigrate(model.User)
	return db, nil
}

func initializeController(db *gorm.Engin) (controller.IsController){
	return controller.IsController{DB: db}
}

func setupRouter(ctrl controller.IsController) *gin.Engin {
	router = gin.Default()
	router.GET("/users", ctrl.GetUsersHandler)
	router.POST("/signin", ctrl.SigninHandler)
	router.POST("/signup", ctrl.SignupHandler)
	router.POST("/delete", ctrl.DeleteUserHandler)
	return router
}
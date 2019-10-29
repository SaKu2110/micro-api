package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
	"github.com/SaKu2110/micro-api/sign/model"
)

type IsController struct {
	DB	*gorm.Engin
}

func (ctrl *IsController)GetUsersHandler(context *gin.Context) {
	var	user model.User

	ctrl.DB.Find(&user)
	g.JSON(200, gin.H{"access_token": ""})
}

func (ctrl *IsController)SigninHandler(context *gin.Context) {
	var user, data model.User

	err := context.BindJSON(&user)
	if err != nil {
	}
	ctrl.DB.Find(&data, "name=? and email=? and password=?", user.NAME, user.EMAIL, user.PASSWORD)
	if data.NAME != user.NAME {
		g.JSON(201, gin.H{"msg": ""})
	}
	if data.EMAIL != user.EMAIL {
		g.JSON(201, gin.H{"msg": ""})
	}
	if data.PASSWORD != user.PASSWORD {
		g.JSON(201, gin.H{"msg": ""})
	}

	g.JSON(200, gin.H{"msg": ""})
}

func (ctrl *IsController)SignupHandler(context *gin.Context) {
	var user, data model.User

	err := context.BindJSON(&user)
	if err != nil {
	}
	ctrl.DB.Find(&data, "email=?", user.EMAIL)
	if data.EMAIL == user.EMAIL {
		g.JSON(201, gin.H{"msg": ""})
	}

	g.JSON(200, gin.H{"msg": ""})
}

func (ctrl *IsController)DeleteUserHandler(context *gin.Context) {
	var user, data model.User

	err := context.BindJSON(&user)
	if err != nil {
	}
	if user.ID == "" {
		g.JSON(201, gin.H{"msg": ""})
	}

	ctrl.DB.Find(&data, "id=?", user.ID)
	ctrl.DB.Delete(&data)

	g.JSON(200, gin.H{"msg": ""})
}
package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/micro-api/sign/model"
)

type IsController struct {
	DB	*gorm.DB
}

type userList struct {
	users	[]model.User
}

func (ctrl *IsController)GetUsersHandler(context *gin.Context) {
	var	user []model.User
	ctrl.DB.Find(&user)

	context.JSON(200, gin.H{"users": user})
}

func (ctrl *IsController)SigninHandler(context *gin.Context) {
	var user, data model.User

	err := context.BindJSON(&user)
	if err != nil {
	}
	ctrl.DB.Find(&data, "name=? and email=? and password=?", user.NAME, user.EMAIL, user.PASSWORD)
	if data.NAME != user.NAME {
		context.JSON(201, gin.H{"msg": ""})
		return
	}
	if data.EMAIL != user.EMAIL {
		context.JSON(201, gin.H{"msg": ""})
		return
	}
	if data.PASSWORD != user.PASSWORD {
		context.JSON(201, gin.H{"msg": ""})
		return
	}

	context.JSON(200, gin.H{"msg": ""})
}

func (ctrl *IsController)SignupHandler(context *gin.Context) {
	var user, data model.User

	err := context.BindJSON(&user)
	if err != nil {
	}
	ctrl.DB.Find(&data, "email=?", user.EMAIL)
	if data.EMAIL == user.EMAIL {
		context.JSON(201, gin.H{"msg": ""})
		return
	}
	ctrl.DB.Create(&user)
	context.JSON(200, gin.H{"msg": ""})
}

func (ctrl *IsController)DeleteUserHandler(context *gin.Context) {
	var user, data model.User

	err := context.BindJSON(&user)
	if err != nil {
	}
	if user.ID == "" {
		context.JSON(201, gin.H{"msg": ""})
		return
	}

	ctrl.DB.Find(&data, "id=?", user.ID)
	ctrl.DB.Delete(&data)

	context.JSON(200, gin.H{"msg": ""})
}
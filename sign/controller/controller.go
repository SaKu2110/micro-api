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
}

func (ctrl *IsController)SigninHandler(context *gin.Context) {
}

func (ctrl *IsController)SignupHandler(context *gin.Context) {
}

func (ctrl *IsController)DeleteUserHandler(context *gin.Context) {
}
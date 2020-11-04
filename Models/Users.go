package Models

import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/DeniesKresna/gocms/Helpers"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password     string `json:"-"`
	RoleId   uint `json:"role_id"`
	Email string `json:"email"`
}

type UserCreate struct{
	Name string `form:"name" validate:"required"`
	Password string `form:"password" validate:"required"`
	Email string `form:"email" validate:"required,email"`
	RoleId uint `form:"role_id" validate:"required"`
}

type UserUpdate struct{
	Name string `json:"name" binding:"required"`
}

var v *validator.Validate

func ValidateUserCreate(c *gin.Context, input *UserCreate) error{
	generatePassword, _ := Api.HashPassword(c.PostForm("password"))
	roleid, _ := strconv.Atoi(c.PostForm("role_id"))
	*input = UserCreate{
		Name: c.PostForm("name"), 
		Password: generatePassword, 
		RoleId: uint(roleid),
		Email: c.PostForm("email"),
	}
	v = validator.New()
	return v.Struct(input)
}

/*
func GetAllBook(b *[]Book) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBook(b *Book) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Book, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Book, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteBook(b *Book, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
*/
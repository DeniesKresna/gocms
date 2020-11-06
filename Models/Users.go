package Models

import (
	"github.com/DeniesKresna/gocms/Helpers"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	Name string `json:"name" form:"name" validate:"required"`
}

func ValidateUserCreate(c *gin.Context, input *UserCreate) error{
	generatePassword, _ := Api.HashPassword(c.PostForm("password"))
	if err := c.MustBindWith(input, binding.FormMultipart); err != nil{
		return err
	}
	v = validator.New()
	if err := v.Struct(input); err != nil{
		return err
	}
	input.Password = generatePassword
	return nil
}

func ValidateUserUpdate(c *gin.Context, input *UserUpdate) error{
	if err := c.Bind(input); err != nil{
		return err
	}
	v = validator.New()
	return v.Struct(input)
}
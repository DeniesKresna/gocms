package Models

import (
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Role struct {
	gorm.Model
	Name string `json: "name"`
}

type RoleCreate struct {
	Name string `form: "name" validate:"required"`
}

type RoleUpdate struct {
	Name string `json: "name"`
}

func ValidateRoleCreate(c *gin.Context, input *RoleCreate) error{
	if err := c.BindWith(input, binding.FormMultipart); err != nil{
		return err
	}
	v = validator.New()
	return v.Struct(input)
}

func ValidateRoleUpdate(c *gin.Context, input *RoleUpdate) error{
	if err:= c.ShouldBindJSON(&input); err != nil{
		return err
	}
	v = validator.New()
	return v.Struct(&input)
}
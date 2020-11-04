package Controllers

import (
	"github.com/DeniesKresna/gocms/Helpers"
	"github.com/DeniesKresna/gocms/Models"
	"github.com/DeniesKresna/gocms/Config"
	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var users []Models.UserShow
	err := Config.DB.Find(&users).Error
	if err != nil {
		Api.RespondJSON(c, 404, users)
	} else {
		Api.RespondJSON(c, 200, users)
	}
}

func UserStore(c *gin.Context) {
	var input Models.UserCreate
	if err:= Models.ValidateUserCreate(c, &input); err != nil{
		Api.RespondJSON(c, 422, err)
	}

	generatePassword, _ := Api.HashPassword(input.Password)
	var user = Models.UserCreate{
		Name: input.Name, 
		Password: generatePassword, 
		RoleId: input.RoleId,
		Email: input.Email,
	}
	err := Config.DB.Create(&user)
	if err != nil {
		Api.RespondJSON(c, 404, user)
	} else {
		Api.RespondJSON(c, 200, user)
	}
}
/*
func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models.PutOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}
*/
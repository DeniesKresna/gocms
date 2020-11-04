package Controllers

import (
	"github.com/DeniesKresna/gocms/Helpers"
	"github.com/DeniesKresna/gocms/Models"
	"github.com/DeniesKresna/gocms/Config"
	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var datas []Models.User
	err := Config.DB.Find(&datas).Error
	if err != nil {
		Api.RespondJSON(c, 404, datas)
	} else {
		Api.RespondJSON(c, 200, datas)
	}
}

func UserStore(c *gin.Context) {
	var input Models.UserCreate
	if err := Models.ValidateUserCreate(c, &input,); err != nil{
		Api.RespondJSON(c, 422, err.Error())
    	return
	}

	var data = Models.User{
		Name: input.Name, 
		Password: input.Password, 
		RoleId: input.RoleId,
		Email: input.Email,
	}

	err := Config.DB.Create(&data).Error
	if err != nil {
		Api.RespondJSON(c, 404, "unknown error")
	} else {
		Api.RespondJSON(c, 200, data)
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
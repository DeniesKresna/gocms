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

func UserShow(c *gin.Context){
	var data Models.User
	err := Config.DB.First(&data, c.Param("id")).Error
	if err != nil {
		Api.RespondJSON(c, 404, "no data")
	}else{
		Api.RespondJSON(c, 200, data)
	}
}

func UserUpdate(c *gin.Context){
	var input Models.UserUpdate
	if err := Models.ValidateUserUpdate(c, &input); err != nil{
		Api.RespondJSON(c, 422, err.Error())
	}

	var data Models.User

	if err := Config.DB.First(&data, c.Param("id")).Error; err != nil{
		Api.RespondJSON(c, 404, "no data")
		return
	}

	err := Config.DB.Model(&data).Updates(input).Error
	if err != nil{
		Api.RespondJSON(c, 404, "unknown error")
	} else {
		Api.RespondJSON(c, 200, data)
	}
}

func UserDestroy(c *gin.Context){
	var data Models.User
	if err := Config.DB.First(&data, c.Param("id")).Error; err != nil{
		Api.RespondJSON(c, 404, "no data")
		return
	}

	name := data.Name
	if err := Config.DB.Delete(&data).Error; err != nil{
		Api.RespondJSON(c, 404, "unknown error")
	}else{
		Api.RespondJSON(c, 200, "success delete " + name)
	}
}

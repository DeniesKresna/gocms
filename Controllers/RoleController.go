package Controllers

import (
	"github.com/DeniesKresna/gocms/Helpers"
	"github.com/DeniesKresna/gocms/Models"
	"github.com/DeniesKresna/gocms/Config"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

func RoleIndex(c *gin.Context){
	var datas Models.Role
	if err := Config.DB.Find(&datas).Error; err != nil{
		Api.RespondJSON(c, 404, err)
	}else{
		Api.RespondJSON(c, 200, datas)
	}
}

func RoleStore(c *gin.Context){
	var input Models.RoleCreate

	if err := Models.ValidateRoleCreate(c, &input); err != nil{
		Api.RespondJSON(c, 422, err)
		return
	}

	var jsoninput,_ = json.Marshal(input)
	var bytejsoninput = []byte(jsoninput)

	var data Models.Role
	if err := json.Unmarshal(bytejsoninput, &data); err != nil{
		Api.RespondJSON(c, 404, "unknown error")
		return
	}

	if err := Config.DB.Create(&data).Error; err != nil{
		Api.RespondJSON(c, 404, "unknown error")
	}else{
		Api.RespondJSON(c, 200, data)
	}
}
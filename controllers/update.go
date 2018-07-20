package controllers

import (
	"github.com/gin-gonic/gin"
	"web/info"
	"net/http"
	"web/api"
)

func Update(c *gin.Context)  {
	var form   update

	ErrJson := new(retJson)

	if err := c.ShouldBind(&form); err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoArg
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}

	err := api.Update(form.Address,form.Address,form.NewPassword)

	if err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoPassword
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}

	ErrJson.Error = 0
	ErrJson.Info = info.UpdateOk
	c.JSON(http.StatusOK,&ErrJson)


}

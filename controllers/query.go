package controllers

import (
	"github.com/gin-gonic/gin"
	"web/info"
	"net/http"
	eth "web/api"
	"encoding/json"
)

func Query(c *gin.Context)  {
	var form  AddressForm
	ErrJson := new(retJson)

	if err := c.ShouldBind(&form); err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoArg
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}
	response,err := eth.QueryTransactions(form.Address)

	RetJson,err := checkerr(response,err)

	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}
	ret, err := json.Marshal(response)
	RetJson,err = checkerr(ret,err)

	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}
	RetJson.Error = 0
	RetJson.Info = string(ret)
	c.JSON(http.StatusOK,&RetJson)


}

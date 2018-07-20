package controllers

import (
	"github.com/gin-gonic/gin"
	"web/info"
	"net/http"
	"fmt"
	eth "web/api"
)

func GetBalance(c *gin.Context) {
	var form AddressForm

	ErrJson := new(retJson)

	if err := c.ShouldBind(&form); err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoArg
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}

	response,err := eth.EthGetBalance(form.Address,"latest")

	RetJson,err := checkerr(response,err)

	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}


	num := fmt.Sprint(response)

	RetJson.Error = 0
	RetJson.Info = num
	c.JSON(http.StatusOK,&RetJson)

}

package controllers

import (
	"github.com/gin-gonic/gin"
	eth "web/api"
	"net/http"
)

func GetPrice(c *gin.Context){

	response,err := eth.Ethprice()
	RetJson,err := checkerr(response,err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}
	RetJson.Error = 0
	RetJson.Info = response.Ethusd
	c.JSON(http.StatusOK,&RetJson)

}


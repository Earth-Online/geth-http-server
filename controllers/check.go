package controllers

import (
	"github.com/gin-gonic/gin"
	"web/info"
	"net/http"
	eth "web/api"
	"log"
)

func Check(c *gin.Context){
	var form   check

	ErrJson := new(retJson)
	log.Print(form)
	if err := c.ShouldBind(&form); err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoArg
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}

	log.SetOutput(gin.DefaultWriter) // You may need this



	response,err := eth.EthGetTransactionByHash(form.Hash)
	log.Print(err)

	RetJson,err := checkerr(response,err)

	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}

	if response.To != form.To || response.Input != form.Info || response.Value != form.Value {
		RetJson.Error = 1
		RetJson.Info = info.ErrNoCheck
		c.JSON(http.StatusOK,&RetJson)
		return
	}

	RetJson.Error = 0
	RetJson.Info = info.CheckOk
	c.JSON(http.StatusOK,&RetJson)

}

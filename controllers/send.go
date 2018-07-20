package controllers

import (
	"github.com/gin-gonic/gin"
	"web/info"
	"net/http"
	eth "github.com/FactomProject/EthereumAPI"
	"log"
)


func Send(c *gin.Context){
	// 提供简化接口
	var form   send

	ErrJson := new(retJson)
	log.SetOutput(gin.DefaultErrorWriter)

	if err := c.ShouldBind(&form); err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoArg
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}

	transaction := eth.TransactionObject{
		From:form.From,
		To:form.To,
		Value:form.Value,
		Gas:form.Gas,
		Input:form.Input,
		GasPrice:form.GasPrice,
	}

	response, err :=  eth.PersonalSendTransaction(&transaction,form.Password)
	RetJson,err := checkerr(response,err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}

	RetJson.Error = 0
	RetJson.Info = response
	c.JSON(http.StatusOK,&RetJson)




}

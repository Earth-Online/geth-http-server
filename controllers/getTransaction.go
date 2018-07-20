package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/info"
	eth "github.com/FactomProject/EthereumAPI"
	"encoding/json"
)

func GetTransaction(c *gin.Context)  {
	var RetJson  = new(retJson)
	var form HashForm
	if err := c.ShouldBind(&form); err != nil {
			RetJson.Error = 1
			RetJson.Info  = info.ErrNoArg

			c.JSON(http.StatusOK,&RetJson)
			return
	}

	response,err := eth.EthGetTransactionByHash(form.Hash)
	Json,err := checkerr(response,err)

	if err != nil {
		c.JSON(http.StatusInternalServerError,&Json)
		return
	}

	ret, err := json.Marshal(response)
	Json,err = checkerr(ret,err)

	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}

	RetJson.Error = 0
	RetJson.Info = string(ret)
	c.JSON(http.StatusOK,&RetJson)

}

package controllers

import (
	"github.com/gin-gonic/gin"
	"web/conf"
	"web/info"
	"net/http"
	eth "github.com/FactomProject/EthereumAPI"
	"log"
)

func init()  {
	eth.SetServer(conf.Rpc.Key("RpcHost").String())
	log.Print(conf.Rpc.Key("RpcHost").String())
	// client = jsonrpc.NewClient(conf.Rpc.Key("RpcHost").String())
}

func Auth(c *gin.Context)   {
	apikey := c.GetHeader("apikey")
	if apikey != conf.Api.Key("apikey").String(){
		ErrJson := new(retJson)
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoApiKey
		c.JSON(http.StatusForbidden,&ErrJson)
		c.Abort()
	}
	}

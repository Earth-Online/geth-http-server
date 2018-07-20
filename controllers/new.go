package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/info"
	eth "web/api"
)

// var client jsonrpc.RPCClient



func NewUser(c *gin.Context)  {
	var form  newAccount
	ErrJson := new(retJson)

	if err := c.ShouldBind(&form); err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoArg
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}




	key,address, err :=  eth.NewKey(form.Password)
	if err != nil {
		ErrJson.Error =1
		ErrJson.Info = info.ErrNewKey
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}
	/* RetJson,err := checkerr(response,err)
	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}


	address := fmt.Sprint(response)
	user := model.User{From:address,Password:form.Password}
	model.DB.Create(user)

*/
	RetJson := new(NewKey)
	RetJson.Error = 0
	RetJson.PrivateKey = key
	RetJson.Address = address
	c.JSON(http.StatusOK,&RetJson)

}

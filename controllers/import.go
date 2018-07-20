package controllers

import (
	"web/info"
	"net/http"
	"github.com/gin-gonic/gin"
	eth "github.com/FactomProject/EthereumAPI"
	"fmt"
)

func ImportKey(c *gin.Context)  {
	var form   importKey

	ErrJson := new(retJson)

	if err := c.ShouldBind(&form); err != nil {
		ErrJson.Error = 1
		ErrJson.Info = info.ErrNoArg
		c.JSON(http.StatusInternalServerError,&ErrJson)
		return
	}

	response,err := eth.PersonalImportRawKey(form.Keydata,form.Password)

	RetJson,err := checkerr(response,err)

	if err != nil {
		c.JSON(http.StatusInternalServerError,&RetJson)
		return
	}
	RetJson.Error = 0
	RetJson.Info = fmt.Sprint(response)
	c.JSON(http.StatusOK,&RetJson)


}
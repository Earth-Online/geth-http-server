package controllers

import (
	"fmt"

)

func checkerr(result interface{},err error) (retJson,error)  {
	var RetJson retJson

	if err != nil {
		RetJson.Error = 1
		RetJson.Info = fmt.Sprintf("系统出现未知错误  错误信息%s",err)
		return RetJson,err

	}
	return RetJson,nil
}

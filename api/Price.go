package api

import "fmt"

func Ethprice() (*PriceObject,error){
	arg := new(map[string]string)
	resp, err :=  Call(StatsModule,"ethprice",*arg)

	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(PriceObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil


}

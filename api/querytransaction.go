package api

func QueryTransactions(address string) (*[]TransactionInfo,error) {
	arg := map[string]string{
		"address":address,
		"sort":"asc",

	}
	resp, err :=  Call(AccountModule,"txlist",arg)
	if err != nil {
		return nil,err
	}

	answer := new([]TransactionInfo)
	if err != nil {
		return nil,err
	}
	err = MapToObject(resp.Result, answer)


	return answer,nil

}

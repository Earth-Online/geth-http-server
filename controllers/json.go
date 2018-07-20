package controllers

type retJson struct {
	Error int `json:"error"`
	Info string `json:"info"`
}

type newAccount struct {
	Password string `form:"Password" binding:"required"`
}

type AddressForm struct {
	Address string `form:"address" binding:"required"`
}

type check struct {
	Hash string `form:"hash" binding:"required"`
	Info string `form:"info" binding:"required"`
	To string `form:"address" binding:"required"`
	Value string `form:"value" binding:"required"`
}

type HashForm struct {
	Hash string `form:"hash" binding:"required"`
}

type send struct {
	From     string `form:"from" binding:"required"`
	To       string `form:"to"`
	Gas      string `form:"gas"`
	GasPrice string `form:"gasPrice"`
	Value    string `form:"value"`
	Input    string `form:"input"`
	Password    string `form:"Password" binding:"required"`
}

type importKey struct {
	Keydata string  `form:"keydata" binding:"required"`
	Password string `form:"Password" binding:"required"`
}

type update struct {
	Address string `form:"From" binding:"required"`
	Password string `form:"Password" binding:"required"`
	NewPassword string `form:"NewPassword" binding:"required"`
}

type NewKey struct {
	Error int `json:"error"`
	PrivateKey string `json:"PrivateKey"`
	Address string `json:"From"`
}







package conf

import (
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
	Rpc *ini.Section
	Mysql *ini.Section
	Api *ini.Section
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		panic("Fail to parse 'app.ini'")
	}
	Rpc, err = Cfg.GetSection("rpc")
	if err != nil {
		panic(err)
	}

	Mysql, err = Cfg.GetSection("mysql")
	if err != nil {
		panic(err)
	}


	Api, err = Cfg.GetSection("api")
	if err != nil {
		panic(err)
	}

}

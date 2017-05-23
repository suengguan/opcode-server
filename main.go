package main

import (
	_ "opcode-server/routers"

	appApi "api/app_service"
	"github.com/astaxie/beego"
)

func main() {
	var cfg = beego.AppConfig
	appApi.AccountApi.Init(cfg.String("AccountService"))
	appApi.AlgorithmApi.Init(cfg.String("AlgorithmService"))
	appApi.BussinessApi.Init(cfg.String("BussinessService"))
	appApi.DataApi.Init(cfg.String("DataService"))
	appApi.LogApi.Init(cfg.String("LogService"))
	appApi.LoginApi.Init(cfg.String("LoginService"))
	appApi.StatusApi.Init(cfg.String("StatusService"))
	appApi.SummaryApi.Init(cfg.String("SummaryService"))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

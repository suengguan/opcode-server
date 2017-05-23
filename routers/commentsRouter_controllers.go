package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["opcode-server/controllers:DispatchController"] = append(beego.GlobalControllerRouter["opcode-server/controllers:DispatchController"],
		beego.ControllerComments{
			Method: "Run",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}

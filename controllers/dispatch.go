package controllers

import (
	"model"

	"opcode-server/service"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"

	"encoding/json"
)

// Operations about Dispatch
type DispatchController struct {
	beego.Controller
}

// @Title Run
// @Description Dispatch service
// @Param	body		body 	models.MessageRequest	true		"body for MessageRequest content"
// @Success 200 {object} models.MessageResponse
// @Failure 403 body is empty
// @router / [post]
func (this *DispatchController) Run() {
	var err error
	var msgRequest model.MessageRequest
	var msgRespond model.MessageResponse

	// unmarshal data
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &msgRequest)
	msgRespond.Token = msgRequest.Token
	msgRespond.SessionId = msgRequest.SessionId
	msgRespond.Target = msgRequest.Parameter.Target
	if err == nil {
		// run
		var result []byte
		var svc service.DispatchService
		result, err = svc.Run(&msgRequest)
		if err == nil {
			if msgRequest.Parameter.Action == model.MSG_PARAM_ACTION_LOGIN {
				msgRespond.Token = uuid.NewV4().String()
				msgRespond.SessionId = uuid.NewV4().String()
			}
			msgRespond.ResultCode = model.MSG_RESULTCODE_SUCCESS
			msgRespond.Reason = "success"
			msgRespond.Result = string(result)
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		msgRespond.ResultCode = model.MSG_RESULTCODE_FAILED
		msgRespond.Reason = err.Error()
	}

	this.Data["json"] = &msgRespond

	this.ServeJSON()
}

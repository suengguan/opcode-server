package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"model"
	_ "opcode-server/routers"
)

const (
	base_url = "http://localhost:8080/v1/dispatch"
)

func Test_Dispatch_Run(t *testing.T) {
	// request message
	var reqMsg model.MessageRequest

	reqMsg.Token = ""
	reqMsg.SessionId = ""

	var from model.MessageRole
	from.Id = 1
	from.Type = model.MSG_ROLE_TYPE_USER_PC_APP
	reqMsg.From = &from

	var to model.MessageRole
	to.Id = model.LOGIN_SERVICE
	to.Type = model.MSG_ROLE_TYPE_SERVICE
	reqMsg.To = &to

	var parameter model.MessageParameter
	parameter.Action = model.MSG_PARAM_ACTION_LOGIN
	parameter.Target = model.MSG_PARAM_TYPE_USER

	var parameterData model.MessageParameterData
	parameterData.Type = model.MSG_PARAM_TYPE_USER

	var admin model.User
	admin.Name = "admin"
	admin.EncryptedPassword = "admin"
	adminData, err := json.Marshal(&admin)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	parameterData.Content = string(adminData)
	parameter.Data = append(parameter.Data, &parameterData)
	reqMsg.Parameter = &parameter

	requestData, err := json.Marshal(&reqMsg)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// post
	res, err := http.Post(base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

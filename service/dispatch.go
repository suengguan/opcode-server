package service

import (
	"encoding/json"
	"fmt"
	"model"
	"time"

	appApi "api/app_service"
	"github.com/astaxie/beego"
)

type DispatchService struct {
}

func (this *DispatchService) Run(msgReq *model.MessageRequest) ([]byte, error) {
	var err error
	var result []byte

	//beego.Debug("->to id:", msgReq.To.Id)
	//beego.Debug("->action:", msgReq.Parameter.Action)
	//beego.Debug("->target:", msgReq.Parameter.Target)
	var data *model.MessageParameterData

	if msgReq.To.Id == model.ACCOUNT_SERVICE {
		beego.Debug("->dispatch to ACCOUNT_SERVICE")
		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_REGISTER {
			beego.Debug("->REGISTER")
			data, _ = this.getTargetData(model.MSG_PARAM_TYPE_USER, msgReq.Parameter.Data)

			this.writeLog(msgReq, "register user:"+data.Content)
			beego.Debug("->ApiRegisterAccount")
			result, err = appApi.AccountApi.Register(([]byte)(data.Content))
		} else if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_GET {
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_USER_LIST {
				beego.Debug("->ApiGetAllAccount")
				result, err = appApi.AccountApi.GetAll()
			}
		} else if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_UPDATE {
			beego.Debug("->UPDATE")
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_USER {
				beego.Debug("->USER")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_USER, msgReq.Parameter.Data)

				this.writeLog(msgReq, "update user:"+data.Content)
				beego.Debug("->ApiUpdateAccount")
				result, err = appApi.AccountApi.Update(([]byte)(data.Content))
			}
		} else if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_DELETE {
			beego.Debug("->DELETE")
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_USER_LIST {
				beego.Debug("->USER_LIST")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_USER_LIST, msgReq.Parameter.Data)

				this.writeLog(msgReq, "delete users:"+data.Content)
				beego.Debug("->ApiDeleteAccounts")
				result, err = appApi.AccountApi.Delete(([]byte)(data.Content))
			}
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else if msgReq.To.Id == model.ALGORITHM_SERVICE {
		beego.Debug("->dispatch to ALGORITHM_SERVICE")
		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_CREATE {
			beego.Debug("->CREATE")
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_ALGORITHM {
				beego.Debug("->ALGORITHM")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_ALGORITHM, msgReq.Parameter.Data)

				this.writeLog(msgReq, "create algorithm:"+data.Content)
				beego.Debug("->ApiCreateAlgorithm")
				result, err = appApi.AlgorithmApi.Create(([]byte)(data.Content))
			}
		} else if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_DELETE {
			beego.Debug("->DELETE")
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_ALGORITHM {
				beego.Debug("->ALGORITHM")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_ALGORITHM, msgReq.Parameter.Data)

				this.writeLog(msgReq, "delete algorithms:"+data.Content)
				beego.Debug("->ApiDeleteAlgorithm")
				err = appApi.AlgorithmApi.Delete(([]byte)(data.Content))
			}
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else if msgReq.To.Id == model.BUSSINESS_SERVICE {
		beego.Debug("->dispatch to BUSSINESS_SERVICE")

		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_GET {
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_CURRENT_POD {
				beego.Debug("->ApiGetCurrentPods")
				result, err = appApi.BussinessApi.GetCurrentPods(msgReq.From.Id)
			}
		} else if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_CREATE {
			beego.Debug("->CREATE")
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_PROJECT {
				beego.Debug("->PROJECT")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_PROJECT, msgReq.Parameter.Data)

				this.writeLog(msgReq, "create project:"+data.Content)
				beego.Debug("->ApiCreateProject")
				result, err = appApi.BussinessApi.CreateProject(([]byte)(data.Content))
			} else if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_JOB {
				beego.Debug("->JOB")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_JOB, msgReq.Parameter.Data)

				this.writeLog(msgReq, "create job:"+data.Content)
				beego.Debug("->ApiCreateJob")
				result, err = appApi.BussinessApi.CreateJob(([]byte)(data.Content))
			}
		} else if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_UPDATE {
			beego.Debug("->UPDATE")
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_JOB_UPDATE {
				beego.Debug("->JOB")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_JOB_UPDATE, msgReq.Parameter.Data)

				this.writeLog(msgReq, "update job:"+data.Content)
				beego.Debug("->ApiUpdateJob")
				result, err = appApi.BussinessApi.UpdateJob(([]byte)(data.Content))
			}
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else if msgReq.To.Id == model.DATA_SERVICE {
		beego.Debug("->dispatch to DATA_SERVICE")
		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_GET {
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_INPUT_FILES {
				beego.Debug("->ApiGetInputFiles")
				result, err = appApi.DataApi.GetInputFiles(msgReq.From.Id)
			}
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else if msgReq.To.Id == model.LOG_SERVICE {
		beego.Debug("->dispatch to LOG_SERVICE")

		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_GET {
			beego.Debug("->GET")
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_POD_LOG {
				beego.Debug("->POD_LOG")
				data, _ = this.getTargetData(model.MSG_PARAM_TYPE_POD_LOG, msgReq.Parameter.Data)

				var pod model.Pod
				err = json.Unmarshal(([]byte)(data.Content), &pod)
				if err == nil {
					beego.Debug("->ApiGetPodLog")
					result, err = appApi.LogApi.GetPodLog(pod.Id)
				}
			} else if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_ACTION {
				beego.Debug("->ACTION")
				beego.Debug("->ApiGetAllActions")
				result, err = appApi.LogApi.GetAllActions(msgReq.From.Id)
			} else {
				err = fmt.Errorf("%d%s", msgReq.Parameter.Target, " is invalid target")
			}
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else if msgReq.To.Id == model.LOGIN_SERVICE {
		beego.Debug("->dispatch to LOGIN_SERVICE")
		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_LOGIN {
			beego.Debug("->LOGIN")
			data, _ = this.getTargetData(model.MSG_PARAM_TYPE_USER, msgReq.Parameter.Data)

			beego.Debug("->ApiLogin")
			result, err = appApi.LoginApi.Login(([]byte)(data.Content))
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else if msgReq.To.Id == model.STATUS_SERVICE {
		beego.Debug("->dispatch to STATUS_SERVICE")

		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_GET {
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_STATUS {
				beego.Debug("->ApiGetAllJobStatus")
				result, err = appApi.StatusApi.GetAllJobStatus(msgReq.From.Id)
			}
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else if msgReq.To.Id == model.SUMMARY_SERVICE {
		beego.Debug("->dispatch to SUMMARY_SERVICE")

		if msgReq.Parameter.Action == model.MSG_PARAM_ACTION_GET {
			if msgReq.Parameter.Target == model.MSG_PARAM_TYPE_SUMMARY {
				beego.Debug("->ApiGetSummary")
				result, err = appApi.SummaryApi.Get(msgReq.From.Id)
			}
		} else {
			err = fmt.Errorf("%d%s", msgReq.Parameter.Action, " is invalid action type")
		}
	} else {
		err = fmt.Errorf("%d%s", msgReq.To.Id, " is invalid service id")
	}

	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	beego.Debug("result:", string(result))

	return result, err
}

func (this *DispatchService) getTargetData(target int, dataList []*model.MessageParameterData) (*model.MessageParameterData, bool) {
	dataCnt := len(dataList)
	for i := 0; i < dataCnt; i++ {
		if dataList[i].Type == target {
			return dataList[i], true
		}
	}

	return nil, false
}

func (this *DispatchService) writeLog(msgReq *model.MessageRequest, content string) error {
	var action model.Action
	var err error

	action.Time = time.Now().Unix()
	action.SessionId = msgReq.SessionId
	var u model.User
	u.Id = msgReq.From.Id
	action.User = &u
	action.DevType = msgReq.From.Type
	action.Type = model.LOG_TYPE_INFO

	action.Content = content

	// write
	var data []byte
	data, err = json.Marshal(&action)
	if err != nil {
		beego.Debug("json Unmarshal data failed")
		return err
	}

	err = appApi.LogApi.WriteAction(data)
	if err != nil {
		return err
	}

	return err
}

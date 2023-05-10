package lorawan

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/emq-auth-http/internal/db/model"
)

type OtaaResp struct {
	app.Response
	AppKey string `json:"appKey,omitempty"`
	DevEui string `json:"devEui,omitempty"`
}

// getOTAAKey
// 获取key
func getOTAAKey(ctx iris.Context) {
	var err error
	var resp OtaaResp

	var req GenOtaaReq
	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.DevEui) == 0 {
		resp.Code = tool.RespCodeSuccess
		tool.ResponseJSON(ctx, resp)
		return
	}

	device, err := getDevice(req.DevEui)
	if err != nil {
		resp.Code = tool.RespCodeSuccess
		resp.Message = "unknown device"
		tool.ResponseJSON(ctx, resp)
		return
	}
	var key = tool.GetRandomString(16)
	device.OtaaKey = key
	resp.AppKey = key

	tool.ResponseJSON(ctx, resp)
}

type GenOtaaReq struct {
	DevEui string `json:"devEui"`
}

type GenOtaaResponse struct {
	app.Response
	AppKey string `json:"appKey,omitempty"`
	DevEui string `json:"devEui,omitempty"`
}

// otaaKey
// 输入为DevEUI(ID),输出此ID对应的AppKey,
// 格式符合LoraWAN规范: 16byte(hex)
func genOTAAKey(ctx iris.Context) {
	var err error
	var resp GenOtaaResponse

	var req GenOtaaReq
	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.DevEui) == 0 {
		resp.Code = tool.RespCodeSuccess
		tool.ResponseJSON(ctx, resp)
		return
	}

	device, err := getDevice(req.DevEui)
	if err != nil {
		resp.Code = tool.RespCodeSuccess
		resp.Message = "unknown device"
		tool.ResponseJSON(ctx, resp)
		return
	}
	var key = tool.GetRandomString(16)
	device.OtaaKey = key
	resp.AppKey = key

	tool.ResponseJSON(ctx, resp)
}

func getDevice(id string) (device model.LoraDeviceSecretSecret, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.LoraDeviceSecretSecret)).
		Filter("DevEui", id).
		One(&device)
	return
}

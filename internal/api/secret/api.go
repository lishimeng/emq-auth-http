package secret

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/emq-auth-http/internal/passwd"
	"github.com/lishimeng/go-log"
)

type Req struct {
	Sn       string `json:"sn,omitempty"`
	Password string `json:"password,omitempty"`
}

type Resp struct {
	app.Response
	Sn       string `json:"sn,omitempty"`
	Password string `json:"password,omitempty"`
}

func add(ctx iris.Context) {
	var err error
	var req Req
	var resp Resp
	err = ctx.ReadJSON(&req)
	if err != nil {

		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Sn) == 0 {

		resp.Code = tool.RespCodeSuccess
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Password) == 0 {
		log.Debug("password nil, use random")
		req.Password = tool.GetRandomString(8)
	}

	d, err := createDevice(req.Sn)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	// 此处必须重新在数据库中取一次, 加密需要使用时间作为盐值
	d, err = getDeviceById(d.Id)
	r := passwd.Generate(req.Password, d)

	err = updatePassword(d, r)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Password = req.Password
	tool.ResponseJSON(ctx, resp)
}

type DeviceSecretResp struct {
	app.Response
	Sn string `json:"sn,omitempty"`
}

func get(ctx iris.Context) {
	var resp DeviceSecretResp
	sn := ctx.Params().Get("sn")
	u, err := getDevice(sn)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Sn = u.Username
	tool.ResponseJSON(ctx, resp)
}

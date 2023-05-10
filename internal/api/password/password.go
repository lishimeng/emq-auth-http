package password

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/emq-auth-http/internal/db/model"
	"github.com/lishimeng/emq-auth-http/internal/passwd"
	"github.com/lishimeng/go-log"
)

type Req struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty" json:"password,omitempty"`
}

type Resp struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty" json:"password,omitempty"`
	PasswordEncode string `json:"password_encode,omitempty" json:"password_encode,omitempty"`
}

// gen
// 配置密码. 账号必须存在才能创建密码
func gen(ctx iris.Context) {

	log.Info("gen password")

	var req Req
	var resp Resp
	var err error
	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Debug("read json failed")
		log.Debug(err)
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Username) == 0 {
		log.Debug("username nil")
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Username = req.Username

	if len(req.Password) == 0 {
		log.Debug("password nil, use random")
		req.Password = tool.GetRandomString(8)
	}

	u, err := getUser(req.Username)

	if err != nil {
		log.Info(err)
		tool.ResponseJSON(ctx, resp)
		return
	}

	r := passwd.Generate(req.Password, u)

	resp.Password = req.Password
	resp.PasswordEncode = r
	tool.ResponseJSON(ctx, resp)
}

func getUser(username string) (ds model.DeviceSecret, err error) {

	err = app.GetOrm().Context.QueryTable(new(model.DeviceSecret)).Filter("Username", username).One(&ds)
	return
}

func Router(p iris.Party) {
	p.Post("/", gen)
}

package lorawan

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
)

type OtaaResp struct {
	app.Response
	AppKey string `json:"app_key,omitempty"`
	DevEui string `json:"dev_eui,omitempty"`
}

func otaaKey(ctx iris.Context) {

}

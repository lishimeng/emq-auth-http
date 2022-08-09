package emq

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/emq-auth-http/internal/db/model"
)

func getUser(username string) (ds model.DeviceSecret, err error) {

	err = app.GetOrm().Context.QueryTable(new(model.DeviceSecret)).Filter("Username", username).One(&ds)
	return
}

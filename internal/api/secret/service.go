package secret

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/emq-auth-http/internal/db/model"
	persistence "github.com/lishimeng/go-orm"
)

// createDevice 创建device secret. 如果已经创建就忽略
func createDevice(username string) (d model.DeviceSecret, err error) {
	// select --> insert or update
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		var tmp []model.DeviceSecret
		_, e = ctx.Context.QueryTable(new(model.DeviceSecret)).Filter("Username", username).All(&tmp)
		if e != nil {
			return
		}
		if len(tmp) > 0 { // 已经创建
			d = tmp[0]
			return
		}
		d.Username = username
		_, e = ctx.Context.Insert(&d)
		return
	})
	return
}

func updatePassword(d model.DeviceSecret, passwordSecret string) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		d.Password = passwordSecret
		d.Status = model.DeviceSecretEnable
		_, e = ctx.Context.Update(&d)
		return
	})
	return
}

func getDeviceById(id int) (ds model.DeviceSecret, err error) {

	ds.Id = id
	err = app.GetOrm().Context.Read(&ds)
	return
}

func getDevice(username string) (ds model.DeviceSecret, err error) {

	err = app.GetOrm().Context.QueryTable(new(model.DeviceSecret)).Filter("Username", username).One(&ds)
	return
}

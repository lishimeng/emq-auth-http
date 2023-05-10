package model

import "github.com/lishimeng/app-starter"

type LoraDeviceSecretSecret struct {
	app.Pk
	DevEui  string `orm:"column(dev_eui);unique"`
	OtaaKey string `orm:"column(otaa_key)"`
	app.TableChangeInfo
}

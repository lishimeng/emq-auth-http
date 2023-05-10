package model

import "github.com/lishimeng/app-starter"

type DeviceSecret struct {
	app.Pk
	Username string `orm:"column(username);unique"`
	Password string `orm:"column(password)"`
	app.TableChangeInfo
}

const (
	DeviceSecretDisable = iota // 不可用
	DeviceSecretEnable         // 可用
)

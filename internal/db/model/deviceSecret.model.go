package model

type DeviceSecret struct {
	Pk
	Username string `orm:"column(username);unique"`
	Password string `orm:"column(password)"`
	TableChangeInfo
}

const (
	DeviceSecretDisable = iota // 不可用
	DeviceSecretEnable         // 可用
)

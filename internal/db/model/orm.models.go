package model

func Tables() (t []interface{}) {
	t = append(t,
		new(DeviceSecret),
		new(LoraDeviceSecretSecret),
	)
	return
}

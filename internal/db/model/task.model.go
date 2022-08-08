package model

// MessageTask 消息投送任务
type MessageTask struct {
	Pk
	MessageId         int `orm:"column(message_id)"`          // message
	MessageInstanceId int `orm:"column(message_instance_id)"` // sms id/mail id
	TableChangeInfo
}

const (
	MessageTaskInit        = iota // 新建,初始化
	MessageTaskSending            // 投送中
	MessageTaskSendSuccess        // 投送成功
	MessageTaskSendFailed         // 投送失败
	MessageTaskCancelled   = -1   // 取消
	MessageTaskSendExpired = -9   // 投送失败
)

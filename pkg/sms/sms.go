package sms

import (
	"gohub/pkg/config"
	"sync"
)

// Message 短信的结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// SMS 是发生短信的操作类
type SMS struct {
	Driver Driver
}

// 单例模式
var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Alun{},
		}
	})

	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}

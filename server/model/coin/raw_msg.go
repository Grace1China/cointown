// 自动生成模板RawMsg
package coin

import (
	"time"

	"github.com/Grace1China/cointown/server/global"
)

// 原始消息 结构体  RawMsg
type RawMsg struct {
	global.GVA_MODEL
	CurrentTime *time.Time `json:"currentTime" form:"currentTime" gorm:"column:current_time;comment:;"`           //时间
	ChatId      string     `json:"chatId" form:"chatId" gorm:"column:chat_id;comment:;size:20;"`                  //对话Id
	TopicId     string     `json:"topicId" form:"topicId" gorm:"column:topic_id;comment:;size:20;"`               //话题Id
	MessageText string     `json:"messageText" form:"messageText" gorm:"column:message_text;comment:;size:4000;"` //消息体
	IsNew       bool       `json:"IsNew" form:"IsNew" gorm:"column:is_new;comment:;size:20;"`
}

// TableName 原始消息 RawMsg自定义表名 raw_msg
func (RawMsg) TableName() string {
	return "raw_msg"
}

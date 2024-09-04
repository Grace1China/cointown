// 自动生成模板Channel
package coin

import (
	"time"

	"github.com/Grace1China/cointown/server/global"
)

// 频道 结构体  Channel
type Channel struct {
	global.GVA_MODEL
	CurrentTime     *time.Time `json:"currentTime" form:"currentTime" gorm:"column:current_time;comment:;"`                        //时间
	ChatId          string     `json:"chatId" form:"chatId" gorm:"column:chat_id;comment:;size:20;"`                               //对话Id
	TopicId         string     `json:"topicId" form:"topicId" gorm:"column:topic_id;comment:;size:20;"`                            //话题Id
	LastMessageText string     `json:"lastMessageText" form:"lastMessageText" gorm:"column:last_message_text;comment:;size:4000;"` //消息体
	Template        string     `json:"template" form:"template" gorm:"column:template;comment:;size:4000;"`                        //消息体模版
	Hits            uint       `json:"hits" form:"hits" gorm:"column:hits;comment:;"`
	HitAt           *time.Time `json:"hitAt" form:"hitAt" gorm:"column:hit_at;comment:;size:20;"`
}

// TableName 频道 Channel自定义表名 channel
func (Channel) TableName() string {
	return "channel"
}

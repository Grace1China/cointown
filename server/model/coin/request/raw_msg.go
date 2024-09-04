package request

import (
	"time"

	"github.com/Grace1China/cointown/server/model/common/request"
)

type RawMsgSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ChatId         string     `json:"chatId" form:"chatId"`
	Id             string     `json:"id" form:"id"`

	request.PageInfo
}

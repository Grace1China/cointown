package request

import (
	"time"

	"github.com/Grace1China/cointown/server/model/common/request"
)

type CoinsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

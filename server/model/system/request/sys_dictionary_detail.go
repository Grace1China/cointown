package request

import (
	"github.com/Grace1China/cointown/server/model/common/request"
	"github.com/Grace1China/cointown/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}

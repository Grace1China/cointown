package request

import (
	"github.com/Grace1China/cointown/server/model/common/request"
	"github.com/Grace1China/cointown/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}

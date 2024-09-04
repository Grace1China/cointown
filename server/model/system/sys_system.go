package system

import (
	"github.com/Grace1China/cointown/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}

package response

import "github.com/Grace1China/cointown/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}

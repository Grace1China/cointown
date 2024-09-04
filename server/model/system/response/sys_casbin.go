package response

import (
	"github.com/Grace1China/cointown/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}

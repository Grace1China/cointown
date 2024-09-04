package ast

import (
	"path/filepath"

	"github.com/Grace1China/cointown/server/global"
)

func init() {
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("../../../")
	global.GVA_CONFIG.AutoCode.Server = "server"
}

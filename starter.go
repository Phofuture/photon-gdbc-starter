package gdbcStarter

import (
	"github.com/Phofuture/photon-core-starter/core"
	"github.com/Phofuture/photon-gdbc-starter/gdbc"
)

func init() {
	core.RegisterAddModule(gdbc.Start)
}

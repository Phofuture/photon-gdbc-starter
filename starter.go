package gdbcStarter

import (
	"github.com/Phofuture/photon-gdbc-starter/gdbc"
	"github.com/dennesshen/photon-core-starter/core"
)

func init() {
	core.RegisterAddModule(gdbc.Start)
}

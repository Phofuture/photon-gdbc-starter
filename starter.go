package gdbcStarter

import (
	"github.com/dennesshen/photon-core-starter/core"
	"github.com/dennesshen/photon-gdbc-starter/gdbc"
)

func init() {
	core.RegisterAddModule(gdbc.Start)
}

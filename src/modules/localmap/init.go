package localmap

import (
	"modules/app"

	"github.com/herb-go/util"
)

// ModuleName module name
const ModuleName = "900localmap"

var TileWidth = 8

var TileHeight = 3

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.RegisterDataFolder("data")
		util.InitOrderByName(ModuleName)
		if app.System.TileHeight > 0 {
			TileHeight = app.System.TileHeight
		}
		if app.System.TileWidth > 0 {
			TileWidth = app.System.TileWidth
		}
		MustLoad()
	})
}

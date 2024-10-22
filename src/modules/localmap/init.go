package localmap

import (
	"modules/app"
	"strings"

	"github.com/herb-go/util"
)

// ModuleName module name
const ModuleName = "900localmap"

var TileWidth = 8

var TileHeight = 3

var filters = []string{"│", "[", "]", "↑", "↓", "∨", "∧", "╱", "╲", "─", "┅", "┊", "〓", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

var Replacerfilters *strings.Replacer

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
		linkreplace := []string{}
		for _, v := range filters {
			var newtoken = " "
			if []rune(v)[0] > 255 {
				newtoken = "  "
			}
			linkreplace = append(linkreplace, v, newtoken)
		}
		Replacerfilters = strings.NewReplacer(linkreplace...)

		MustLoad()
	})
}

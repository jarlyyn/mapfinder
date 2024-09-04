package search

import (
	"modules/app"
	"strings"

	"github.com/herb-go/util"
)

// ModuleName module name
const ModuleName = "901search"

var TileMinChinese = 2
var MatchedMinWidth = 12
var MatchedMinHeight = 5
var TrustPercent = 50

var ReplaceToken = []string{"¤"}
var Replacer *strings.Replacer
var Replacer2 = strings.NewReplacer("(", "[", ")", "]")

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
		if app.System.TileMinChinese > 0 {
			TileMinChinese = app.System.TileMinChinese
		}
		replace := []string{}
		for _, v := range ReplaceToken {
			replace = append(replace, v, "●")
		}
		Replacer = strings.NewReplacer(replace...)
	})
}

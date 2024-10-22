package search

import (
	"modules/app"
	"os"
	"strings"

	"github.com/herb-go/util"
)

// ModuleName module name
const ModuleName = "901search"

var TileMinChinese = 2
var MatchedMinSize = 12
var TrustPercent = 30

var ReplaceToken = []string{"¤"}
var Replacer *strings.Replacer
var Replacer2 = strings.NewReplacer("(", "[", ")", "]")
var ChineseMap = map[rune]bool{}

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
		data, err := os.ReadFile(util.System("data", "chinese.txt"))
		if err != nil {
			panic(err)
		}
		allchinese := []rune(string(data))
		for _, v := range allchinese {
			ChineseMap[v] = true
		}
	})
}

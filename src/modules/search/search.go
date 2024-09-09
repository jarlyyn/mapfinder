package search

import (
	"modules/components/textmap"
	"modules/localmap"
	"regexp"
	"sort"
	"strings"
)

type Result struct {
	ID       string
	Name     string
	Room     string
	AreaBox  *Box
	MapBox   *Box
	AreaData string
	MapData  string
}

// 中文正则
var rechinese = regexp.MustCompile("[\u4e00-\u9fa5]")

// 地图常见符号正则
var remapsymbol = regexp.MustCompile("[ │\\[\\]]↑↓∨∧╱╲─┏━┓┃┅〓")

type Box struct {
	Left   int
	Top    int
	Width  int
	Height int
}

func (b *Box) Clone() *Box {
	return &Box{
		Left:   b.Left,
		Top:    b.Top,
		Width:  b.Width,
		Height: b.Height,
	}
}
func (b *Box) Expand(direct int) {
	switch direct {
	case directLeft:
		b.Left = b.Left - 1
		b.Width = b.Width + 1
		break
	case directRight:
		b.Width = b.Width + 1
		break
	case directUp:
		b.Top = b.Top - 1
		b.Height = b.Height + 1
		break
	case directDown:
		b.Height = b.Height + 1
		break
	}
}

func NewBox(left int, top int, width int, height int) *Box {
	return &Box{
		Left:   left,
		Top:    top,
		Width:  width,
		Height: height,
	}
}

// 当前匹配信息
type Matched struct {
	AreaMap string
	Area    *textmap.TextMap
	Map     *localmap.LocalMap
	AreaBox *Box
	MapBox  *Box
}

func Search(text string) *Result {
	text = Replacer2.Replace(text)
	text = Replacer.Replace(text)
	textarea := textmap.Import(localmap.Replacerfilters.Replace(text))
	if textarea.Width < localmap.TileWidth || textarea.Height < localmap.TileHeight {
		return nil
	}
	if textarea.Width*textarea.Height < MatchedMinSize {
		return nil
	}

	//匹配地图碎片
	tilelist := []*Matched{}

	for y := 0; y < textarea.Height-localmap.TileHeight+1; y++ {
		for x := 0; x < textarea.Width-localmap.TileWidth+1; x++ {
			key := strings.Join(textarea.Crop(x, y, localmap.TileWidth, localmap.TileHeight), "\n")
			chineses := rechinese.FindAllString(key, -1)
			if len(chineses) < TileMinChinese {
				continue
			}
			tiles := localmap.DefaultManager.GetTiles(key)
			for _, tile := range tiles {
				tilelist = append(tilelist, &Matched{
					AreaMap: text,
					Area:    textarea,
					Map:     tile.Map,
					AreaBox: NewBox(x, y, localmap.TileWidth, localmap.TileHeight),
					MapBox:  NewBox(tile.Left, tile.Top, tile.Width, tile.Height),
				})
			}
		}
	}
	// 	//过滤太近的碎片
	// 	filteredtiles := []*Matched{}
	// Tile:
	// 	for _, tile := range tilelist {
	// 		for _, target := range filteredtiles {
	// 			if target.Map.ID != tile.Map.ID {
	// 				continue
	// 			}
	// 			xoffset := tile.MapBox.Left - target.MapBox.Left
	// 			yoffset := tile.MapBox.Top - target.MapBox.Top
	// 			if xoffset < localmap.TileWidth || xoffset > -localmap.TileWidth || yoffset < localmap.TileHeight || yoffset > -localmap.TileHeight {
	// 				continue Tile
	// 			}
	// 		}
	// 		filteredtiles = append(filteredtiles, tile)
	// 	}
	results := []*Result{}
	for _, tile := range tilelist {
		r := tryMatch(tile)
		if r != nil {
			results = append(results, r)
		}
	}
	if len(results) > 0 {
		var result *Result
		for _, r := range results {
			if result == nil {
				result = r
				continue
			}
			if result.AreaBox.Width < r.AreaBox.Width {
				result = r
			}
		}
		return result
	}
	return nil
}

type diffResult struct {
	Text      string
	TextDiff  string
	PositionX int
	PositionY int
}

type diffResults []*diffResult

func (r diffResults) Len() int {
	return len(r)
}
func (r diffResults) Less(i, j int) bool {
	if r[i].PositionY != r[j].PositionY {
		return r[i].PositionY < r[j].PositionY
	}
	if r[i].PositionX != r[j].PositionX {
		return r[i].PositionX < r[j].PositionX
	}
	return r[i].Text < r[j].Text
}
func (r diffResults) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func diff(src []string, dst []string, x int, y int) (bool, []*diffResult) {
	result := []*diffResult{}
	if len(src) == 0 || len(src) != len(dst) {
		return false, nil
	}
	for srcindex, sc := range src {
		srccurrent := []rune(sc)
		dstcurrent := []rune(dst[srcindex])
		if len(srccurrent) != len(dstcurrent) {
			return false, nil
		}
		currentx := 0
		for srctokenindex, srctokencurrent := range srccurrent {
			dsttokencurrent := dstcurrent[srctokenindex]
			if srctokencurrent != dsttokencurrent {
				if (srctokencurrent < 256 && dsttokencurrent > 255) || (dsttokencurrent < 256 && srctokencurrent > 255) {
					return false, nil
				}
				result = append(result, &diffResult{
					Text:      string(srctokencurrent),
					TextDiff:  string(dsttokencurrent),
					PositionX: x + currentx,
					PositionY: y + srcindex,
				})

			}
			if srctokencurrent < 256 {
				currentx = currentx + 1
			} else {
				currentx = currentx + 2
			}
		}
	}
	return true, result
}

const directLeft = 0
const directRight = 1
const directUp = 2
const directDown = 3

// src 为area,即传入数据
func tryZone(src []string, dst []string) bool {
	if strings.Join(src, "\n") == strings.Join(dst, "\n") {
		return true
	}
	matched, result := diff(src, dst, 0, 0)
	if !matched {
		return false
	}
	// if (len(result) * 100 / len(strings.Join(src, ""))) > TrustPercent {
	// 	return false
	// }
NextDiff:
	for _, d := range result {
		if len(remapsymbol.FindAllString(d.Text, -1)) > 0 {
			return false
		}
		if len(remapsymbol.FindAllString(d.TextDiff, -1)) > 0 {
			return false
		}
		for _, v := range []rune(d.Text) {
			if v > 255 && ChineseMap[v] != true {
				continue NextDiff
			}
		}
		return false
		// if len(rechinese.FindAllString(d.Text, -1)) > 0 {
		// 	return false
		// }

		if len(d.Text) == 1 {
			return false
		}
	}
	return true
}

// 向指定方向尝试扩展匹配区域
func expandDirect(matched *Matched, direct int) bool {
	var src []string
	var dst []string
	switch direct {
	case directLeft:
		if matched.AreaBox.Left <= 0 || matched.MapBox.Left <= 0 {
			return false
		}
		src = matched.Area.Crop(matched.AreaBox.Left-1, matched.AreaBox.Top, 1, matched.AreaBox.Height)
		dst = matched.Map.Map.Crop(matched.MapBox.Left-1, matched.MapBox.Top, 1, matched.MapBox.Height)
		break
	case directRight:
		if matched.AreaBox.Left+matched.AreaBox.Width >= matched.Area.Width || matched.MapBox.Left+matched.MapBox.Width >= matched.Map.Map.Width {
			return false
		}
		src = matched.Area.Crop(matched.AreaBox.Left+matched.AreaBox.Width, matched.AreaBox.Top, 1, matched.AreaBox.Height)
		dst = matched.Map.Map.Crop(matched.MapBox.Left+matched.MapBox.Width, matched.MapBox.Top, 1, matched.MapBox.Height)
		break
	case directUp:
		if matched.AreaBox.Top <= 0 || matched.MapBox.Top <= 0 {
			return false
		}
		src = matched.Area.Crop(matched.AreaBox.Left, matched.AreaBox.Top-1, matched.AreaBox.Width, 1)
		dst = matched.Map.Map.Crop(matched.MapBox.Left, matched.MapBox.Top-1, matched.MapBox.Width, 1)
		break
	case directDown:
		if matched.AreaBox.Top+matched.AreaBox.Height >= matched.Area.Height || matched.MapBox.Top+matched.MapBox.Height >= matched.Map.Map.Height {
			return false
		}
		src = matched.Area.Crop(matched.AreaBox.Left, matched.AreaBox.Top+matched.AreaBox.Height, matched.AreaBox.Width, 1)
		dst = matched.Map.Map.Crop(matched.MapBox.Left, matched.MapBox.Top+matched.MapBox.Height, matched.MapBox.Width, 1)
		break
	default:
		return false
	}
	if !tryZone(src, dst) {
		return false
	}
	matched.AreaBox.Expand(direct)
	matched.MapBox.Expand(direct)
	return expandDirect(matched, direct)
}

// 向四个方向匹配最大区域
func expand(matched *Matched) {
	expandDirect(matched, directLeft)
	expandDirect(matched, directRight)
	expandDirect(matched, directUp)
	expandDirect(matched, directDown)
}

func tryMatch(matched *Matched) *Result {
	expand(matched)
	if matched.AreaBox.Width*matched.AreaBox.Height < MatchedMinSize {
		return nil
	}
	areadata := matched.Area.Crop(matched.AreaBox.Left, matched.AreaBox.Top, matched.AreaBox.Width, matched.AreaBox.Height)
	mapdata := matched.Map.Map.Crop(matched.MapBox.Left, matched.MapBox.Top, matched.MapBox.Width, matched.MapBox.Height)
	match, diffresult := diff(
		areadata,
		mapdata,
		matched.MapBox.Left, matched.MapBox.Top,
	)
	if !match || len(diffresult) == 0 {
		return nil
	}
	sort.Sort(diffResults(diffresult))
	y := diffresult[0].PositionY
	x := diffresult[0].PositionX
	text := ""
	for _, diff := range diffresult {
		if diff.PositionY == y {
			if x == diff.PositionX {
				text = text + diff.TextDiff
				if len(diff.TextDiff) == 1 {
					x = x + 1
				} else {
					x = x + 2
				}
			} else {
				break
			}

		}
	}
	if text == "" {
		return nil
	}
	return &Result{
		ID:       matched.Map.ID,
		Name:     matched.Map.Name,
		Room:     text,
		AreaBox:  matched.AreaBox,
		MapBox:   matched.MapBox,
		AreaData: strings.Join(textmap.Import(matched.AreaMap).Crop(matched.AreaBox.Left, matched.AreaBox.Top, matched.AreaBox.Width, matched.AreaBox.Height), "\n"),
		MapData:  strings.Join(textmap.Import(matched.Map.Raw).Crop(matched.MapBox.Left, matched.MapBox.Top, matched.MapBox.Width, matched.MapBox.Height), "\n"),
	}
}

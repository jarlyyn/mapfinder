package textmap

import "strings"

type TextMap struct {
	Data   [][]rune
	Raw    string
	Height int
	Width  int
}

// 裁切，按指定的显示位置对给到的文字地图进行裁切
// 显示位置，每行算1,列的话，ascii字符算1,其他算2。
// 裁切超过边界时不报错。
// 裁切到宽字符的一半时，该字符也算入裁切结果。
// 裁切结果为按行分割的字符串
func (m *TextMap) Crop(left int, top int, width int, height int) []string {
	result := []string{}
	overY := top + height
	overX := left + width
	for currentY := top; currentY < overY; currentY++ {
		if currentY < 0 || currentY >= len(m.Data) {
			result = append(result, strings.Repeat(" ", width))
			continue
		}
		line := m.Data[currentY]
		currentline := []rune{}
		for i := left; i < 0; i++ {
			if i >= left && i < left+width {
				currentline = append(currentline, rune(32))
			}
		}
		currentX := 0
		lasttoken := rune(0)
		for _, token := range line {
			if currentX == left+1 && lasttoken > 255 {
				currentline = append(currentline, lasttoken)
			}
			if currentX >= left && currentX < overX {
				currentline = append(currentline, token)
			}
			if token > 255 {
				currentX = currentX + 2
			} else {
				currentX = currentX + 1
			}
			if currentX > overX {
				break
			}
			lasttoken = token
		}
		if currentX == left+1 && lasttoken > 255 {
			currentline = append(currentline, lasttoken)
		}
		for i := currentX; i < left+width; i++ {
			if i >= left && i < left+width {
				currentline = append(currentline, rune(32))
			}
		}
		result = append(result, string(currentline))
	}
	return result
}
func New() *TextMap {
	return &TextMap{
		Data: [][]rune{},
		Raw:  "",
	}
}

func Import(data string) *TextMap {
	maplines := [][]rune{}
	data = strings.ReplaceAll(data, "\r", "")
	lines := strings.Split(data, "\n")
	maxwidth := 0
	for _, line := range lines {
		mapline := []rune{}
		width := 0
		for _, token := range line {
			mapline = append(mapline, token)
			if token < 256 {
				width = width + 1
			} else {
				width = width + 2
			}
		}
		if maxwidth < width {
			maxwidth = width
		}
		maplines = append(maplines, mapline)
	}
	m := New()
	m.Raw = data
	m.Data = maplines
	m.Height = len(lines)
	m.Width = maxwidth
	return m
}

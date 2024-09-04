package textmap

import (
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	//
	//ABC甲乙D
	//AB甲乙CD
	//A甲B乙CD
	data := `
ABC甲乙D
AB甲乙CD
A甲B乙CD`
	m := Import(data)
	if len(m.Data) != 4 {
		t.Fatal()
	}
	if m.Height != 4 {
		t.Fatal()
	}
	if m.Width != 8 {
		t.Fatal()
	}
	result := strings.Join(m.Crop(0, 1, 1, 1), "")
	if result != "A" {
		t.Fatal()
	}
	result = strings.Join(m.Crop(1, 1, 6, 2), "")
	if result != "BC甲乙B甲乙C" {
		t.Fatal()
	}
	//过上边界
	result = strings.Join(m.Crop(0, -1, 1, 3), "")
	if result != "  A" {
		t.Fatal()
	}
	//过下边界
	result = strings.Join(m.Crop(1, 3, 2, 2), "")
	if result != "甲  " {
		t.Fatal()
	}
	//完全不在上边界
	result = strings.Join(m.Crop(0, -4, 1, 3), "")
	if result != "   " {
		t.Fatal()
	}
	//完全不在下边界
	result = strings.Join(m.Crop(1, 5, 2, 2), "")
	if result != "    " {
		t.Fatal()
	}
	//过左边界
	result = strings.Join(m.Crop(-1, 1, 2, 1), "")
	if result != " A" {
		t.Fatal()
	}
	//完全不在左边界
	result = strings.Join(m.Crop(-3, 1, 2, 1), "")
	if result != "  " {
		t.Fatal()
	}
	//过右边界
	result = strings.Join(m.Crop(7, 1, 2, 1), "")
	if result != "D " {
		t.Fatal()
	}
	//完全不在右边界
	result = strings.Join(m.Crop(9, 1, 2, 1), "")
	if result != "  " {
		t.Fatal()
	}
	//切半个字符
	result = strings.Join(m.Crop(2, 3, 4, 1), "")
	if result != "甲B乙" {
		t.Fatal()
	}
	m = Import("a甲b")
	result = strings.Join(m.Crop(1, 0, 1, 1), "")
	if result != "甲" {
		t.Fatal()
	}
	result = strings.Join(m.Crop(2, 0, 1, 1), "")
	if result != "甲" {
		t.Fatal()
	}
	result = strings.Join(m.Crop(2, 0, 2, 1), "")
	if result != "甲b" {
		t.Fatal()
	}
	m = Import("a甲")
	result = strings.Join(m.Crop(2, 0, 1, 1), "")
	if result != "甲" {
		t.Fatal()
	}

}

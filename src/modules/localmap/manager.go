package localmap

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
	"sync"

	"github.com/herb-go/util"
)

type Manager struct {
	Locker     sync.Mutex
	RawDataMap map[string]*RawData
	Maps       map[string]*LocalMap
	TileIndex  map[string][]*Tile
}

func (m *Manager) GetTiles(key string) []*Tile {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.TileIndex[key]
}
func (m *Manager) index() {
	m.TileIndex = map[string][]*Tile{}
	for _, n := range m.Maps {
		for _, tile := range n.TileList {
			if m.TileIndex[tile.Data] == nil {
				m.TileIndex[tile.Data] = []*Tile{}
			}
			m.TileIndex[tile.Data] = append(m.TileIndex[tile.Data], tile)
		}
	}
}
func (m *Manager) Reset() {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	m.RawDataMap = map[string]*RawData{}
	m.Maps = map[string]*LocalMap{}
	m.TileIndex = map[string][]*Tile{}
}
func (m *Manager) Import(rd ...*RawData) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	for _, data := range rd {
		m.RawDataMap[data.ID] = data
		m.Maps[data.ID] = data.CreateLocalMap()
	}
	m.index()
}

type ListItem struct {
	ID     string
	Name   string
	Width  int
	Height int
}
type ListItems []*ListItem

func (l ListItems) Len() int {
	return len(l)
}
func (l ListItems) Less(i, j int) bool {
	return l[i].ID < l[j].ID
}
func (l ListItems) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (m *Manager) List() []*ListItem {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	result := []*ListItem{}
	for _, data := range m.Maps {
		result = append(result, &ListItem{
			ID:     data.ID,
			Name:   data.Name,
			Width:  data.Map.Width,
			Height: data.Map.Height,
		})
	}
	sort.Sort(ListItems(result))
	return result
}
func (m *Manager) GetMap(id string) *LocalMap {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	return m.Maps[id]
}
func (m *Manager) Export() []*RawData {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	result := []*RawData{}
	for _, data := range m.RawDataMap {
		result = append(result, data)
	}
	return result
}
func (m *Manager) Remove(idlist ...string) {
	m.Locker.Lock()
	defer m.Locker.Unlock()
	for _, id := range idlist {
		delete(m.RawDataMap, id)
		delete(m.Maps, id)
	}
	m.index()

}
func NewManager() *Manager {
	return &Manager{
		Maps:       map[string]*LocalMap{},
		TileIndex:  map[string][]*Tile{},
		RawDataMap: map[string]*RawData{},
	}
}

var DefaultManager = NewManager()

type DataList struct {
	Maps []*RawData
}

func NewDataList() *DataList {
	return &DataList{
		Maps: []*RawData{},
	}
}
func MustLoad() {
	data, err := os.ReadFile(util.AppData("data", "maps.json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			DefaultManager.Reset()
			return
		}
		panic(err)
	}
	dl := NewDataList()
	err = json.Unmarshal(data, dl)
	if err != nil {
		panic(err)
	}
	DefaultManager.Import(dl.Maps...)
}

func MustSave() {
	dl := NewDataList()
	dl.Maps = DefaultManager.Export()
	data, err := json.Marshal(dl)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(util.AppData("data", "maps.json"), data, util.DefaultFileMode)
	if err != nil {
		panic(err)
	}
}

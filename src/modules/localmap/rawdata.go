package localmap

type RawData struct {
	ID   string
	Name string
	Data string
}

func (d *RawData) CreateLocalMap() *LocalMap {
	return Create(d.ID, d.Name, d.Data)
}

func NewRawData(id string, name string, data string) *RawData {
	return &RawData{
		ID:   id,
		Name: name,
		Data: data,
	}
}

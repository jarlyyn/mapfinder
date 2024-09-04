package actionsearch

import (
	"io"
	"modules/search"
	"net/http"

	"github.com/herb-go/herb/ui/render"
)

var Action = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	result := search.Search(string(data))
	render.MustJSON(w, result, 200)
}

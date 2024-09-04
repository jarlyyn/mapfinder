package actionviewmap

import (
	"modules/localmap"
	"net/http"

	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/ui/render"
)

type Result struct {
	ID   string
	Name string
	Data string
}

var Action = func(w http.ResponseWriter, r *http.Request) {
	id := router.GetParams(r).Get("id")
	lmap := localmap.DefaultManager.GetMap(id)
	if lmap == nil {
		render.MustError(w, 404)
		return
	}
	result := &Result{
		ID:   lmap.ID,
		Name: lmap.Name,
		Data: lmap.Map.Raw,
	}
	render.MustJSON(w, result, 200)
}

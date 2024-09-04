package actionlistmap

import (
	"modules/localmap"
	"net/http"

	"github.com/herb-go/herb/ui/render"
)

var Action = func(w http.ResponseWriter, r *http.Request) {
	list := localmap.DefaultManager.List()
	render.MustJSON(w, list, 200)
}

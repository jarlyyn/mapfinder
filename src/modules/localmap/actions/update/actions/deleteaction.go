package actions

import (
	"modules/localmap"
	"net/http"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/ui/render"
)

var ActionDelete = action.New(func(w http.ResponseWriter, r *http.Request) {
	id := router.GetParams(r).Get("id")
	localmap.DefaultManager.Remove(id)
	localmap.MustSave()
	render.MustJSON(w, "ok", 200)
})

package actionutf8update

import (
	"io"
	"modules/localmap"
	"net/http"
	"strings"

	"github.com/herb-go/herb/middleware/router"
)

var Action = func(w http.ResponseWriter, r *http.Request) {
	id := router.GetParams(r).Get("id")
	name := r.Header.Get("mapname")
	if name == "" || strings.Contains(name, " ") {
		w.WriteHeader(400)
		w.Write([]byte("Header field mapname wrong"))
		return
	}
	defer r.Body.Close()
	lm, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	localmap.DefaultManager.Import(localmap.NewRawData(id, name, string(lm)))
	localmap.MustSave()
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

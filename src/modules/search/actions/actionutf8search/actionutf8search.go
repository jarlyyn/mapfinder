package actionutf8search

import (
	"io"
	"modules/search"
	"net/http"
)

var Action = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	result := search.Search(string(body))
	if result == nil {
		w.Write([]byte(""))
		return
	}
	w.Write([]byte(result.Name + " " + result.Room))
}

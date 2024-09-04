package actiongbkupdate

import (
	"io"
	"modules/localmap"
	"net/http"
	"strings"

	"github.com/herb-go/herb/middleware/router"
	"golang.org/x/text/encoding/simplifiedchinese"
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
	decoder := simplifiedchinese.GBK.NewDecoder()
	utfid, err := decoder.Bytes([]byte(id))
	if err != nil {
		panic(err)
	}
	utfname, err := decoder.Bytes([]byte(name))
	if err != nil {
		panic(err)
	}
	utfdata, err := decoder.Bytes(lm)
	if err != nil {
		panic(err)
	}
	localmap.DefaultManager.Import(localmap.NewRawData(string(utfid), string(utfname), string(utfdata)))
	localmap.MustSave()
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

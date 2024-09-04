package actiongbksearch

import (
	"io"
	"modules/search"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var Action = func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	decoder := simplifiedchinese.GBK.NewDecoder()
	data, err := decoder.Bytes(body)
	if err != nil {
		panic(err)
	}
	result := search.Search(string(data))
	if result == nil {
		w.Write([]byte(""))
		return
	}
	encoder := simplifiedchinese.GBK.NewEncoder()
	output, err := encoder.Bytes([]byte(result.Name + " " + result.Room))
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

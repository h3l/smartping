package http

import (
	"../g"
	"net/http"
	"path/filepath"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/") {
		if !g.IsExist(filepath.Join(g.Root, "/html", r.URL.Path, "index.html")) {
			http.NotFound(w, r)
			return
		}
	}
	http.FileServer(http.Dir(filepath.Join(g.Root, "/html"))).ServeHTTP(w, r)
}

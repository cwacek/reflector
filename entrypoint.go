package main

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func InjectLogger(next http.Handler) http.Handler {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL,
		}).Debug("R:")

		next.ServeHTTP(w, r)
	})
	return handler
}

type VersionedDocHandler struct {
	Source   *FileSystem
	projects map[string]bool
}

func (h *VersionedDocHandler) Respond(project, version string, w http.ResponseWriter) error {
	if have, ok := h.projects[project]; !ok {
		http.Error(w, "Not Found", 400)
		return nil
	}

	var path string
	if version == "" {
		path = filepath.Join(project, "latest")
	} else {
		path = filepath.Join(project, version)
	}

}

func makeDocRouter(handler VersionedDocHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		var pathParts = strings.SplitN(r.URL.Path, "/", 1)
		project := pathParts[0]
		var version string
		version = r.Form.Get("version")

		handler.Respond(project, version, w)
	}
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	var view VersionedDocHandler
	view = S3View()
	docRouter := makeDocRouter(view)

	handler := InjectLogger(docRouter)

	http.ListenAndServe("0.0.0.0:8091", handler)

}

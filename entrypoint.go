package main

import (
	"net/http"

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

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	s3View := NewS3View()
	handler := InjectLogger(s3View)

	http.ListenAndServe("0.0.0.0:8091", handler)

}

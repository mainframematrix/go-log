package middleware

import (
	"log"
	"net/http"

	logger "github.com/mainframematrix/go-log/src"
)

func PathLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l, err := logger.CreateLogger(logger.INFO, "stdout", true)
		if err != nil {
			log.Fatalf("Error initializing logger: %v", err)
		}
		l.Log(logger.INFO, "%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

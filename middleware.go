package tweetlator

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"time":  time.Now().Sub(t).Seconds(),
			"route": r.RequestURI,
		}).Info()
	})
}

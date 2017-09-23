package main

import (
	"net/http"
	"os"

	"github.com/cbelsole/tweetlator"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

// Instantiates a translator that is backed by the Microsoft Translation API and passes it to helloWorld.
// Get your own subscription key by registering a Microsoft Text Translation service in Azure.
// See http://docs.microsofttranslator.com/text-translate.html.
func main() {
	transKey := os.Getenv("TRANSLATOR_KEY")
	if transKey == "" {
		log.Fatal("translation key missing")
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Infof("cannot parse log level `%s`, using warn instead")
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(level)
	}

	r := mux.NewRouter()
	server := http.Server{
		Addr:    ":8000",
		Handler: tweetlator.LogRequests(r),
	}
	log.Debug("starting server")
	tr := tweetlator.NewTranslationHandler(transKey)
	r.HandleFunc("/translate", tr.Translation).Methods("POST")
	r.HandleFunc("/", tr.Health).Methods("GET")
	r.HandleFunc("/health", tr.Health).Methods("GET")
	server.ListenAndServe()
}

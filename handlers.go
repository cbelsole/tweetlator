package tweetlator

import (
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type (
	translateReq struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Tweet string `json:"tweet"`
	}
	translateRes struct {
		Tweet string `json:"tweet"`
	}
)
type TranslationHandler struct {
	key string
}

func NewTranslationHandler(key string) *TranslationHandler {
	return &TranslationHandler{key}
}

func (t *TranslationHandler) Translation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req translateReq
	err := decoder.Decode(&req)
	if err != nil {
		log.Debug(err)
		jsonError(w, r, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tr := NewTranslate(t.key)
	tweet, err := tr.translate(req.Tweet, req.To, req.From)
	if err != nil {
		log.Debug(err)
		jsonError(w, r, err, http.StatusInternalServerError)
		return
	}
	jsonResponse(w, r, translateRes{tweet}, http.StatusOK)
}

func (t *TranslationHandler) Health(w http.ResponseWriter, r *http.Request) {
	tr := NewTranslate(t.key)
	if err := tr.health(); err != nil {
		log.Debug(err)
		jsonError(w, r, errors.New("we are unhealthy"), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, r, map[string]string{"message": "we are healthy"}, http.StatusOK)
}

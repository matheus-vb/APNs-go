package api

import (
	"github.com/go-chi/chi/v5"
	"go-server/util"
	"log"
	"net/http"
)

type JSONPayload struct {
}

func (app *App) SendNotification(w http.ResponseWriter, r *http.Request) {
	deviceToken := chi.URLParam(r, "deviceToken")

	payload := util.JSONResponse{
		Error:   false,
		Message: "token",
		Data:    deviceToken,
	}

	err := util.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		log.Println(err)
	}
}

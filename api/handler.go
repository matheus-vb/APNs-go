package api

import (
	"github.com/go-chi/chi/v5"
	"go-server/util"
	"log"
	"net/http"
)

func (app *App) SendNotification(w http.ResponseWriter, r *http.Request) {
	deviceToken := chi.URLParam(r, "deviceToken")

	res, err := app.Provider.SendNotification(deviceToken, "Message from Go!")
	if err != nil {
		log.Fatal(err)
	}

	payload := util.JSONResponse{
		Error:   false,
		Message: "result",
		Data:    res.Sent(),
	}

	err = util.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		log.Println(err)
	}
}

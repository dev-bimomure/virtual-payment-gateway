package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) VirtualPayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stringMap := make(map[string]string)
	stringMap["publicKey"] = app.config.apiKeys.publicKey
	err := app.renderTemplate(w, r, "payment", &templateData{
		StringMap: stringMap,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

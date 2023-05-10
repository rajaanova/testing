package handler

import (
	"codebase/intve/service"
	"encoding/json"
	"net/http"
)

type routes struct {
	service service.User
}

func (r *routes) getUser(w http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	if val, ok := r.service.GetUser(id); ok {
		resBody, err := json.Marshal(val)
		if err != nil {
			//send error here
		}
		w.Write(resBody)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	return

}

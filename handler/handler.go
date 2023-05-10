package handler

import (
	"codebase/intve/service"
	"net/http"
)

type Routes struct {
	Service service.Service
}

func (r *Routes) GetUser(w http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	val, err := r.Service.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if val == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

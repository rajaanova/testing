package main

import (
	"codebase/intve/handler"
	"codebase/intve/service"
	"codebase/intve/storage"
	"net/http"
)

func init() {

}
func main() {
	ma := map[string]*storage.User{}
	dataVal := storage.MapStorage{IdByUser: ma}
	serviceVal := service.Service{Store: dataVal}
	routes := handler.Routes{Service: serviceVal}
	http.HandleFunc("/user", routes.GetUser)
	http.ListenAndServe(":8080", nil)

}

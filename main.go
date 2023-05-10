package main

import (
	"codebase/intve/handler"
	"codebase/intve/service"
	"codebase/intve/storage"
	"fmt"
	"net/http"
)

func init() {

	a := make([]int, 32)
	b := a[1:16]
	//a = append(a, 1)
	a[2] = 42
	fmt.Println(b)
	fmt.Println("hello raj")
}
func main() {
	ma := map[string]*storage.User{}
	dataVal := storage.MapStorage{IdByUser: ma}
	serviceVal := service.Service{Store: dataVal}
	routes := handler.Routes{Service: serviceVal}
	http.HandleFunc("/user", routes.GetUser)
	http.ListenAndServe(":8080", nil)

}

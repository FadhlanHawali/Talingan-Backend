package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"

	"github.com/Talingan-Backend/pkg/api"
)

func main() {

	if err := api.VisionClassificationPredict(); err != nil {
		fmt.Printf(err.Error())
		return
	}

	router := mux.NewRouter()

	handler := cors.AllowAll().Handler(router)
	port := fmt.Sprintf(":%s", viper.Get("host.port"))
	log.Printf("Server Running on port %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func checkHealth() {

}

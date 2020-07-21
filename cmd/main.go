package main

import (
	"log"
	"net/http"

	"github.com/Talingan-Backend/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// viper.SetConfigFile("./config/dev.json")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// }

	// if err := api.VisionClassificationPredict(); err != nil {
	// 	fmt.Printf(err.Error())
	// 	return
	// }

	router := mux.NewRouter()
	router.HandleFunc("/", http.HandlerFunc(checkHealth))
	handler := cors.AllowAll().Handler(router)
	// port := fmt.Sprintf(":%s", viper.Get("host.port"))
	log.Printf("Server Running on port %s", "8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "Success", 200)
	return
}

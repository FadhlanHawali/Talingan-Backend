package main

import (
	"context"
	"github.com/Talingan-Backend/database"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"

	"github.com/Talingan-Backend/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	viper.SetConfigFile("./configs/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// if err := api.VisionClassificationPredict(); err != nil {
	// 	fmt.Printf(err.Error())
	// 	return
	// }

	db := database.DBInit("mongodb://127.0.0.1:27017")
	inDB := &InDB{DB: db}
	router := mux.NewRouter()
	router.HandleFunc("/", http.HandlerFunc(inDB.checkHealth))
	handler := cors.AllowAll().Handler(router)
	// port := fmt.Sprintf(":%s", viper.Get("host.port"))
	log.Printf("Server Running on port %s", "8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func (idb *InDB) checkHealth(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "Success", 200)
	err := idb.DB.Ping(context.Background(), readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println("Connected !")
	}
	return
}

type InDB struct{
	DB *mongo.Client
}

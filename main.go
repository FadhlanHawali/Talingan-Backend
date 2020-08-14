package main

import (
	"context"
	"fmt"
	config "github.com/Talingan-Backend/configs"
	"github.com/Talingan-Backend/database"
	"github.com/Talingan-Backend/pkg/file"
	"github.com/Talingan-Backend/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)
const maxUploadSize = 2 * 1024 * 1024 // 2 mb
const uploadPath = "./tmp"
func main() {
	var configuration config.Configuration

	viper.SetConfigFile("./configs/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// if err := api.VisionClassificationPredict(); err != nil {
	// 	fmt.Printf(err.Error())
	// 	return
	// }

	db := database.DBInit(configuration.Database.ConnectionURI)
	inDB := &InDB{DB: db}
	router := mux.NewRouter()
	router.HandleFunc("/", http.HandlerFunc(inDB.checkHealth))
	router.HandleFunc("/file", http.HandlerFunc(file.UploadFileHandler))
	router.HandleFunc("/api/v1/webhook",checkGithub)
	handler := cors.AllowAll().Handler(router)
	// port := fmt.Sprintf(":%s", viper.Get("host.port"))
	port:=configuration.Server.Port
	log.Printf("Server Running on port %d", port)
	log.Println("checks122")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), handler))

}

func checkGithub(w http.ResponseWriter, r *http.Request){
	utils.WrapAPISuccess(w,r,"success",http.StatusOK)
	return
}

func (idb *InDB) checkHealth(w http.ResponseWriter, r *http.Request) {

	err := idb.DB.Ping(context.Background(), readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println("Connected !")
		utils.WrapAPISuccess(w, r, "Success", 200)
	}
	return
}

type InDB struct{
	DB *mongo.Client
}


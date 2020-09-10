package services

import (
	"fmt"
	"github.com/Talingan-Backend/v2/internal/app/services"
	"github.com/Talingan-Backend/v2/internal/delivery/rest/services/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)



func initRouter(router *mux.Router, svc *service.Services){
	router.Handle("/api/v1/upload-file",http.HandlerFunc(svc.ServicesService.UploadFileHandler))
	handler := cors.AllowAll().Handler(router)

	//fmt.Println("Apps served on :" + args.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":8008"), handler))
}

func Start (app *services.ServicesRest){

	svc := service.GetServices(app)
	router := mux.NewRouter()
	initRouter(router, svc)
}

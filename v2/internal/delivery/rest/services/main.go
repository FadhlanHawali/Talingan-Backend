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
	router.Handle("/api/v1/kandang/new",http.HandlerFunc(svc.ServicesService.InsertNewKandang))
	router.Handle("/api/v1/kandangs",http.HandlerFunc(svc.ServicesService.GetKandangs))
	router.Handle("/api/v1/monitoring/kandang/{idKandang}",http.HandlerFunc(svc.ServicesService.GetMonitoring))
	router.Handle("/api/v1/notification",http.HandlerFunc(svc.ServicesService.GetNotification))
	handler := cors.AllowAll().Handler(router)

	fmt.Println("Apps served on :8008")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":8008"), handler))
}

func Start (app *services.ServicesRest){

	svc := service.GetServices(app)
	router := mux.NewRouter()
	initRouter(router, svc)
}

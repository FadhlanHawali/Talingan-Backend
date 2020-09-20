package main

import (
	"github.com/Talingan-Backend/v2/internal/app/services"
	services2 "github.com/Talingan-Backend/v2/internal/delivery/rest/services"
	"log"
)

func main(){
	servicesRestApp, err := services.NewServicesRest()
	if err != nil {
		log.Fatalf("marshal error %+v", err)
	}
	defer func() {
		errs := servicesRestApp.Close()
		for e := range errs{
			log.Println(e)
		}
	}()

	services2.Start(servicesRestApp)
}

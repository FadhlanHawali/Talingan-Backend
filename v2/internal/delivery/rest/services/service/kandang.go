package service

import (
	"encoding/json"
	"github.com/Talingan-Backend/v2/internal/entity"
	"github.com/Talingan-Backend/v2/internal/helper"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *ServicesService) InsertNewKandang (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		helper.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		helper.WrapAPIError(w,r,"can't read body",http.StatusBadRequest)
		return
	}
	log.Println(r.Body)
	var kandang entity.Kandang
	var cobo interface{}
	err = json.Unmarshal(body,&cobo)
	log.Println(cobo)
	err = json.Unmarshal(body, &kandang)
	if err != nil {
		log.Println("ERROR : " + err.Error())
		helper.WrapAPIError(w,r,"error unmarshal : "+err.Error(),http.StatusInternalServerError)
		return
	}

	flag,err := s.uc.InsertNewKandang(&kandang); if err != nil {
		helper.WrapAPIError(w,r,err.Error(),http.StatusBadRequest)
		return
	}
	log.Println(flag)
	helper.WrapAPISuccess(w,r,"success",http.StatusOK)
	return
}

func (s *ServicesService) GetKandangs (w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		helper.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}

	_,err,res := s.uc.GetKandangs();if err != nil{
		helper.WrapAPIError(w,r,err.Error(),http.StatusBadRequest)
		return
	}

	helper.WrapAPIData(w,r,res,http.StatusOK,"success")
	return
}

func (s *ServicesService) GetKandang (w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		helper.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}
}

func (s *ServicesService) InsertMonitoring (w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		helper.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		helper.WrapAPIError(w,r,"can't read body",http.StatusBadRequest)
		return
	}

	var monitoring entity.DeteksiKandang
	err = json.Unmarshal(body, &monitoring)
	if err != nil {
		helper.WrapAPIError(w,r,"error unmarshal : "+err.Error(),http.StatusInternalServerError)
		return
	}

	flag,err := s.uc.InsertMonitoring(&monitoring); if err != nil {
		helper.WrapAPIError(w,r,err.Error(),http.StatusBadRequest)
		return
	}
	log.Println(flag)
	helper.WrapAPISuccess(w,r,"success",http.StatusOK)
	return
}

func (s *ServicesService) GetMonitoring (w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		helper.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	idKandang := vars["idKandang"]
	_,err,res := s.uc.GetMonitoring(idKandang);if err != nil{
		helper.WrapAPIError(w,r,err.Error(),http.StatusBadRequest)
		return
	}

	helper.WrapAPIData(w,r,res,http.StatusOK,"success")
	return
}

func (s *ServicesService) GetNotification (w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		helper.WrapAPIError(w,r,"Bad request method", http.StatusBadRequest)
		return
	}
	_,err,res := s.uc.GetNotification();if err != nil{
		helper.WrapAPIError(w,r,err.Error(),http.StatusBadRequest)
		return
	}

	helper.WrapAPIData(w,r,res,http.StatusOK,"success")
	return
}

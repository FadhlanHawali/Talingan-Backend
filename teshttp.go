package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main (){
	//fmt.Sprintf("TOTAL DRAW : %d",getNumDraws(2011))
	fmt.Sprintf("TOTAL GOAL : %d",getWinnerTotalGoals("UEFA Champions League",2011))
}



func getNumDraws(year int32) int32 {
	var matches Matches
	var totalDraw int32
	totalDraw = 0

	for i := 0;i<=10;i ++ {
		req, err := http.NewRequest("GET",fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?year=%d&team1goals=%d&team2goals=%d",year,i,i),nil)
		if err != nil {
			log.Println(err.Error())
		}

		client :=  &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err.Error())
		}else{
			if err := json.NewDecoder(resp.Body).Decode(&matches);err != nil {
				log.Println(err.Error())
			} else {
				totalDraw = matches.Total + totalDraw
			}
		}
	}
	log.Println("TOTALDRAW : ",totalDraw)

	return totalDraw
}


func getWinnerTotalGoals(competition string, year int32) int32 {
	var matches Matches
	var matchWinner MatchWinner
	var totalPages, totalGoal int32
	var clubWinner string
	totalPages = 0
	totalGoal = 0
	competitionReplaced := strings.ReplaceAll(competition," ","%20")

	req, err := http.NewRequest("GET",fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_competitions?name=%s",competitionReplaced),nil)
	if err != nil {
		log.Println(err.Error())
	}

	client :=  &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}else{
		if err := json.NewDecoder(resp.Body).Decode(&matchWinner);err != nil {
			log.Println(err.Error())
		} else {

			for _, v:= range matchWinner.Data{

				if v.Year == year && v.Name == competition{
					clubWinner = v.Winner
					log.Println("match winner : ",clubWinner)
				}
			}
		}
	}

	log.Println("WINNER : ",clubWinner)
	req, err = http.NewRequest("GET",fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?competition=%s&year=%d",competitionReplaced,year),nil)
	if err != nil {
		log.Println(err.Error())
	}

	client =  &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}else{
		if err := json.NewDecoder(resp.Body).Decode(&matches);err != nil {
			log.Println(err.Error())
		} else {
			totalPages = matches.TotalPages
			log.Println("Matches : ",matches)
			log.Println("Total pages : ",totalPages)
		}
	}

	var j int32
	var matches2 Matches2
	for j= 1; j<=totalPages;j++{
		competition = strings.ReplaceAll(competition," ","%20")
		req, err := http.NewRequest("GET",fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?competition=%s&year=%d&page=%d",competitionReplaced,year,j),nil)
		if err != nil {
			log.Println(err.Error())
		}

		client :=  &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err.Error())
		}else{
			if err := json.NewDecoder(resp.Body).Decode(&matches2);err != nil {
				log.Println("ERROR HERE : ", err.Error())
			} else {

				totalGoal = totalGoal + getData(clubWinner,matches2.Data)
			}
		}
	}
	log.Println("TOTAL GOAL = ",totalGoal)
	return totalGoal
}

func getData(team string, data []Data) int32{

	var totalGoal int
	totalGoal = 0
	for _,v := range data{
		if v.Team1 == team || v.Team2 == team {
			team1goals ,_ := strconv.Atoi(v.Team1Goals)
			team2goals, _ := strconv.Atoi(v.Team2Goals)

			if v.Team1 == team {
				log.Println(fmt.Sprintf("%s : %d",v.Team1,team1goals))
				totalGoal = team1goals + totalGoal
			}else if v.Team2 == team {
				log.Println(fmt.Sprintf("%s : %d",v.Team2,team2goals))
				totalGoal = team2goals + totalGoal
			}
		}

	}

	return int32(totalGoal)
}

type Matches struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	Total int32 `json:"total"`
	TotalPages int32 `json:"total_pages"`
	Data []Data `json:"data"`
}

type Data struct {
	Competition string `json:"competition"`
	Year int `json:"year"`
	Round string `json:"round"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Team1Goals string `json:"team1goals"`
	Team2Goals string `json:"team2goals"`
}

type Matches2 struct {
	Page string `json:"page"`
	PerPage int `json:"per_page"`
	Total int32 `json:"total"`
	TotalPages int32 `json:"total_pages"`
	Data []Data `json:"data"`
}

type Data2 struct {
	Competition string `json:"competition"`
	Year int `json:"year"`
	Round string `json:"round"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Team1Goals string `json:"team1goals"`
	Team2Goals string `json:"team2goals"`
}

type MatchWinner struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	Total int32 `json:"total"`
	TotalPages int32 `json:"total_pages"`
	Data []DataWinner `json:"data"`
}

type DataWinner struct {
	Name string `json:"name"`
	Country string `json:"country"`
	Year int32 `json:"year"`
	Winner string `json:"winner"`
	RunnerUp string `json:"runnerup"`
}
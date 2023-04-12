package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for {
		postData()
		time.Sleep(15 * time.Second)
	}
}

func postData() {
	data := map[string]interface{}{
		"water": rand.Intn(100) + 1,
		"wind":  rand.Intn(100) + 1,
	}

	requestJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		log.Fatalln(err)
	}

	statusWater, statusWind := getStatus(responseMap)
	fmt.Println("status water: ", statusWater)
	fmt.Println("status wind: ", statusWind)
}

func getStatus(data map[string]interface{}) (string, string) {
	water := int(data["water"].(float64))
	wind := int(data["wind"].(float64))

	var waterStat, windStat string

	switch {
	case water < 5:
		waterStat = "aman"
	case water >= 6 && water <= 8:
		waterStat = "siaga"
	case water > 8:
		waterStat = "bahaya"
	}

	switch {
	case wind < 6:
		windStat = "aman"
	case wind >= 7 && wind <= 15:
		windStat = "siaga"
	case wind > 15:
		windStat = "bahaya"
	}

	return waterStat, windStat
}

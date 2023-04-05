package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const url = "https://jsonplaceholder.typicode.com/posts"

type Data struct {
	ValueWater int    `json:"valuewater"`
	ValueWind  int    `json:"valuewind"`
	StatusWater string `json:"statuswater"`
	StatusWind  string `json:"statuswind"`
}

func getStatus(value int) string {
	if value < 5 {
		return "aman"
	} else if value >= 5 && value <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1
		statusWater := getStatus(valueWater)
		statusWind := getStatus(valueWind)

		data := Data{
			ValueWater: valueWater,
			ValueWind:  valueWind,
			StatusWater: statusWater,
			StatusWind:  statusWind,
		}

		payload, _ := json.Marshal(data)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp.StatusCode)
			fmt.Println(resp.Body)
			resp.Body.Close()
		}

		time.Sleep(15 * time.Second)package main

		import (
			"bytes"
			"encoding/json"
			"fmt"
			"math/rand"
			"net/http"
			"time"
		)
		
		const url = "https://jsonplaceholder.typicode.com/posts"
		
		type Data struct {
			ValueWater  int    `json:"valuewater"`
			ValueWind   int    `json:"valuewind"`
			StatusWater string `json:"statuswater"`
			StatusWind  string `json:"statuswind"`
		}
		
		func getStatus(value int) string {
			if value < 5 {
				return "aman"
			} else if value >= 5 && value <= 8 {
				return "siaga"
			} else {
				return "bahaya"
			}
		}
		
		func main() {
			rand.Seed(time.Now().UnixNano())
		
			for {
				valueWater := rand.Intn(100) + 1
				valueWind := rand.Intn(100) + 1
				statusWater := getStatus(valueWater)
				statusWind := getStatus(valueWind)
		
				data := Data{
					ValueWater:  valueWater,
					ValueWind:   valueWind,
					StatusWater: statusWater,
					StatusWind:  statusWind,
				}
		
				payload, _ := json.Marshal(data)
		
				resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(resp.StatusCode)
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(resp.Body)
					fmt.Println(buf.String())
					_ = resp.Body.Close()
				}
		
				time.Sleep(15 * time.Second)
			}
		}
		
	}
}

package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Info struct {
	City string `json:"city"`
}

type Coordinates struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Track struct {
	Items []struct {
		Track struct {
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusOK)

}

func ReceiveInfo(w http.ResponseWriter, r *http.Request) {
	var info Info

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%+v", info)
    fmt.Print(info.City)


	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + info.City + "&appid=b77e07f479efe92156376a8b07640ced")
	if err != nil {
		// handle error
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var coord Coordinates

	if err := json.Unmarshal(body, &coord); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	temper := coord.Main.Temp
	fmt.Println(temper)

	var ftc float64

	ftc = (temper - 273.15)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("%.2f", ftc)

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/37i9dQZF1DXb1oSegSL8id/tracks?market=ES&fields=items(track(name(name)))&limit=10",nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer BQBk3QWuNpXkJXz-Jeidk1Da71-Pd1-Q2YCAQaRK8o_C3JVgcGSNmWjveychm6UPm77QXGSmC54_x6YeqZ0W1Ptgu2h4iNufbPwM5hS--AekGDMI9yP6KtFAEPEYwSiasYh5TY88-5MQBUO4n-dTAhFLLr0SxOYu-g22rZX5SUBmyXACWkB1pO_0cjyHPYI4mU71W7_K_uVlMw")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var track map[string]interface{}

	err = json.Unmarshal(body, &track)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ftc > 30 {
    	log.Print("Party")
		for _, val := range track {
			party := val
			log.Print(party)
		}
	} else if ftc > 15 && ftc < 30 {
		log.Print("Chilli Beat")
	} else if ftc > 10 && ftc < 24 {
		log.Print("Rock")
	} else if ftc <= 10 {
		log.Print("Classical Music")
	}

}
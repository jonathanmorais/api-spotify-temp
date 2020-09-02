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

type Chilli struct {
	Items []struct {
		Track struct {
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
}

type Playlist struct {
	PartyId     string `json:"partyid"`
	ChilliId    string `json:"chilliid"`
	RockId      string `json:"rockid"`
	ClassicalId string `json:"classicalid"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusOK)

}

func newPlaylist(PartyId, ChilliId, RockId, ClassicalId string ) *Playlist {
	p := Playlist{
		PartyId:     "37i9dQZF1DX8mBRYewE6or",
		ChilliId:    "2rN3mSrzUcgjlj1TcEDTX7",
		RockId:      "37i9dQZF1DX8mBRYewE6or",
		ClassicalId: "37i9dQZF1DWWEJlAGA9gs0",
	}
	log.Printf(p.PartyId)

	return newPlaylist(PartyId, ChilliId, RockId, ClassicalId)
}

func ReceiveCity(w http.ResponseWriter, r *http.Request) float64 {
	var info Info

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0
	}

	fmt.Fprintf(w, "%+v", info)
	fmt.Print("%s\n", info.City)

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

	var ftc float64

	ftc = (temper - 273.15)
	if err != nil {
		log.Print(err)
	}

	log.Printf("%.2f", ftc)

	return ftc

}

func GetTrack(w http.ResponseWriter, r *http.Request) {

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/2rN3mSrzUcgjlj1TcEDTX7/tracks?market=ES&fields=items(track(name(name)))&limit=1", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer BQBAP8qKdNDZeQc_vMf9iIqV-Mh68J6_DuH_WCf26rg-kAc78YosL0_W6z8HhtGnXS-JbcXwNb6AXz3sFMckX_4tawesbyb-9yz66y9N1uJDJDyi6WfWgKWLmlq9JuULrhpe6FoRQMn_RUfvoygpxaaajBdnBN3dCDoE9e6ZKs9X6BbAHIj0URFzwSpTUfHaKlw15J3yXVmoeA")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var track Chilli

	err = json.Unmarshal(body, &track)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuggestionTrack() {
	if ftc > 30 {
    	log.Print("Party")
		for _, val := range track.Items {
			party := val.Track.Name
			log.Print(party)
		}
	} else if ftc > 15 && ftc < 30 {
		log.Print("Chilli Beat")
		for _, val := range track.Items {
			chilli := val.Track.Name
			fmt.Print(chilli)
		}
	} else if ftc > 10 && ftc < 24 {
		log.Print("Rock")
	} else if ftc <= 10 {
		log.Print("Classical Music")
	}

	return 0
}
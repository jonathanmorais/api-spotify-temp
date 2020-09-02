package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	http "net/http"
	"os"
	"strconv"
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
	PartyId     string
	ChilliId    string
	RockId      string
	ClassicalId string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusOK)

}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func newPlaylist() (string, string, string, string) {
	p := Playlist{
		PartyId:     "37i9dQZF1DX8mBRYewE6or",
		ChilliId:    "2rN3mSrzUcgjlj1TcEDTX7",
		RockId:      "37i9dQZF1DX8mBRYewE6or",
		ClassicalId: "37i9dQZF1DWWEJlAGA9gs0",
	}
	log.Printf(p.PartyId)

	return p.PartyId, p.ChilliId, p.RockId, p.ClassicalId
}

func ReceiveCoordinates(w http.ResponseWriter, r *http.Request) string {
	var info Info

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	temp := (temper - 273.15)
	if err != nil {
		log.Print(err)
	}

     ftc := strconv.FormatFloat(temp, 'E', -1, 64)
	 fmt.Printf("%T, %v\n", ftc, ftc)

	return ftc
}

func GetTrack() Chilli {

	party, _, _, _ := newPlaylist()

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/playlists/" + party + "/tracks?market=ES&fields=items(track(name(name)))&limit=1", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer BQBAP8qKdNDZeQc_vMf9iIqV-Mh68J6_DuH_WCf26rg-kAc78YosL0_W6z8HhtGnXS-JbcXwNb6AXz3sFMckX_4tawesbyb-9yz66y9N1uJDJDyi6WfWgKWLmlq9JuULrhpe6FoRQMn_RUfvoygpxaaajBdnBN3dCDoE9e6ZKs9X6BbAHIj0URFzwSpTUfHaKlw15J3yXVmoeA")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var track Chilli

	err = json.Unmarshal(body, &track)
	if err != nil {
		log.Printf(err.Error())
	}

	return track
}

func SuggestionTrack(w http.ResponseWriter, r *http.Request) {

	ftc := ReceiveCoordinates(w, r)
	f, err := strconv.ParseFloat(ftc, 64)
	if err != nil{log.Println(err)}

	track := GetTrack()

	if f > 30 {
    	log.Print("Party")
		for _, val := range track.Items {
			party := val.Track.Name
			log.Print(party)
		}
	} else if f > 15 && f < 30 {
		log.Print("Chilli Beat")
		for _, val := range track.Items {
			chilli := val.Track.Name
			fmt.Print(chilli)
		}
	} else if f > 10 && f < 24 {
		log.Print("Rock")
	} else if f <= 10 {
		log.Print("Classical Music")
	}

	return
}
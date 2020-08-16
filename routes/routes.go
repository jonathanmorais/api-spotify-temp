package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Info struct {
	City string `json:"city"`
}

type Coordinates struct {
	Dates []string `json:"dates"`
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

	fmt.Fprintf(w, "City: %+v", info)
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
	json.Unmarshal(body, &coord)
	fmt.Printf("Results: %v\n", coord)


}
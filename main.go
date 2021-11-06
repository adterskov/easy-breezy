package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/yosssi/gohtml"
)

// Find out the Weather Forecast
func weatherForecast(w http.ResponseWriter, r *http.Request) {
	// environment variable
	// TODO: create tests for input values
	apiKey := os.Getenv("API_KEY")
	outputFormat := strings.ToLower(os.Getenv("OUTPUT_FORMAT"))
	myCity := strings.ToLower(os.Getenv("MY_CITY"))

	logrus.Info("Find out the Weather Forecast")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&mode=%s&appid=%s", myCity, outputFormat, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// return http code 200
	w.WriteHeader(http.StatusOK)

	// output
	fmt.Fprint(w, string(body))
}

// Return PONG in HTML format, status code 200 OK
func pingHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Check Availability")

	h := `<!DOCTYPE html><html><head><title>Availability Check</title>
<style type="text/css">
div {font-size: 16px; font-weight: bold;}
</style></head><body>
<div>PONG</div>
</body></html>`

	// return http code 200
	w.WriteHeader(http.StatusOK)
	// output
	fmt.Fprint(w, gohtml.Format(h))
}

// Return status code 200, in JSON format
func healthHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Check Health Status")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["status"] = "200 OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	//return http code 200
	w.WriteHeader(http.StatusOK)
	// output
	w.Write(jsonResp)
}

func main() {
	router := mux.NewRouter()
	logrus.Info("Web server is running")

	// configure routes
	router.HandleFunc("/", weatherForecast)
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/health", healthHandler)

	// TODO: custom web server port
	log.Fatal(http.ListenAndServe(":8080", router))
}

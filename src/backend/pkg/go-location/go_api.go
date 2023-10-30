package golocation

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

func InitRoutes(route *mux.Router) {
	golocation, err := New()
	if err != nil {
		log.Fatal(err)
	}

	route.HandleFunc("/locations/country", func(w http.ResponseWriter, r *http.Request) {
		countries, err := golocation.AllCountries()
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(countries); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/state", func(w http.ResponseWriter, r *http.Request) {
		states, err := golocation.AllStates()
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(states); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/city", func(w http.ResponseWriter, r *http.Request) {
		cities, err := golocation.AllCities()
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(cities); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/country/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		country := golocation.GetCountry(id)

		if err := json.NewEncoder(w).Encode(country); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/city/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		city, err := golocation.GetCity(id)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(city); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/state/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])

		state, err := golocation.GetState(id)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(state); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/states/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		states, err := golocation.GetCountryStates(id)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(states); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/cities/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		cities, err := golocation.GetStateCites(id)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(cities); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/locations/language", func(w http.ResponseWriter, r *http.Request) {
		languages, err := golocation.GetLanguages()
		if err != nil {
			log.Fatal(err)
		}

		if err := json.NewEncoder(w).Encode(languages); err != nil {
			logrus.Errorf("failed to parse and send response: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

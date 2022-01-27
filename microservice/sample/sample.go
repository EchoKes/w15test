package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Busstop struct {
	BusStopCode string `json:"BusStopCode"`
	Description string `json:"Description"`
}

var busstopArray []Busstop

// var db *sql.DB

func landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "~ Sample Landing ~")
}

// Updates detailed information of the specified bus stop code
func BusStop(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") == "application/json" {
		var bs Busstop
		regBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(regBody, &bs)
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("422 - Please enter account details in JSON format"))
		}

		if r.Method == "PUT" {
			// Update tutor's rating for student
			if checkExist(bs) {
				updateValue(bs)
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("202 - Accepted response status code appropriately"))
			} else {
				busstopArray = append(busstopArray, bs)
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("201 - Created success status response code"))
			}

		}
		if r.Method == "DELETE" {
			if checkExist(bs) {
				// remove element from array
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte("202 - Accepted success status response code"))
			} else {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found response status code appropriately"))
			}
		}
	}
}

// create api functions here

func checkExist(bs Busstop) bool {
	results := false
	for _, busstop := range busstopArray {
		if busstop.BusStopCode == bs.BusStopCode {
			results = true
		}
	}
	return results
}

func updateValue(bs Busstop) {
	for _, busstop := range busstopArray {
		if busstop.BusStopCode == bs.BusStopCode {
			busstop.Description = bs.Description
		}
	}
}

func main() {
	// start router
	router := mux.NewRouter()

	// setup routers
	router.HandleFunc("/landing", landing)
	router.HandleFunc("/v1/BusStops", BusStop).Methods("PUT", "DELETE")
	// // establish db connection
	// var err error
	// db, err = sql.Open("mysl", "root:password@tcp(db_sample:8182)/sample_db")
	// if err != il {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// specify allowd headers, methods, & origins to allow CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	fmt.Println("listening at port 5080")
	log.Fatal(http.ListenAndServe(":5080", handlers.CORS(headers, origins, methods)(router)))
}

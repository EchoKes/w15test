package main

type Sample struct {
	id int `json:"id"`
	sample int `json:"sample"`
	datetime string `json:"datetime"`
	boolvalue bool `json"boolValue"`
}

var db *sql.DB

func landing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "~ Sample Landing ~")
}

// create api functions here


// create db api functions here



func main() {
	// start router
	router := mux.NewRouter()

	// setup routers
	router.HandleFunc("/landing", landing)

	// establish db connection
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(db_sample:8182)/sample_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// specify allowed headers, methods, & origins to allow CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	fmt.Println("listening at port 5081")
	log.Fatal(http.ListenAndServe(":5081", handlers.CORS(headers, origins, methods)(router)))
}

package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)
//ArticleHandler is a function handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	//mux.Vars returns all path parameters as a map
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}
//Query-based matching
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"][0])
}

func main() {
	//create a new router
	r := mux.NewRouter()
	//attach a path with handler
	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		//good practice: enforce timeouts for servers you create
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

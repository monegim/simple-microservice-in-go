package cmd

import (
	"github.com/monegim/simple-microservice-in-go/rating/internal/ repository/memory"
	"github.com/monegim/simple-microservice-in-go/rating/internal/controller/rating"
	httphandler "github.com/monegim/simple-microservice-in-go/rating/internal/handler/http"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/paytabs_task/api"
	"github.com/mohamedveron/paytabs_task/service"
	"net/http"
)


func main() {

	serviceLayer := service.NewService()
	server := api.NewServer(serviceLayer)

	fmt.Println("Wait for consuming the accounts")

	serviceLayer.ConsumeAccount()

	fmt.Println("You Can make a transfer now")

	// prepare router
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("server starting on port 8080...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}

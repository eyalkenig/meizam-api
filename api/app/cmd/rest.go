package cmd

import (
	"github.com/eyalkenig/meizam-api/api/app"

	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"

	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func NewRESTCommand(port string) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "rest runs a REST server",
		Run:   getRESTRunFunction(port),
	}
}

func getRESTRunFunction(port string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		router := app.NewRouter()
		allowedOrigins := handlers.AllowedOrigins([]string{"*"})
		allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

		go func() {
			log.Println("Starting application server")
			log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
		}()

		signalsChan := make(chan os.Signal)
		signal.Notify(signalsChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		receivedSignal := <-signalsChan
		log.Println(fmt.Sprintf("Received signal: %s", receivedSignal.String()))
		log.Println("Application server stopped")
	}
}
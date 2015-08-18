package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/spf13/viper"
)

func Run() {
	viper.SetDefault("Port", 6680)
	viper.SetEnvPrefix("newmotion")
	viper.BindEnv("Port")

	viper.SetConfigName("config")
	viper.ReadInConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	app := Application{}

	mux.HandleFunc("/api/transaction", app.TransactionHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	port := viper.GetString("Port")
	if env_port := os.Getenv("PORT"); env_port != "" {
		port = env_port
	}

	address := fmt.Sprintf(":%s", port)
	n.Run(address)
}

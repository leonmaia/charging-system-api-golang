package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/leonmaia/newmotion-golang/keydb"
	"github.com/spf13/viper"
)

func Start() {
	viper.SetDefault("Port", 6680)
	viper.SetEnvPrefix("newmotion")
	viper.BindEnv("Port")

	viper.SetConfigName("config")
	viper.ReadInConfig()

	db := keydb.NewDB()
	app := Application{DB: db}

	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", HealthCheckHandler)
	mux.HandleFunc("/api/transaction", app.TransactionHandler)
	mux.HandleFunc("/api/overview", app.OverviewHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	port := viper.GetString("Port")
	if env_port := os.Getenv("PORT"); env_port != "" {
		port = env_port
	}

	address := fmt.Sprintf(":%s", port)
	n.Run(address)
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
	w.WriteHeader(http.StatusOK)
}

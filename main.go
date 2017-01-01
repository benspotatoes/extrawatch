package main

import (
	"log"
	"net/http"

	"github.com/benspotatoes/extrawatch/api"
	"github.com/benspotatoes/extrawatch/backend"
	"github.com/benspotatoes/extrawatch/config"
	"goji.io"
)

func main() {
	conf := config.NewConfig()

	backendConf := &backend.Config{
		ConnectOpts: conf.Db,
	}
	backend, err := backend.NewBackend(backendConf)
	if err != nil {
		log.Fatalf("unable to initialize backend: %s\n", err.Error())
	}
	mux := api.NewRouter(backend, &api.Config{Cors: conf.Cors})
	serve(conf, mux)
}

func serve(conf *config.Extrawatch, mux *goji.Mux) {
	log.Printf("serving application at %s", conf.HostPort)
	log.Fatal(http.ListenAndServe(conf.HostPort, mux))
}

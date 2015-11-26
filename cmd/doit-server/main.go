package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bryanl/doit-server"
	"github.com/kelseyhightower/envconfig"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/digitalocean"
)

const (
	scope = "read write"
)

// Specification is the runtime specifiation for doit-server. Pertinent environment
// variables will be captured here.
type Specification struct {
	Addr               string `default:":8935"`
	Callback           string
	DigitalOceanKey    string `envconfig:"digitalocean_key"`
	DigitalOceanSecret string `envconfig:"digitalocean_secret"`
	EncodingKey        string `envconfig:"encoding_key"`
}

func main() {
	log.SetPrefix("doit-server: ")
	var s Specification
	err := envconfig.Process("doit", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	if os.Getenv("SESSION_SECRET") == "" {
		log.Fatal("environment variable SESSION_SECRET required")
	}

	if s.DigitalOceanKey == "" {
		log.Fatal("environment variable DOIT_DIGITALOCEAN_KEY required")
	}

	if s.DigitalOceanSecret == "" {
		log.Fatal("environment variable DOIT_DIGITALOCEAN_SECRET required")
	}

	if s.EncodingKey == "" {
		log.Fatal("environment variable DOIT_ENCODING_KEY required")
	}

	goth.UseProviders(
		digitalocean.New(s.DigitalOceanKey, s.DigitalOceanSecret, s.Callback, scope),
	)

	serv := doitserver.NewServer(s.EncodingKey)

	log.Printf("server listening at %s", s.Addr)
	http.ListenAndServe(s.Addr, serv.Mux)
}

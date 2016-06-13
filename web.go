package main

import (
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/owulveryck/simple-framework/server"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Debug       bool
	Scheme      string
	Port        int
	Address     string
	PrivateKey  string
	Certificate string
}

var config Configuration

func main() {

	var help = flag.Bool("help", false, "show help message")
	flag.Parse()
	// Default values
	config.Port = 443
	config.Scheme = "https"
	config.Address = "0.0.0.0"
	config.Debug = false
	config.PrivateKey = "ssl/server.key"
	config.Certificate = "ssl/server.pem"
	defaultConf := config
	err := envconfig.Process("WEB", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("==> WEB_PORT: %v (default: %v)", config.Port, defaultConf.Port)
	log.Printf("==> WEB_SCHEME: %v (default: %v)", config.Scheme, defaultConf.Scheme)
	log.Printf("==> WEB_ADDRESS: %v (default: %v)", config.Address, defaultConf.Address)
	log.Printf("==> WEB_DEBUG: %v (default: %v)", config.Debug, defaultConf.Debug)
	log.Printf("==> WEB_PRIVATEKEY: %v (default: %v)", config.PrivateKey, defaultConf.PrivateKey)
	log.Printf("==> WEB_CERTIFICATE: %v (default: %v)", config.Certificate, defaultConf.Certificate)
	if *help {
		os.Exit(0)
	}

	router := server.NewRouter()

	addr := fmt.Sprintf("%v:%v", config.Address, config.Port)
	if config.Scheme == "https" {
		log.Fatal(http.ListenAndServeTLS(addr, config.Certificate, config.PrivateKey, router))

	} else {
		log.Fatal(http.ListenAndServe(addr, router))

	}
}

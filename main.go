package main

import (
	"flag"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	MaxIdleSeconds int    `yaml:"max_idle_seconds"`
	Domain         string `yaml:"domain"`
	ApiKey         string `yaml:"api_key"`
	PublicApiKey   string `yaml:"public_api_key"`
	Address        string `yaml:"address"`
}

func main() {
	config_file := flag.String("config_file", "/etc/mailgun_router.yaml", "Path to configuration file")

	flag.Parse()

	data, err := ioutil.ReadFile(*config_file)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error in configuration format: %v", err)
	}

	log.Printf("Listening on %s...\n", config.Address)

	sm := NewSmtpRouter(
		&config,
		log.New(os.Stdout, "", log.Ldate|log.Ltime),
		log.New(os.Stderr, "", log.Ldate|log.Ltime))
	if err := sm.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

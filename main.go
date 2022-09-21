package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/kr/pretty"
)

type Config struct {
	Programs map[string]Program `json:"programs"`
}

type Program struct {
	Path         string   `json:"path"`
	Args         []string `json:"args"`
	Description  string   `json:"description"`
	Retries      int      `json:"retries"`
	IgnoreErrors bool     `json:"ignoreErrors"`
	Directory    string   `json:"directory"`
}

func main() {
	// Use a fake "$HOME" for the purposes of this demo
	cfpath := filepath.Join("home", ".config", "demo", "config.json")

	cf, err := os.ReadFile(cfpath)
	if err != nil {
		log.Fatal(err)
	}

	var conf Config
	if err := json.Unmarshal(cf, &conf); err != nil {
		log.Fatal(err)
	}

	pretty.Printf("%# v\n", conf)
}

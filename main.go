package main

import (
	_ "embed"
	"log"
	"path/filepath"

	"github.com/cue-exp/cueconfig"
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

var (
	//go:embed schema.cue
	schema []byte

	//go:embed defaults.cue
	defaults []byte
)

func main() {
	// Use a fake "$HOME" for the purposes of this demo
	cfpath := filepath.Join("home", ".config", "demo", "config.cue")

	var conf Config
	if err := cueconfig.Load(cfpath, schema, defaults, nil, &conf); err != nil {
		log.Fatal(err)
	}

	pretty.Printf("%# v\n", conf)
}

package main

import (
	"fmt"
	"log"

	"github.com/integrii/flaggy"
	"github.com/mlitwiniuk/passha/pkg/config"
	"github.com/mlitwiniuk/passha/pkg/runner"
)

func main() {
	configFile := "config.yml"
	runInline := false
	// timeout := 15
	flaggy.String(&configFile, "c", "config", "Config file")
	flaggy.Bool(&runInline, "i", "inline", "run inline rather than in parallel")
	// flaggy.Int64(&timeout, "t", "timeout", "execution timeout")
	flaggy.Parse()
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Executing command `%s` on defined hosts\n", cfg.Cmd)
	if runInline {
		runner.RunOneByOne(cfg)
	} else {
		runner.RunInParallel(cfg)
	}
}

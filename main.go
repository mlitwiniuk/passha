package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gookit/color"
	"github.com/integrii/flaggy"
	"github.com/mlitwiniuk/passha/pkg/config"
	"github.com/mlitwiniuk/passha/pkg/runner"
)

func main() {
	configFile := "config.yml"
	flaggy.String(&configFile, "c", "config", "Config file")
	flaggy.Parse()
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Executing command `%s` on defined hosts\n", cfg.Cmd)
	for _, host := range cfg.Hosts {

		output, err := runner.RunCmdOnHost(cfg.Cmd, host)
		if err != nil {
			color.Red.Println(strings.ToUpper(host))
			color.Red.Printf("Error: %s\n\n", err)
		} else {
			color.Green.Println(strings.ToUpper(host))
			fmt.Printf("---\n%s--\n\n", output.String())
		}
	}
}

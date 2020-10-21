package runner

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/gookit/color"
	"github.com/mlitwiniuk/passha/pkg/config"
)

type executionResult struct {
	Host   string
	Output string
	Error  error
}

func RunOneByOne(cfg *config.DefaultConfig) {
	for _, host := range cfg.Hosts {
		res := runCmdOnHost(cfg.Cmd, host)
		outputRunnerResult(res)
	}
}

func RunInParallel(cfg *config.DefaultConfig) {
	queue := make(chan executionResult, len(cfg.Hosts))
	var wg sync.WaitGroup
	for _, host := range cfg.Hosts {
		wg.Add(1)
		go runCmdOnHostWithWg(cfg.Cmd, host, queue, &wg)
	}
	wg.Wait()
	close(queue)
	for output := range queue {
		outputRunnerResult(&output)
	}
}

func runCmdOnHostWithWg(command string, host string, queue chan executionResult, wg *sync.WaitGroup) {
	queue <- *runCmdOnHost(command, host)
	wg.Done()
}

func runCmdOnHost(command string, host string) *executionResult {
	cmd := exec.Command("ssh", host, command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return &executionResult{Error: err, Host: host}
	}
	return &executionResult{Output: out.String(), Host: host}
}

func outputRunnerResult(res *executionResult) {
	if res.Error != nil {
		color.Red.Println(strings.ToUpper(res.Host))
		color.Red.Printf("Error: %s\n\n", res.Error)
	} else {
		color.Green.Println(strings.ToUpper(res.Host))
		fmt.Printf("---\n%s--\n\n", res.Output)
	}
}

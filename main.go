package main

import (
	"os"

	"github.com/logrhythm/kafkabeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

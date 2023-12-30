package main

import (
	"os"

	"github.com/ancuongnguyen07/BlurHashGo/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

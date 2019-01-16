package main

import (
	"os"

  "github.com/channelgrabber/protologbeat/cmd"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
    os.Exit(1)
  }
}

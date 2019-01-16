package cmd

import (
  "github.com/figure-of-late/protologbeat/beater"

  cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
var Name = "protologbeat"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, "", beater.New)
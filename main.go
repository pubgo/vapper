package main

import (
	"github.com/pubgo/vapper/cmds"
	"os"
)

func main() {
	cmds.Execute("V", os.ExpandEnv("$PWD"))
}

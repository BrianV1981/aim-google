package main

import (
	"os"

	"github.com/BrianV1981/aim-google/internal/cmd"
	_ "github.com/BrianV1981/aim-google/internal/tzembed" // Embed IANA timezone database for Windows support
)

func main() {
	if err := cmd.Execute(os.Args[1:]); err != nil {
		os.Exit(cmd.ExitCode(err))
	}
}

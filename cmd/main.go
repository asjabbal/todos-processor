package main

import (
	"flag"
	"os"
	"todos-processor/internal"
)

func main() {
	// read flag arguments
	args, err := internal.ArgsReader{}.Read()
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	// start the main app
	App{}.Start(args)
}

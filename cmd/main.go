package main

import (
	"todos-processor/internal"
)

func main() {
	args := internal.ArgsReader{}.Read()
	App{}.Start(args)
}

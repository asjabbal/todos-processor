package internal

import (
	"flag"
	"fmt"
	"os"
	"todos-processor/config"
)

type ArgsReader struct{}

func (ArgsReader) Read() map[string]interface{} {
	typePtr := flag.String("type", "all", "which type of todos to process: all, even, odd")
	sizePtr := flag.Int("size", 10, "number of todos to process at a time: range is 1 to 100")
	flag.Parse()

	fmt.Println("type of todos to process:", *typePtr)
	fmt.Println("size of todos to process:", *sizePtr)

	invalidValue := false
	if *typePtr != config.Constant.TodoTypes["all"] &&
		*typePtr != config.Constant.TodoTypes["even"] &&
		*typePtr != config.Constant.TodoTypes["odd"] {
		fmt.Println("invalid value for type!")
		invalidValue = true
	}
	if *sizePtr < config.Constant.MinTodosSize || *sizePtr > config.Constant.MaxTodosSize {
		fmt.Println("invalid value for size!")
		invalidValue = true
	}
	if invalidValue {
		flag.Usage()
		os.Exit(1)
	}

	args := map[string]interface{}{
		"type": *typePtr,
		"size": *sizePtr,
	}

	return args
}

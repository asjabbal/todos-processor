package internal

import (
	"errors"
	"flag"
	"fmt"
	"todos-processor/config"
)

type ArgsReader struct{}

func (ArgsReader) Read() (map[string]interface{}, error) {
	typePtr := flag.String("type", "all", "which type of todos to process: all, even, odd")
	sizePtr := flag.Int("size", 10, "number of todos to process at a time: range is 1 to 1000")

	flag.Parse()

	fmt.Println("type of todos to process:", *typePtr)
	fmt.Println("size of todos to process:", *sizePtr)

	if *typePtr != config.Constant.TodoTypes["all"] &&
		*typePtr != config.Constant.TodoTypes["even"] &&
		*typePtr != config.Constant.TodoTypes["odd"] {
		return nil, errors.New("invalid value for type!")
	}
	if *sizePtr < config.Constant.MinTodosSize || *sizePtr > config.Constant.MaxTodosSize {
		fmt.Println("invalid value for size!")
		return nil, errors.New("invalid value for size!")
	}

	args := map[string]interface{}{
		"type": *typePtr,
		"size": *sizePtr,
	}

	return args, nil
}

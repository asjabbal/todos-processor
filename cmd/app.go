package main

import (
	"fmt"
	"os"
	"todos-processor/internal"
)

type App struct{}

func (App) Start(options map[string]interface{}) {
	fmt.Println("processing started")

	// generate ids based on odd/even and size
	ids, err := internal.IdsGenerator{}.Generate(options["type"].(string), options["size"].(int))
	if err != nil {
		fmt.Printf("unable to generate ids due to an error: %v\n", err)
		os.Exit(1)
	}

	// execute the process of fetching and processing todos
	Processor{}.Execute(ids)
}

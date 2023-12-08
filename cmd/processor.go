package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"todos-processor/api"
	"todos-processor/model"
)

type Processor struct{}

func (Processor) Execute(ids []int) {
	fmt.Println("processing ids =", ids)

	var wg sync.WaitGroup

	for i := 0; i < len(ids); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			responseData, err := api.TodosApi{}.FetchById(id)
			if err != nil {
				fmt.Printf("unable to fetch todo with id=%v due to an error: %v\n", id, err)
				return
			}

			todo := model.Todo{}
			err = json.Unmarshal(responseData, &todo)
			if err != nil {
				fmt.Printf("unable to parse response data of todo with id=%v due to an error: %v\n", id, err)
				return
			}

			todo.PrintValues()
		}(ids[i])
	}

	wg.Wait()
}

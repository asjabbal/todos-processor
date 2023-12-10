package internal

import (
	"errors"
	"todos-processor/config"
)

type IdsGenerator struct{}

func (IdsGenerator) Generate(todosType string, numberOfTodos int) ([]int, error) {
	var ids []int

	if numberOfTodos < config.Constant.MinTodosSize || numberOfTodos > config.Constant.MaxTodosSize {
		return nil, errors.New("invalid todos size given!")
	}

	switch todosType {
	case config.Constant.TodoTypes["all"]:
		ids = generateAllIds(numberOfTodos)
	case config.Constant.TodoTypes["even"]:
		ids = generateEvenIds(numberOfTodos)
	case config.Constant.TodoTypes["odd"]:
		ids = generateOddIds(numberOfTodos)
	default:
		return nil, errors.New("invalid todos type given!")
	}

	return ids, nil
}

func generateAllIds(size int) []int {
	var ids []int
	for i := 1; i <= size; i++ {
		ids = append(ids, i)
	}

	return ids
}

func generateEvenIds(size int) []int {
	return generateEvenOrOddIds("even", size)
}

func generateOddIds(size int) []int {
	return generateEvenOrOddIds("odd", size)
}

func generateEvenOrOddIds(idType string, size int) []int {
	ref := 2
	if idType == config.Constant.TodoTypes["odd"] {
		ref = 1
	}

	var ids []int
	for i := 0; i < size; i++ {
		ids = append(ids, ref)
		ref = ref + 2
	}

	return ids
}

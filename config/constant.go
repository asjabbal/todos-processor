package config

type constant struct {
	MinTodosSize    int
	MaxTodosSize    int
	TodoTypes       map[string]string
	TodosApiHostUrl string
}

var Constant = constant{
	MinTodosSize: 1,
	MaxTodosSize: 1000,
	TodoTypes: map[string]string{
		"all":  "all",
		"even": "even",
		"odd":  "odd",
	},
	TodosApiHostUrl: "https://jsonplaceholder.typicode.com/todos",
}

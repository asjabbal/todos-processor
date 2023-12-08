# Todos Processor

## Requirements:-
- Golang 1.21.5
- Docker (if running project with Docker)

## How to run without Docker:-
- Go to project's root folder
- Run `build/todos-processor -type=even -size=20`

## How to run with Docker:-
- Go to project's root folder
- Run `docker build -t todos-processor .`
- Then `docker run -it --rm --name todos-processor-container todos-processor`
- It will open up `bash` in Docker container
- Now run command `todos-processor -type=even -size=20`

## How to use `todos-processor` command:-
- For help, run `todos-processor -h` and it will give you usage info:-
```
Usage of todos-processor:
  -size int
    	number of todos to process at a time: range is 1 to 1000 (default 10)
  -type string
    	which type of todos to process: all, even, odd (default "all")
```
- To process only first 20 even todos, run `todos-processor -type=even -size=20`
- To process only first 30 odd todos, run `todos-processor -type=odd -size=30`
- To process only first 50 all todos, run `todos-processor -type=all -size=50`
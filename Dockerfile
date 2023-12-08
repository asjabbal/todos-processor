FROM golang:1.21

COPY . /todos-processor

WORKDIR /todos-processor

RUN go mod tidy

WORKDIR /todos-processor/cmd

RUN go build -o ../build/todos-processor

WORKDIR /todos-processor

RUN cp build/todos-processor /usr/local/bin

CMD [ "bash"]
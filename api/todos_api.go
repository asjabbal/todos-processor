package api

import (
	"io"
	"net/http"
	"strconv"
	"todos-processor/config"
)

type TodosApi struct{}

var hostUrl string = config.Constant.TodosApiHostUrl

func (TodosApi) FetchById(id int) ([]byte, error) {
	response, err := http.Get(hostUrl + "/" + strconv.Itoa(id))

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

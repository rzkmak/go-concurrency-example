package main

import (
	"encoding/json"
	"fmt"
	"github.com/aeidelos/go-concurrency/domain"
	"io/ioutil"
	"net/http"
)

func main() {
	todoChannel := make(chan domain.Todo)
	var todo domain.Todo
	go func() {
		result := domain.Todo{}
		CallApiConcurrent("https://jsonplaceholder.typicode.com/todos/1", &result)
		todoChannel <- result
	}()
	todo = <- todoChannel
	fmt.Println(todo)
}

func CallApiConcurrent(url string, data interface{}) {
	fmt.Println("Starting api call on " + url)
	response, err := http.Get(url)
	if err != nil {
		panic("error api call")
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			panic("error closing body request")
		}
	}()
	body := response.Body
	result, err := ioutil.ReadAll(body)
	if err != nil {
		panic("error parsing")
	}
	fmt.Println(string(result))
	errMar := json.Unmarshal(result, &data)
	if errMar != nil {
		panic("error unmarshal json")
	}
}

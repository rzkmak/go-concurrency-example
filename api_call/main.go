package main

import (
	"encoding/json"
	"fmt"
	"github.com/aeidelos/go-concurrency/domain"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Starting the application")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic("error api call")
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			panic("error closing body request")
		}
	}()
	data := response.Body
	result, err := ioutil.ReadAll(data)
	if err != nil {
		panic("error parsing")
	}
	fmt.Println(string(result))
	todo := domain.Todo{}
	errMar := json.Unmarshal(result, &todo)
	if errMar != nil {
		panic("error unmarshal json")
	}
	fmt.Println("Todo parsing json")
	fmt.Println(todo)

}

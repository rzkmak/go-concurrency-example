package main

import (
	"encoding/json"
	"fmt"
	"github.com/aeidelos/go-concurrency/domain"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	var todo [] domain.Todo
	var wg sync.WaitGroup
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			result := domain.Todo{}
			CallApiConcurrent("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id), &result, &wg)
			todo = append(todo, result)
		}(i)
	}

	wg.Wait()
	fmt.Println(todo)
}

func CallApiConcurrent(url string, data interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
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

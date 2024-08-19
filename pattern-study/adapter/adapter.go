package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ToDo is a struct that represents a todo item
type ToDo struct {
	UserID    int    `json:"userId" xml:"userId"`
	ID        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

// DataInterface is an adapter that adapts the data to the ToDo struct
type DataInterface interface {
	GetData() (*ToDo, error)
}

// RemoteService is a struct that represents a remote service
type RemoteService struct {
	Remote DataInterface
}

// CallRemoteService calls the remote service
func (r *RemoteService) CallRemoteService() (*ToDo, error) {
	return r.Remote.GetData()
}

// JSONBackend is a struct that represents a JSON backend
type JSONBackend struct{}

// GetData gets the data from the JSON backend
func (j *JSONBackend) GetData() (*ToDo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo ToDo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// XMLBackend is a struct that represents an XML backend
type XMLBackend struct{}

// GetData gets the data from the XML backend
func (x *XMLBackend) GetData() (*ToDo, error) {
	xmlFile := `<?xml version="1.0" encoding="UTF-8"?>
	<todo>
		<userId>1</userId>
		<id>1</id>
		<title>delectus aut autem</title>
		<completed>false</completed>
	</todo>`

	var todo ToDo
	err := xml.Unmarshal([]byte(xmlFile), &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func main() {
	// No adapter
	todo := getRemoteDta()
	fmt.Printf("TODO without adapter: %+v\n", todo)

	// With adapter, using JSON
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	tdFromJSON, err := jsonAdapter.CallRemoteService()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("TODO with JSON adapter: %+v\n", tdFromJSON)

	// With adapter, using XML
	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	tdFromXML, err := xmlAdapter.CallRemoteService()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("TODO with XML adapter: %+v\n", tdFromXML)
}

func getRemoteDta() *ToDo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo ToDo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatalln(err)
	}

	return &todo
}

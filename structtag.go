package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Person some comment
type Person struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name"`
}

func main() {
	p := &Person{FirstName: "Zhihui", LastName: "Tang"}
	j, _ := json.Marshal(p)
	fmt.Println(string(j))

	t := reflect.TypeOf(p)
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
}

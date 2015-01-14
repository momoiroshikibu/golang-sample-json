package main

import (
  "encoding/json"
  "net/http"
)

type Person struct {
    Name string
    Age int
}

func NewPerson (name string, age int) *Person {
    person := new(Person)
    person.Name = name
    person.Age = age
    return person
}

func main() {
  http.HandleFunc("/", getPersonalInfo)
  http.ListenAndServe(":3000", nil)
}

func getPersonalInfo(w http.ResponseWriter, r *http.Request) {
  person := NewPerson("Kosuke", 28)

  json_response, err := json.Marshal(person)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(json_response)
}

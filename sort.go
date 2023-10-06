package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserSlice []User

func (value UserSlice) Len() int { return len(value) }
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
func main() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 23},
		{"Habi", 19},
		{"Wahyu", 28},
	}

	sort.Sort(UserSlice(users))

	jsons, err := json.Marshal(users)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsons))
}

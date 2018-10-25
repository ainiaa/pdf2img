package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person Person结构
type Person struct {
	ID        string   `json:"id,omitemty"`         //id
	Firstname string   `json:"firstname,omitempty"` //姓
	Lastname  string   `json:"lastname,omitempty"`  //名字
	Address   *Address `json:"address,omitempty"`   //地址
}

//Address 地址结构
type Address struct {
	City     string `json:"city,omitempty"`     //城市
	Province string `json:"province,omitempty"` //省份
}

var people []Person

// GetPerson 获取用户信息
func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

// GetPeople 获取用户信息
func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// PostPerson 更新用户信息
func PostPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson 删除用户信息
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main4() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "xi", Lastname: "dada", Address: &Address{City: "Shenyang", Province: "Liaoning"}})
	people = append(people, Person{ID: "2", Firstname: "li", Lastname: "xiansheng", Address: &Address{City: "Changchun", Province: "Jinlin"}})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", PostPerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}

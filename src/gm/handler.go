package gm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := Users{
		User{Name:"胡兵", Sex :"male", Level:23},
		User{Name:"游族", Sex :"female", Level:999},
		}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ChoreIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(chores); err != nil {
		panic(err)
	}
}

func ChoreShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var choreId int
	var err error
	if choreId, err = strconv.Atoi(vars["choreId"]); err != nil {
		panic(err)
	}
	
	chore := RepoFindChore(choreId)
	if chore.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(chore); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/chores
*/
func ChoreCreate(w http.ResponseWriter, r *http.Request) {
	var chore Chore
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	
	if err := json.Unmarshal(body, &chore); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateChore(chore)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
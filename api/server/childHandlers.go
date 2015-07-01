package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func ChildrenIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(children); err != nil {
		panic(err)
	}
}

func ChildShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var childId int
	var err error
	if childId, err = strconv.Atoi(vars["childId"]); err != nil {
		panic(err)
	}
	
	child := RepoFindChild(childId)
	if child.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(child); err != nil {
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
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
func ChildCreate(w http.ResponseWriter, r *http.Request) {
	var child Child
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &child); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateChild(child)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func ChildShowChores(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var childId int
	var err error
	if childId, err = strconv.Atoi(vars["childId"]); err != nil {
		panic(err)
	}
	
	child := RepoFindChild(childId)
	if child.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(child.Chores); err != nil {
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

func ChildAddChore(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var childId int
	var choreId int
	var err error
	
	if childId, err = strconv.Atoi(vars["childId"]); err != nil {
		panic(err)
	}
	if choreId, err = strconv.Atoi(vars["choreId"]); err != nil {
		panic(err)
	}

	child := RepoFindChild(childId)
	if child.Id > 0 {
		RepoAddChoreToChild(child.Id, choreId)
		fmt.Fprint(w, child.Chores)
		
		
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(child.Chores); err != nil {
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
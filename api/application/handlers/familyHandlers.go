package handlers

import (
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/dicknaniel/Chillowance/api/application/repositories"
	"github.com/dicknaniel/Chillowance/api/application/models"
)

func Family_Create(w http.ResponseWriter, r *http.Request) {
	var family models.Family

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &family); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := repositories.RepoCreateFamily(family)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}



package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Member struct {
	ID          string
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Bank        string `json:"bank"`
}

var MemberDatabase = map[string]Member{}

func CreateMember(w http.ResponseWriter, r *http.Request) {
	var m Member
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	m.ID = uuid.New().String()
	MemberDatabase[m.ID] = m
	JsonResponse(w, m)
}

func GetMembers(w http.ResponseWriter, r *http.Request) {
	var ml []Member
	for _, v := range MemberDatabase {
		ml = append(ml, v)
	}
	JsonResponse(w, ml)
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mID := vars["memberId"]
	m, ok := MemberDatabase[mID]
	if !ok {
		ErrorResponse(w, fmt.Errorf("User Not Found"), http.StatusNotFound)
		return
	}
	JsonResponse(w, m)
}

func UpdateMembers() {

	// fetch a member by their user id and change the required parameter, return the result and save it

}

func DeleteMembers() {

	// delete a member for our database

}

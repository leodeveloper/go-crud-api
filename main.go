package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Education struct {
	ID      string   `json:"id"`
	Degree  string   `json:"degree"`
	GPA     string   `json:"gpa"`
	Student *Student `json:"student"`
}

type Student struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var educations []Education

func main() {
	r := mux.NewRouter()

	educations = append(educations, Education{ID: "1", Degree: "Bachelor", GPA: "3.1", Student: &Student{Firstname: "Muhammad", Lastname: "Suleman"}})
	educations = append(educations, Education{ID: "2", Degree: "Master", GPA: "3.2", Student: &Student{Firstname: "Muhammad", Lastname: "Noman"}})

	r.HandleFunc("/educations", getEducations).Methods("GET")
	r.HandleFunc("/educations/{id}", getEducation).Methods("GET")
	r.HandleFunc("/educations", createEducation).Methods("POST")
	r.HandleFunc("/educations/{id}", updateEducation).Methods("PUT")
	r.HandleFunc("/educations/{id}", deleteEducation).Methods("DELETE")

	fmt.Printf("Starting server at port 8081\n")
	log.Fatal(http.ListenAndServe(":8081", r))

}

func getEducations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(educations)
}

func deleteEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range educations {
		if item.ID == params["id"] {
			educations = append(educations[:index], educations[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(educations)
}

func getEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//we do not need index that why we use blank _
	for _, item := range educations {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var education Education
	_ = json.NewDecoder(r.Body).Decode(&education)
	education.ID = strconv.Itoa(rand.Intn(100000000))
	educations = append(educations, education)
	json.NewEncoder(w).Encode(education)
}

func updateEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range educations {
		if item.ID == params["id"] {
			educations = append(educations[:index], educations[index+1:]...)
			var education Education
			_ = json.NewDecoder(r.Body).Decode(&education)
			education.ID = params["id"]
			educations = append(educations, education)
			json.NewEncoder(w).Encode(education)
			return
		}
	}
}

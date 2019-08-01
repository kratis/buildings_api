package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "log"
    "strconv"
    "net/http"
)

type Address struct {
    City  string `json:"city"`
    State string `json:"state"`
}

type Building struct {
    Id string `json:"id,omitempty"`
    Address *Address `json:"address,omitempty"`
    Floors []int `json:"floors,omitempty"`
}

var buildings[]Building


func index(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(buildings)
}

func show(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(req)
    for _, entry := range buildings {
        if entry.Id == params["id"] {
            json.NewEncoder(w).Encode(entry)
            return
        }
    }

    json.NewEncoder(w).Encode(&Building{})
}

func create(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var newBuilding Building
    _ = json.NewDecoder(req.Body).Decode(&newBuilding)
    newBuilding.Id = strconv.Itoa(len(buildings) + 1)
    
    buildings = append(buildings, newBuilding)
    json.NewEncoder(w).Encode(buildings)
}

func delete(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(req)
    for index, entry := range buildings {
        if entry.Id == params["id"] {
            buildings = append(buildings[:index], buildings[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(buildings)
}

func main() {
    router := mux.NewRouter()
    buildings = append(buildings, Building{Id: "1", Floors: []int{1,2,3}})
    buildings = append(buildings, Building{Id: "2", Floors: []int{1,2,3,4}})
    router.HandleFunc("/buildings", index).Methods("GET")
    router.HandleFunc("/buildings/{id}", show).Methods("GET")
    router.HandleFunc("/buildings", create).Methods("POST")
    router.HandleFunc("/buildings/{id}", delete).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345", router))
}

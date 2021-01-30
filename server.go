package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func server() {
	handlers := newHandlers()
	http.HandleFunc("/random", handlers.getRandom)
	http.HandleFunc("/portada/", handlers.getPortada)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

type handlers struct {
	sync.Mutex
	store map[int]Portada
}

func newHandlers() *handlers {
	return &handlers{
		store: extractData(),
	}
}

func (h *handlers) getPortada(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Wrong number of elements in request. Expected '/portada/$id, but found %v'", r.URL.String())
		return
	}
	portadaId, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal(err)
	}

	h.Lock()
	portada, ok := h.store[portadaId]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("Portada with id: %v not found.\n", portadaId)
		return
	}

	jsonBytes, err := json.Marshal(portada)
	if err != nil {
		log.Printf(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonBytes)
	log.Printf("Portada with id: %v served.\n", portadaId)
}

func (h *handlers) getRandom(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	h.Lock()
	portada := h.store[rand.Intn(len(h.store))]
	h.Unlock()

	jsonBytes, err := json.Marshal(portada)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonBytes)
	log.Println("Random portada served.")
}

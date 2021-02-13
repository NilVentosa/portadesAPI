package portades

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func Server() {
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
	data *Data
}

func newHandlers() *handlers {
	db, err := sql.Open("sqlite3", "portades.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	return &handlers{
		data: NewData(db),
	}
}

func (h *handlers) getPortada(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 || len(parts[2]) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error"))
		log.Printf("Wrong number of elements in request. Expected '/portada/$id, but found %v'", r.URL.String())
		return
	}

	portadaId, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal(err)
	}

	h.Lock()
	portada, ok := h.data.GetPortada(portadaId)
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

	h.Lock()
	portada, ok := h.data.GetRandomPortada()
	h.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error"))
		return
	}

	jsonBytes, err := json.Marshal(portada)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonBytes)
	log.Println("Random portada served.")
}

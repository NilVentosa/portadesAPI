package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var dataFile string = "./portades.csv"

type Portada struct {
	Id        int    `json:"id"`
	Intro     string `json:"intro"`
	Newspaper string `json:"newspaper"`
	Headline  string `json:"headline"`
	Result    bool   `json:"result"`
	Video     string `json:"video"`
	Episode   string `json:"episode"`
}

func main() {
	server()
}

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

func extractData() map[int]Portada {
	_, currentFileName, _, _ := runtime.Caller(0)
	filePath := path.Dir(currentFileName)

	csvFile, _ := os.Open(filePath + "/" + dataFile)
	reader := csv.NewReader(csvFile)

	var portades map[int]Portada
	portades = make(map[int]Portada)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}

		result, err := strconv.ParseBool(line[4])
		if err != nil {
			log.Fatal(err)
		}

		portades[id] = Portada{
			Id:        id,
			Intro:     line[1],
			Headline:  line[3],
			Newspaper: line[2],
			Result:    result,
			Video:     line[5],
			Episode:   line[6],
		}
	}

	return portades
}

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
	"time"
)

var dataFile string = "./portades.csv"

type Portada struct {
	Id        int    `json:"id"`
	Intro     string `json:"intro"`
	Newspaper string `json:"newspaper"`
	Headline  string `json:"headline"`
	Result    bool   `json:"result"`
	Image     string `json:"image"`
}

func main() {
	server()
}

func server() {
	randomHandlers := newRandomHandlers()
	http.HandleFunc("/random", randomHandlers.get)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

type randomHandlers struct {
	store map[int]Portada
}

func newRandomHandlers() *randomHandlers {
	return &randomHandlers{
		store: extractData(),
	}
}

func (h *randomHandlers) get(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	portada := h.store[rand.Intn(len(h.store))]

	jsonBytes, err := json.Marshal(portada)
	if err != nil {
		log.Fatal(err.Error())
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	//w.Header().Add("Access-Control-Allow-Origin", "*")
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
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(error)
		}

		result, err := strconv.ParseBool(line[4])
		if err != nil {
			log.Fatal(error)
		}

		portades[id] = Portada{
			Id:        id,
			Intro:     line[1],
			Headline:  line[3],
			Newspaper: line[2],
			Result:    result,
			Image:     line[5],
		}
	}

	return portades
}

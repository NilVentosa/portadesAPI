package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var dataFile string = "portades.csv"

type Portada struct {
	Id       int    `json:"id"`
	Intro    string `json:"intro"`
	Headline string `json:"headline"`
	Result   bool   `json:"result"`
	Image    string `json:"image"`
}

func main() {
	data := extractData()
}

func extractData() map[int]Portada {
	csvFile, _ := os.Open(dataFile)
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

		result, err := strconv.ParseBool(line[3])
		if err != nil {
			log.Fatal(error)
		}

		portades[id] = Portada{
			Id:       id,
			Intro:    line[1],
			Headline: line[2],
			Result:   result,
			Image:    line[4],
		}
	}

	portadesJson, _ := json.Marshal(portades)
	fmt.Println(string(portadesJson))
	return portades
}

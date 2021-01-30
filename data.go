package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
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

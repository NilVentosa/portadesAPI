package main

import (
	"os"
	"portadesAPI/platform/portades"
)

func main() {
	portades.Server(os.Args[1])
}

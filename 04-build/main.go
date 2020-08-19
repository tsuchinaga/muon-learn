package main

import (
	"github.com/ImVexed/muon"
	"log"
	"net/http"
	"os"
)

func main() {
	m := muon.New(&muon.Config{Title: "Quit Button", Titled: true, Width: 500, Height: 500, Borderless: true}, http.FileServer(http.Dir("./www")))
	m.Bind("quit", func() { os.Exit(0) })

	if err := m.Start(); err != nil {
		log.Println(err)
	}
}

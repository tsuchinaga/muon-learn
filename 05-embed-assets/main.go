package main

import (
	"github.com/ImVexed/muon"
	"log"
	"net/http"
)

func main() {
	m := muon.New(&muon.Config{Title: "Embed", Titled: true, Width: 500, Height: 500}, http.FileServer(http.Dir("./www")))
	if err := m.Start(); err != nil {
		log.Println(err)
	}
}

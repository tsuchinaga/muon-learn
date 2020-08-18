package main

import (
	"github.com/ImVexed/muon"
	"net/http"
)

func main() {
	m := muon.New(&muon.Config{
		Title:      "Hello, World",
		Height:     500,
		Width:      500,
		Resizeable: true,
		Borderless: true,
		Titled:     true,
	}, http.FileServer(http.Dir("./www")))

	if err := m.Start(); err != nil {
		panic(err)
	}
}

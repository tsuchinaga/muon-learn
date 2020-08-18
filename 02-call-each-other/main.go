package main

import (
	"github.com/ImVexed/muon"
	"log"
	"net/http"
	"sync"
)

func main() {
	s := store{count: 1}
	m := muon.New(&muon.Config{Title: "Counter", Titled: true, Width: 500, Height: 500}, http.FileServer(http.Dir("./www")))
	m.Bind("getCount", func() float64 { return float64(s.get()) })
	m.Bind("countUp", func() {
		s.countUp()
		if _, err := m.Eval("rerender()", nil); err != nil {
			log.Println(err)
		}
	})
	m.Bind("countDown", func() {
		s.countDown()
		if _, err := m.Eval("rerender()", nil); err != nil {
			log.Println(err)
		}
	})

	if err := m.Start(); err != nil {
		log.Println(err)
	}
}

type store struct {
	count int
	mutex sync.Mutex
}

func (s *store) get() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.count
}

func (s *store) countUp() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.count++
}

func (s *store) countDown() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.count--
}

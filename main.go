package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	//go:embed html/last.html
	lastHTML string
)

type Server struct {
	db *DB
}

// GET /last -> HTML
func (s *Server) lastHandler(w http.ResponseWriter, r *http.Request) {
	lastText := "No entries"

	e, err := s.db.Last()
	if err == nil {
		time := e.Time.Format("2006-01-02T15:04")
		lastText = fmt.Sprintf("[%s] %s by %s", time, e.Content, e.Login)
	}
	fmt.Fprintf(w, lastHTML, lastText)
}

// POST /
func (s *Server) newHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var e Entry

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	e.Time = time.Now()
	err := s.db.Add(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := map[string]interface{}{
		"error": nil,
		"time":  e.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) Health() error {
	return s.db.Health()
}

// Get /health
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.Health(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "OK\n")
}

func main() {
	var err error
	dsn := os.Getenv("JOURNAL_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=s3cr3t sslmode=disable"
	}
	db, err := NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	s := Server{db}

	r := mux.NewRouter()
	r.HandleFunc("/last", s.lastHandler).Methods("GET")
	r.HandleFunc("/api/journal", s.newHandler).Methods("POST")
	r.HandleFunc("/api/health", s.healthHandler).Methods("GET")

	http.Handle("/", r)

	addr := os.Getenv("JOURNAL_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

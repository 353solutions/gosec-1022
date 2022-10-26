package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	//go:embed html/last.html
	lastHTML     string
	lastTemplate *template.Template
)

type Server struct {
	db *DB
}

// GET /last -> HTML
func (s *Server) lastHandler(w http.ResponseWriter, r *http.Request) {
	e, err := s.db.Last(r.Context())
	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			http.Error(w, "can't query", http.StatusInternalServerError)
			return
		}

		e = Entry{
			Login:   "no one",
			Content: "nothing",
		}
	}
	lastTemplate.Execute(w, e)
}

// Home exercise: write a size limit middleware
// Hint: io.NopCloser
const maxMsgSize = 1 << 20

// POST /
func (s *Server) newHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var e Entry

	lr := http.MaxBytesReader(w, r.Body, maxMsgSize)
	if err := json.NewDecoder(lr).Decode(&e); err != nil {
		// exercise: send hand-crafted errors
		http.Error(w, "bad JSON", http.StatusBadRequest)
		return
	}

	e.Time = time.Now()

	if err := e.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := s.db.Add(r.Context(), e)
	if err != nil {
		http.Error(w, "can't store entry", http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"error": nil,
		"time":  e.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *Server) Health(ctx context.Context) error {
	return s.db.Health(ctx)
}

// Get /health
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.Health(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "OK\n")
}

func main() {
	var err error
	lastTemplate, err = template.New("last").Parse(lastHTML)
	if err != nil {
		log.Fatalf("ERROR: can't parse HTML - %s", err)
	}

	dsn := os.Getenv("JOURNAL_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=s3cr3t sslmode=disable"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	db, err := NewDB(ctx, dsn)
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

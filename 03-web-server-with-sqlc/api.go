package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/uguremirmustafa/wswsqlc/db"
)

type APIServer struct {
	listenAddr string
	db         *Storage
}

func NewApiServer(addr string, storage *Storage) *APIServer {
	return &APIServer{
		listenAddr: addr,
		db:         storage,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.handleGetNodes)
	mux.HandleFunc("POST /", s.handleCreateNode)

	log.Printf("JSON API Server running on port: %s \n", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, mux)
	return err
}

func (s *APIServer) handleGetNodes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /")

	nodes, err := s.db.q.ListNodes(r.Context())

	if err != nil {
		fmt.Fprintf(w, "fucked up")
	}

	WriteJSON(w, http.StatusOK, nodes)
}

func (s *APIServer) handleCreateNode(w http.ResponseWriter, r *http.Request) {
	req := new(db.CreateNodeParams)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		WriteJSON(w, http.StatusOK, "fucked up while parsing")
		return
	}

	arg := db.CreateNodeParams{
		Name:        req.Name,
		Description: req.Description,
	}

	node, err := s.db.q.CreateNode(r.Context(), arg)
	if err != nil {
		WriteJSON(w, http.StatusOK, fmt.Sprint(err.Error()))
		return
	}

	WriteJSON(w, http.StatusOK, node)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

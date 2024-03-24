package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type APIServer struct {
	listenAddr string
	db         Storage
}

func NewApiServer(addr string, storage Storage) *APIServer {
	return &APIServer{
		listenAddr: addr,
		db:         storage,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /product", (makeHTTPHandleFunc(s.handleGetProducts)))
	mux.HandleFunc("POST /product", (makeHTTPHandleFunc(s.handleCreateProduct)))
	mux.HandleFunc("PUT /product/{id}", (makeHTTPHandleFunc(s.handleUpdateProduct)))
	mux.HandleFunc("DELETE /product/{id}", (makeHTTPHandleFunc(s.handleDeleteProduct)))
	log.Printf("JSON API Server running on port: %s \n", s.listenAddr)

	http.ListenAndServe(s.listenAddr, mux)
}

func (s *APIServer) handleGetProducts(w http.ResponseWriter, r *http.Request) error {
	products, err := s.db.GetProducts()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, products)
}

func (s *APIServer) handleCreateProduct(w http.ResponseWriter, r *http.Request) error {
	productReq := new(CreateProductRequest)
	if err := json.NewDecoder(r.Body).Decode(productReq); err != nil {
		return err
	}

	p := NewProduct(productReq.Name, productReq.Price)
	id, err := s.db.CreateProduct(p)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, id)
}

func (s *APIServer) handleUpdateProduct(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	productReq := new(UpdateProductRequest)
	if err := json.NewDecoder(r.Body).Decode(productReq); err != nil {
		return err
	}

	p := NewProduct(productReq.Name, productReq.Price)
	err = s.db.UpdateProduct(p, id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, ApiSuccess{Success: "update is successful"})
}

func (s *APIServer) handleDeleteProduct(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	err = s.db.DeleteProduct(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, ApiSuccess{Success: fmt.Sprintf("product %d is deleted", id)})
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type ApiSuccess struct {
	Success string `json:"success"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func getID(r *http.Request) (int64, error) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid id given %s", idStr)
	}

	return int64(id), nil
}

package handler

import (
	"crud_chi_with_text/interfaces"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type FileHandler struct {
	service interfaces.FileService
}

func NewFileHandler(service interfaces.FileService) *FileHandler {
	return &FileHandler{
		service: service,
	}
}

func (t *FileHandler) CreateDataHandler(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	t.service.CreateData(data)
	w.WriteHeader(http.StatusCreated)
}

func (t *FileHandler) ReadAllDataHandler(w http.ResponseWriter, r *http.Request) {
	data := t.service.ReadAllData()
	for _, d := range data {
		fmt.Fprintln(w, d)
	}
}

func (t *FileHandler) FindByIDHandler(w http.ResponseWriter, r *http.Request) {
	targetID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := t.service.FindByID(targetID)
	if result == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, result)
}

func (t *FileHandler) UpdateDataHandler(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newData := r.FormValue("data")
	t.service.UpdateData(index, newData)
	w.WriteHeader(http.StatusOK)
}

func (t *FileHandler) DeleteDataHandler(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.service.DeleteData(index)
	w.WriteHeader(http.StatusOK)
}

package main

import (
	"crud_chi_with_text/handler"
	"crud_chi_with_text/repository"
	"crud_chi_with_text/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	repository := repository.NewTextFileRepository("file.txt")
	service := service.NewTextFileService(repository)
	textHandler := handler.NewFileHandler(service)

	r := chi.NewRouter()
	r.Post("/create", textHandler.CreateDataHandler)
	r.Get("/", textHandler.ReadAllDataHandler)
	r.Get("/{id}", textHandler.FindByIDHandler)
	r.Put("/update/{id}", textHandler.UpdateDataHandler)
	r.Delete("/delete/{id}", textHandler.DeleteDataHandler)

	http.ListenAndServe(":8080", r)
}

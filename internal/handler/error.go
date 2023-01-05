package handler

import (
	"fmt"
	"net/http"
	"text/template"

	"forum/models"
)

func (h *Handler) ErrorPage(w http.ResponseWriter, status int, message string) {
	errData := models.Error{Status: status, StatusText: http.StatusText(status), Message: message}
	templ, err := template.ParseFiles("./ui/error.html")
	if err != nil {
		fmt.Printf("error handler: parsefiles: %s\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if err := templ.Execute(w, errData); err != nil {
		fmt.Printf("error handler: execute: %s\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

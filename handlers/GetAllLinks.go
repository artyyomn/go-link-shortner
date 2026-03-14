package handlers

import (
	"log"
	"net/http"
	"url-shortner/models"
)

func GetAllLinks(w http.ResponseWriter, r *http.Request) {
	for key, value := range models.Database {
		log.Println(key, value)
	}
}

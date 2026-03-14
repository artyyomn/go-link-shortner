package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"url-shortner/db"
)

func HandleShortLink(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Write([]byte("wrong request method lil bro"))
		return
	}

	shortURL := r.PathValue("short")

	// value, ok := models.Database[shortURL]
	// TODO refactor this for sqlite database

	longLink, ok := QueryDb(shortURL)
	log.Println(longLink, shortURL)

	if ok {
		http.Redirect(w, r, longLink, http.StatusTemporaryRedirect)
	} else {
		w.Write([]byte("invalid URL"))
	}
}

func QueryDb(shortLink string) (string, bool) {
	var long_link string

	sqldb = db.New()
	defer sqldb.Close()

	err := sqldb.QueryRow("SELECT long_link FROM links WHERE short_link = ?", shortLink).Scan(&long_link)
	if err == sql.ErrNoRows {
		return "", false
	} else if err != nil {
		log.Fatal("couldn't  query row", err)
		return "", false
	}

	return long_link, true
}

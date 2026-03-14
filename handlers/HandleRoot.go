package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"url-shortner/db"
	"url-shortner/models"
	"html/template"
)

var sqldb *sql.DB

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./frontend/index.html"))
		tmpl.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		var URL models.Link

		err := json.NewDecoder(r.Body).Decode(&URL)
		if err != nil {
			log.Fatal("json encoding failed", err)
		}

		// TODO: check legibility of the url???
		// URL.Link = SanitizeLink(URL.Link)

		URL.Short = shorten(URL.Long)
		StoreInDb(URL)

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(URL.Short); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}
}

func StoreInDb(links models.Link) error {
	sqldb = db.New()
	defer sqldb.Close()

	tx, err := sqldb.Begin()
	if err != nil {
		log.Fatal("couldn't begin transaction", err)
	}

	stmt, err := tx.Prepare("INSERT INTO links(long_link, short_link) VALUES(?, ?)")
	if err != nil {
		log.Fatal("Insettion failed", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(links.Long, links.Short)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal("Commit failed :(", err)
	}

	return nil
}

func shorten(url string) string {
	hash := sha256.Sum256([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:])[:8]
}

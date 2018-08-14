package controller

import (
	// "encoding/json"
	// "fmt"
	"net/http"

	"github.com/voyagegroup/go-todo/model"

	"github.com/jmoiron/sqlx"
)

type Search struct {
	DB *sqlx.DB
}

func (s *Search) Get(w http.ResponseWriter, r *http.Request) error {

	query := r.Form["title"]
	title := query[0]
	if len(query) != 1 {
		title = ""
	}
	todos, err := model.Search(s.DB, title)
	if err != nil {
		return err
	}
	return JSON(w, 200, todos)
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dao

import (
	null "github.com/guregu/null/v5"
)

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID          int64       `json:"id"`
	AuthorID    int64       `json:"author_id"`
	Title       string      `json:"title"`
	Description null.String `json:"description"`
}

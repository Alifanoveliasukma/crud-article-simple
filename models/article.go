package models

import "database/sql"

type Article struct {
	ID			int	    `json:"id"`
	Judul		string	`json:"judul"`
	GambarURL	string  `json:"gambar_url"`
	Kategori 	string  `json:"kategori"`
	Content		string  `json:"content"`
	CreatedAt	string  `json:"created_at"`
}

func CreateArticle(db *sql.DB, judul, gambarURL, kategori, content, createdAt string) error {
	query := "INSERT INTO article (judul, gambar_url, kategori, content, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, judul, gambarURL, kategori, content, createdAt)
	if err != nil {
		return err
	}
	return nil
}

func GetArticle(db *sql.DB, id int) (*Article, error) {
	query := "SELECT * FROM article WHERE id = ?"
	row := db.QueryRow(query, id)

	article := &Article{}
	err := row.Scan(&article.ID, &article.Judul, &article.GambarURL, &article.Kategori, &article.Content, &article.CreatedAt)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func UpdateArticle(db *sql.DB, id int, judul, gambarURL, kategori, content, createdAt string) error {
	query := "UPDATE article SET judul = ?, gambar_url = ?, kategori = ?, content = ?, created_at = ?"
	_, err := db.Exec(query, judul, gambarURL, kategori, content, createdAt) 
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(db *sql.DB, id int) error {
	query := "DELETE FROM article WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
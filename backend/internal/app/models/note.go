package models

type Note struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

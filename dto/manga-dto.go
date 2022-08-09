package dto

//MangaUpdateDTO is a model that client use when updating a book
type MangaUpdateDTO struct {
	MangaID       uint64 `json:"manga_id" form:"manga_id" binding:"required"`
	Title         string `json:"title" form:"title" binding:"required"`
	EnglishTitle  string `json:"en_title" form:"en_title"`
	JapaneseTitle string `json:"ja_title" form:"ja_title"`
	Author        string `json:"author" form:"author"`
	Artist        string `json:"artist" form:"artist"`
	Status        string `json:"status" form:"status" binding:"required"`
	Synopsis      string `json:"synopsis" form:"synopsis" binding:"required"`
	UserID        uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

//MangaCreateDTO is is a model that client use when create a new book
type MangaCreateDTO struct {
	Title         string `json:"title" form:"title" binding:"required"`
	EnglishTitle  string `json:"en_title" form:"en_title" binding:"required"`
	JapaneseTitle string `json:"ja_title" form:"ja_title" binding:"required"`
	Author        string `json:"author" form:"author" binding:"required"`
	Artist        string `json:"artist" form:"artist"`
	Status        string `json:"status" form:"status" binding:"required"`
	Synopsis      string `json:"synopsis" form:"synopsis"`
	UserID        uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

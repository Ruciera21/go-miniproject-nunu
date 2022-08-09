package entity

// * Manga represents manga table in database
type Manga struct {
	MangaID       uint64 `gorm:"primary_key:serial;column:manga_id" db:"manga_id"`
	Title         string `gorm:"type:varchar(255);not null;column:title" db:"title"`
	EnglishTitle  string `gorm:"type:varchar(255);column:english_title" db:"english_title"`
	JapaneseTitle string `gorm:"type:varchar(255);column:japaneses_title" db:"japaneses_title"`
	Author        string `gorm:"type:varchar(100);not null;column:author" db:"author"`
	Artist        string `gorm:"type:varchar(100);column:artist" db:"artist"`
	Status        string `gorm:"type:smallint;not null;column:status" db:"status"`
	PublishedOn   string `gorm:"type:date;column:published_on" db:"published_on"`
	FinishedOn    string `gorm:"type:date;column:published_on;default null" db:"finished_on"`
	Synopsis      string `gorm:"type:text;column:Synopsis" db:"synopsis"`
	UserID        uint64 `gorm:"not null" json:"-" db:"-"`
	User          User   `gorm:"foreignkey:UserID;not null" db:"user"`
}

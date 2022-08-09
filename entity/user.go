package entity

//User represents users table in database
type User struct {
	ID       uint64   `gorm:"primary_key:serial" db:"id"`
	Name     string   `gorm:"type:varchar(255)" db:"name"`
	Email    string   `gorm:"uniqueIndex;type:varchar(255)" db:"email"`
	Password string   `gorm:"->;<-;not null" db:"-"`
	Token    string   `gorm:"-" db:"token,omitempty"`
	Manga    *[]Manga `db:"books,omitempty"`
}

package repository

import (
	"mini-project-nunu/entity"

	"gorm.io/gorm"
)

// * MangaRepository is a contract for how mangaRepository behaves
type MangaRepository interface {
	InsertManga(m entity.Manga) entity.Manga
	UpdateManga(m entity.Manga) entity.Manga
	DeleteManga(m entity.Manga)
	AllManga() []entity.Manga
	FindMangaByID(mangaID uint64) entity.Manga
}

type mangaConnection struct {
	connection *gorm.DB
}

// * NewMangaRepository creates an instance of MangaRepository
func NewMangaRepository(dbConn *gorm.DB) MangaRepository {
	return &mangaConnection{connection: dbConn}
}

func (db *mangaConnection) InsertManga(m entity.Manga) entity.Manga {
	db.connection.Save(&m)
	db.connection.Preload("User").Find(&m)
	return m
}

func (db *mangaConnection) UpdateManga(m entity.Manga) entity.Manga {
	db.connection.Save(&m)
	db.connection.Preload("User").Find(&m)
	return m
}

func (db *mangaConnection) DeleteManga(m entity.Manga) {
	db.connection.Delete(&m)
}

func (db *mangaConnection) AllManga() []entity.Manga {
	var allManga []entity.Manga
	db.connection.Preload("User").Find(&allManga)
	return allManga
}

func (db *mangaConnection) FindMangaByID(mangaID uint64) entity.Manga {
	var manga entity.Manga
	db.connection.Preload("User").Find(&manga)
	return manga
}

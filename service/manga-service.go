package service

import (
	"fmt"
	"log"
	"mini-project-nunu/dto"
	"mini-project-nunu/entity"
	"mini-project-nunu/repository"

	"github.com/mashingan/smapping"
)

// * MangaService is a contract for what mangaService can do to db
type MangaService interface {
	Insert(m dto.MangaCreateDTO) entity.Manga
	Update(m dto.MangaUpdateDTO) entity.Manga
	Delete(m entity.Manga)
	AllManga() []entity.Manga
	FindByID(mangaID uint64) entity.Manga
	IsAllowedToEdit(UserID string, mangaID uint64) bool
}

type mangaService struct {
	mangaRepository repository.MangaRepository
}

// NewMangaService ....
func NewMangaService(mangaRepo repository.MangaRepository) MangaService {
	return &mangaService{mangaRepository: mangaRepo}
}

func (s *mangaService) Insert(m dto.MangaCreateDTO) entity.Manga {
	manga := entity.Manga{}
	err := smapping.FillStruct(&manga, smapping.MapFields(&m))
	if err != nil {
		log.Fatalf("Error mapping %v :", err)
	}
	response := s.mangaRepository.InsertManga(manga)
	return response
}

func (s *mangaService) Update(m dto.MangaUpdateDTO) entity.Manga {
	manga := entity.Manga{}
	err := smapping.FillStruct(&manga, smapping.MapFields(&m))
	if err != nil {
		log.Fatalf("Error mapping %v :", err)
	}
	response := s.mangaRepository.UpdateManga(manga)
	return response
}

func (s *mangaService) Delete(m entity.Manga) {
	s.mangaRepository.DeleteManga(m)
}

func (s *mangaService) AllManga() []entity.Manga {
	return s.mangaRepository.AllManga()
}

func (s *mangaService) FindByID(mangaID uint64) entity.Manga {
	return s.mangaRepository.FindMangaByID(mangaID)
}

func (s *mangaService) IsAllowedToEdit(userID string, mangaID uint64) bool {
	m := s.mangaRepository.FindMangaByID(mangaID)
	id := fmt.Sprintf("%v", m.UserID)
	return userID == id
}
